package api

import (
	"github.com/robfig/cron/v3"
	"log"
	"project/proxy_pool/config"
	"project/proxy_pool/worker"
)

func StartWorker() {
	c := cron.New()
	// 启动拉取线程
	if _, err := c.AddFunc(config.CronPull, worker.PullWorker); err != nil {
		log.Fatalln("check task execute error!", err)
		return
	}
	// 启动检测线程
	if _, err := c.AddFunc(config.CronCheck, worker.CheckWorker); err != nil {
		log.Fatalln("check task execute error!", err)
		return
	}
	c.Start()
}
