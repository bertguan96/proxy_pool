package worker

import (
	"encoding/json"
	"log"
	"project/proxy_pool/common"
	"project/proxy_pool/config"
	"project/proxy_pool/db"
	"time"
)

// pullCheck 拉取首次校验
func pullCheck(proxy *common.ProxyGetter) bool {
	requestCheck := CheckHttp(proxy.Host) || CheckHttps(proxy.Host)
	proxy.IsHttps = CheckHttps(proxy.Host) // 是否支持https
	if db.Exists(proxy.Id) {
		log.Printf("host %s has exist!", proxy.Host)
		return false
	}
	proxy.LastCheckTime = time.Now().String()
	proxy.CheckCount = 1
	return requestCheck
}

func PullWorker() {
	log.Printf("start proxy pull worker!")
	for k := range config.Proxy {
		funcExec := config.Proxy[k].(func() []*common.ProxyGetter)
		proxyGetter := funcExec() // 执行代理获取脚本
		for i := 0; i < len(proxyGetter); i++ {
			log.Printf("pull host %s from %s", proxyGetter[i].Host, proxyGetter[i].Name)
			// 通过校验，则加入redis缓存
			if pullCheck(proxyGetter[i]) {
				jsonStr, err := json.Marshal(proxyGetter[i])
				if err != nil {
					log.Fatalln("json marshal failed！")
					return
				}
				db.AddIp(proxyGetter[i].Host, string(jsonStr))
			} else {
				log.Printf("host %s check failed!", proxyGetter[i].Host)
			}
		}
	}
}
