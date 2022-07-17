package config

import "project/proxy_pool/common"

var (
	Version   = 1.0                     // 版本号
	Name      = "proxy pool"            // 名称
	CronCheck = "@every 1m"             // 校验定时
	CronPull  = "@every 1m"             // 拉取定时
	Proxy     = map[string]interface{}{ // 填写解析方法
		"QinGuo": common.QinGuo,
	}
	DBKey         = "proxy_pool"
	DBHost        = "" // Redis的主机地址
	DBPassword    = "" // Redis的密码
	DB            = 0
	ProxyAuth     = "" // 代理权限，当然也可以设置白名单
	HttpsValidUrl = "https://www.qq.com"
	HttpValidUrl  = "http://httpbin.org"
)
