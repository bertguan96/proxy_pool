package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"project/proxy_pool/api"
	"project/proxy_pool/config"
)

var runType = flag.String("type", "", "please input run type!")

func init() {
	flag.Parse()
}

// main 启动入口
func main() {
	switch *runType {
	case "schedule":
		runSchedule()
	case "server":
		runServer()
	default:
		log.Fatalln("please input current type!")
		return
	}
	log.Printf("proxy start success！\nversion: %0.1f \n name: %s", config.Version, config.Name)
}

// runSchedule 运行定时任务
func runSchedule() {
	c := make(chan os.Signal)
	signal.Notify(c)
	go func() {
		api.StartWorker()
	}()
	<-c
	log.Println("Proxy Schedule Task Start Success！")
}

// runServer 运行Web服务
func runServer() {
	api.StartServer()
	log.Println("Proxy Web Server Task Start Success！")
}
