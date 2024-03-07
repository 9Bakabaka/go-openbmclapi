package zh

var areaCN = map[string]string{
	"program.starting":     "正在启动 Go-OpenBmclApi v%s (%s)",
	"error.set.cluster.id": "启动前请在 config.yaml 内设置 cluster-id 和 cluster-secret !",
	"error.init.failed":    "无法初始化节点: %v",

	"program.exited":                          "节点正在退出, 代码 %d",
	"error.exit.please.read.faq":              "请在提交问题前阅读 https://github.com/LiterMC/go-openbmclapi?tab=readme-ov-file#faq",
	"warn.exit.detected.windows.open.browser": "检测到您是新手 Windows 用户. 我们正在帮助您打开浏览器 ...",

	"info.filelist.fetching":      "获取文件列表中",
	"error.filelist.fetch.failed": "文件列表获取失败: %v",

	"error.address.listen.failed": "无法监听地址 %s: %v",

	"info.cert.requesting":              "请求证书中, 请稍候 ...",
	"info.cert.requested":               "证书请求完毕, 域名为 %s",
	"error.cert.not.set":                "配置文件内没有提供证书",
	"error.cert.parse.failed":           "无法解析证书密钥对[%d]: %v",
	"error.cert.request.failed":         "证书请求失败: %v",
	"error.cert.requested.parse.failed": "无法解析已请求的证书: %v",

	"info.server.public.at":         "服务器已在 https://%s:%d (%s) 开放, 使用了 %d 个证书",
	"info.server.alternative.hosts": "备用域名:",
	"info.wait.first.sync":          "正在等待第一次同步 ...",
	"error.cluster.enable.failed":   "无法启用节点: %v",

	"warn.server.closing":          "关闭服务中 ...",
	"warn.server.closed":           "服务器已关闭.",
	"info.cluster.disabling":       "禁用节点中 ...",
	"error.cluster.disable.failed": "节点禁用失败: %v",
	"warn.cluster.disabled":        "节点已禁用",
	"warn.httpserver.closing":      "正在关闭 HTTP 服务器 ...",

	"info.check.start":                "开始在 %s 检测文件. 强检查 = %v",
	"info.check.done":                 "文件在 %s 检查完毕, 缺失 %d 个文件",
	"hint.check.checking":             "> 检查中 ",
	"warn.check.modified.size":        "找到修改过的文件: %q 的大小为 %d, 预期 %d",
	"warn.check.modified.hash":        "找到修改过的文件: %q 的哈希值为 %s, 预期 %s",
	"error.check.unknown.hash.method": "未知的哈希格式 %q",
	"error.check.open.failed":         "无法打开 %q: %v",
	"error.check.hash.failed":         "无法为 %s 计算哈希值: %v",

	"info.sync.prepare":          "准备同步中, 文件列表长度为 %d ...",
	"hint.sync.start":            "开始同步, 总计: %d, 字节: %s",
	"hint.sync.done":             "文件同步完成, 用时: %v, %s/s",
	"info.sync.none":             "所有文件已同步",
	"warn.sync.interrupted":      "同步已中断",
	"info.sync.config":           "同步配置: %#v",
	"hint.sync.total":            "总计: ",
	"hint.sync.downloading":      "> 下载中 ",
	"info.sync.downloaded":       "已下载 %s [%s] %.2f%%",
	"error.sync.download.failed": "下载失败 %s:\n\t%s",
	"error.sync.create.failed":   "无法创建 %s/%s: %v",

	"info.gc.start":       "已为 %s 开启垃圾收集器",
	"info.gc.done":        "垃圾收集已在 %s 完成",
	"warn.gc.interrupted": "垃圾收集器在 %s 中断",
	"info.gc.found":       "找到过期文件 %s",
	"error.gc.error":      "垃圾收集错误: %v",

	"error.config.read.failed":           "无法读取配置文件: %v",
	"error.config.encode.failed":         "无法编码配置文件: %v",
	"error.config.write.failed":          "无法写入配置文件: %v",
	"error.config.not.exists":            "未找到配置文件, 正在创建",
	"error.config.created":               "配置文件已创建, 请修改, 并重启程序",
	"error.config.parse.failed":          "无法解析配置文件: %v",
	"error.config.alias.user.not.exists": "WebDav 别名用户 %q 不存在",
}