package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"project/proxy_pool/common"
	"project/proxy_pool/db"
)

type ProxyPool struct {
}

func StartServer() {
	r := GetRouter()
	err := r.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatalln(" Web Server Task init error!")
		return
	}
	log.Println("Web Server Task Start Success！")
}

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/get", get)
	r.GET("/delete", delete)
	r.GET("/getAll", getAll)
	r.GET("/clear", clear)
	return r
}

func get(c *gin.Context) {
	ip := db.GetIp()
	proxy := &common.ProxyGetter{}
	err := json.Unmarshal([]byte(ip), proxy)
	if err != nil {
		log.Printf("json unmarshal failed!")
		c.JSON(http.StatusOK, "")
		return
	}
	c.JSON(http.StatusOK, proxy)
}

func delete(c *gin.Context) {
	id := c.GetString("id")
	res := db.DelIp(id)
	c.JSON(http.StatusOK, res)
}

func getAll(c *gin.Context) {
	ipMap := db.GetAll()
	proxyList := make([]*common.ProxyGetter, 0)
	for key := range ipMap {
		proxy := &common.ProxyGetter{}
		err := json.Unmarshal([]byte(ipMap[key]), proxy)
		if err != nil {
			log.Printf("json unmarshal failed!")
			return
		}
		proxyList = append(proxyList, proxy)
	}

	c.JSON(http.StatusOK, proxyList)
}

func clear(c *gin.Context) {
	res := db.ClearAll()
	c.JSON(http.StatusOK, res)
}
