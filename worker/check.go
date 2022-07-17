package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"project/proxy_pool/common"
	"project/proxy_pool/config"
	"project/proxy_pool/db"
	"time"
)

func CheckHttp(host string) bool {
	proxyUrl := fmt.Sprintf("http://%s@%s", config.ProxyAuth, host)
	uri, err := url.Parse(proxyUrl)
	if err != nil {
		log.Fatal("parse url error: ", err)
		return false
	}
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(uri),
		},
	}
	resp, err := client.Get(config.HttpValidUrl)
	if err != nil {
		log.Printf(err.Error())
		log.Printf("https valid failed, proxy url %s！", proxyUrl)
		return false
	}
	return resp.StatusCode == 200
}

func CheckHttps(host string) bool {
	proxyUrl := fmt.Sprintf("https://%s@%s", config.ProxyAuth, host)
	uri, err := url.Parse(proxyUrl)
	if err != nil {
		return false
	}
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(uri),
		},
	}
	resp, err := client.Get(config.HttpsValidUrl)
	if err != nil {
		log.Printf(err.Error())
		log.Printf("https valid failed, proxy url %s！", proxyUrl)
		return false
	}
	return resp.StatusCode == 200
}

func CheckWorker() {
	log.Printf("start proxy check worker!")
	ipMap := db.GetAll()
	for key := range ipMap {
		ipJson := ipMap[key]
		proxy := common.ProxyGetter{}
		err := json.Unmarshal([]byte(ipJson), &proxy)
		if err != nil {
			log.Printf("json Unmarshal failed!")
			return
		}
		// 如果都不满足条件则移除
		if !CheckHttps(proxy.Host) && !CheckHttp(proxy.Host) {
			log.Printf("proxy %s time out! will be removed!", proxy.Host)
			db.DelIp(proxy.Host)
		}
		proxy.LastCheckTime = time.Now().String() // 计算检查时间
		proxy.CheckCount = 1 + proxy.CheckCount   // 计算校验次数
	}
}
