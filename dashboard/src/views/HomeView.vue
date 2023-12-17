<script setup lang="ts">
import { onMounted, ref, computed } from 'vue'
import { useRequest } from 'vue-request'
import axios from 'axios'
import ProgressSpinner from 'primevue/progressspinner'
import Skeleton from 'primevue/skeleton'
import { formatNumber, formatBytes, formatTime } from '@/utils'
import HitsChart from '@/components/HitsChart.vue'
import type { StatInstData, APIStatus } from '@/api/v0'
import { tr } from '@/lang'

const now = ref(new Date())
setInterval(() => {
	now.value = new Date()
}, 1000)

const { data, error, loading } = useRequest(
	async () => (await axios.get<APIStatus>('/api/v0/status')).data,
	{
		pollingInterval: 5000,
		loadingDelay: 500,
		loadingKeep: 3000,
	},
)

const status = computed(() => error.value ? 'error' : data.value && data.value.enabled ? 'enabled' : 'disabled')

const stat = computed(() => {
	if (!data.value) {
		return
	}
	const stat = data.value.stats
	stat.days = cutDays(stat.days, stat.date.year, stat.date.month)
	stat.prev.days = cutDays(stat.prev.days, stat.date.year, stat.date.month - 1)

	stat.days[stat.date.day] = stat.hours.reduce((sum, v) => ({
		hits: sum.hits + v.hits,
		bytes: sum.bytes + v.bytes,
	}))
	stat.months[stat.date.month] = stat.days.reduce((sum, v) => ({
		hits: sum.hits + v.hits,
		bytes: sum.bytes + v.bytes,
	}))
	stat.years[stat.date.year.toString()] = stat.months.reduce((sum, v) => ({
		hits: sum.hits + v.hits,
		bytes: sum.bytes + v.bytes,
	}))
	return stat
})

function formatHour(hour: number): string {
	const offset = -new Date().getTimezoneOffset()
	let min = hour * 60 + offset
	hour = Math.floor(min / 60) % 24
	min %= 60
	if (hour < 0) {
		hour += 24
	}
	if (min < 0) {
		min += 60
	}
	return `${hour}:${min.toString().padStart(2, '0')}`
}

function cutDays(days: StatInstData[], year: number, month: number): StatInstData[] {
	const dayCount = new Date(year, month, 0).getDate()
	days.length = dayCount
	return days
}

function formatDay(day: number): string {
	if (!stat.value) {
		return ''
	}
	const date = new Date(Date.UTC(stat.value.date.year, stat.value.date.month, day))
	return `${date.getMonth() + 1}-${date.getDate()}`
}

function formatMonth(month: number): string {
	if (!stat.value) {
		return ''
	}
	const date = new Date(Date.UTC(stat.value.date.year, month + 1, 1))
	return `${date.getFullYear()}-${(date.getMonth() + 1).toString().padStart(2, '0')}`
}

function getDaysInMonth(): number {
	const date = new Date()
	const days = new Date(date.getFullYear(), date.getMonth() + 1, 0).getDate()
	return date.getDate() / days
}
</script>

<template>
	<main>
		<h1>Go-OpenBmclAPI {{ tr('title.dashboard') }}</h1>
		<div class="basic-info">
			<div
				class="info-status"
				:status="status"
				:style="{
					'--status-text': `'${tr(`badge.server.status.${status}`)}'`,
				}"
			></div>

			<ProgressSpinner v-if="loading" class="polling" strokeWidth="6"/>
			<div v-if="error">
				<b>{{ error }}</b>
			</div>
			<div v-else-if="data">
				{{ tr('message.server.run-for') }}
				<span class="info-uptime">
					{{ formatTime(now.getTime() - new Date(data.startAt).getTime()) }}
				</span>
			</div>
		</div>
		<div class="chart-card">
			<h3>Day</h3>
			<HitsChart
				v-if="data && stat"
				class="hits-chart"
				:max="25"
				:offset="23"
				:data="stat.hours"
				:oldData="stat.prev.hours"
				:current="stat.date.hour + new Date().getMinutes() / 60"
				:formatXLabel="formatHour"
			/>
			<Skeleton v-else class="hits-chart"/>
		</div>
		<div class="chart-card">
			<h3>Month</h3>
			<HitsChart
				v-if="data && stat"
				class="hits-chart"
				:max="31"
				:offset="29"
				:data="stat.days"
				:oldData="stat.prev.days"
				:current="stat.date.day + new Date().getHours() / 24"
				:formatXLabel="formatDay"
			/>
			<Skeleton v-else class="hits-chart"/>
		</div>
		<div class="chart-card">
			<h3>Year</h3>
			<HitsChart
				v-if="data && stat"
				class="hits-chart"
				:max="14"
				:offset="12"
				:data="stat.months"
				:oldData="stat.prev.months"
				:current="stat.date.month + getDaysInMonth()"
				:formatXLabel="formatMonth"
			/>
			<Skeleton v-else class="hits-chart"/>
		</div>
		<!-- TODO: show yearly chart -->
	</main>
</template>
<style scoped>
.basic-info {
	display: flex;
	flex-direction: row;
	align-items: center;
	height: 4rem;
	font-weight: 200;
}

.basic-info > div {
	display: inline-block;
}

.info-status {
	--flash-from: unset;
	--flash-out: var(--flash-from);
	display: inline-flex !important;
	flex-direction: row;
	align-items: center;
	padding: 0.5rem;
	margin: 0.5rem;
	border-radius: 0.2rem;
	font-weight: 800;
	user-select: none;
	cursor: pointer;
}

.info-status[status='enabled'] {
	--flash-from: #fff;
	--flash-to: #11dfc3;
	color: #fff;
	background-color: #28a745;
	animation: flash 1s infinite;
}

.info-status[status='disabled'] {
	--flash-from: #fff;
	--flash-to: #e61a05;
	color: #fff;
	background-color: #f89f1b;
	animation: flash 3s infinite;
}

.info-status[status='error'] {
	--flash-from: #8a8dac;
	color: #fff;
	background-color: #bfadad;
}

.info-status::before {
	content: ' ';
	display: inline-block;
	width: 1.05rem;
	height: 1.05rem;
	margin-right: 0.5rem;
	border: solid #fff 0.25rem;
	border-radius: 50%;
	background-color: var(--flash-out);
	box-shadow: #fff8 inset 0 0 2px;
	transition: background-color 0.15s;
}

.info-status::after {
	content: var(--status-text);
}

.polling {
	width: 1.5rem;
	margin-right: 0.2rem;
}

.info-uptime {
	font-weight: 700;
	font-style: italic;
}

.chart-card {
	margin-bottom: 1rem;
}

.hits-chart {
	width: 45rem !important;
	height: 13rem !important;
}

@media (max-width: 50rem) {
	.hits-chart {
		width: 100% !important;
	}
}
</style>