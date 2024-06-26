import { tr, getLang, Lang } from '@/lang'

const _EN = new Lang('en')
const _ZH = new Lang('zh')

export function formatNumber(num: number): string {
	const lang = getLang()
	if (_ZH.match(lang)) {
		return formatNumberZH(num)
	}
	return formatNumberEN(num)
}

const nUnitsUS = ['k', 'm', 'B', 'T', 'Q']

function formatNumberEN(num: number): string {
	if (num < 1000) {
		return num.toString()
	}
	var unit
	for (const u of nUnitsUS) {
		unit = u
		num /= 1000
		if (num < 1000) {
			break
		}
	}
	return `${num.toFixed(2)} ${unit}`
}

const nUnitsZH = ['万', '亿', '兆', '京']

function formatNumberZH(num: number): string {
	if (num < 9000) {
		return num.toString()
	}
	var unit
	for (const u of nUnitsZH) {
		unit = u
		num /= 10000
		if (num < 9000) {
			break
		}
	}
	return `${num.toFixed(2)} ${unit}`
}

const bUnits = ['KB', 'MB', 'GB', 'TB']

export function formatBytes(bytes: number): string {
	if (bytes < 1000) {
		return bytes.toString()
	}
	var unit
	for (const u of bUnits) {
		unit = u
		bytes /= 1024
		if (bytes < 1000) {
			break
		}
	}
	return `${bytes.toFixed(2)} ${unit}`
}

export function formatTime(ms: number): string {
	var unit = tr('unit.time.ms')
	if (ms > 800) {
		ms /= 1000
		unit = tr('unit.time.s')
		if (ms > 50) {
			ms /= 60
			unit = tr('unit.time.min')
			if (ms > 50) {
				ms /= 60
				unit = tr('unit.time.hour')
				if (ms > 22) {
					ms /= 24
					unit = tr('unit.time.day')
					if (ms > 350) {
						ms /= 356
						unit = tr('unit.time.year')
					}
				}
			}
		}
	}
	return `${ms.toFixed(2)} ${unit}`
}
