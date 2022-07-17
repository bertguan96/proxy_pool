package common

import (
	"encoding/json"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// ProxyGetter 代理管理类，所有的代理拉取都从这里来

type ProxyGetter struct {
	Id            string `json:"id"`            // 生成的随机ID
	Name          string `json:"name"`          // 代理名称
	Host          string `json:"host"`          // 代理IP与Host
	Deadline      string `json:"deadline"`      // 过期时间
	LastCheckTime string `json:"lastCheckTime"` // 最后一次校验时间
	CheckCount    int64  `json:"checkCount"`    // 校验次数
	Lock          bool   `json:"lock"`          // 是否锁定（有的业务不能与其他业务一起共享IP所以需要独占）
	IsHttps       bool   `json:"isHttps"`       // 是否支持https
}

// QinGuo 需要JSON结构的接口才可以导出
func QinGuo() []*ProxyGetter {
	var proxyResult = make([]*ProxyGetter, 0)
	var result map[string]interface{}
	resp, err := http.Get("") // 填写申请的URL
	if err != nil {
		log.Fatalln("request error！")
		return nil
	}
	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		err = json.Unmarshal(body, &result)
	}
	if result["Code"].(float64) == 0 {
		data := result["Data"].([]interface{})
		for i := 0; i < len(data); i++ {
			value := data[i].(map[string]interface{})
			proxy := &ProxyGetter{
				Id:            uuid.New().String(),
				Host:          value["host"].(string),
				Deadline:      value["deadline"].(string),
				Name:          "QinGuo",
				LastCheckTime: time.Now().String(),
				CheckCount:    0,
				Lock:          false,
			}
			proxyResult = append(proxyResult, proxy)
		}
	}
	return proxyResult
}
