package main

import (
	"AuroraDisk/config"
	"AuroraDisk/model"
	"AuroraDisk/router"
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// todo:初始化前端
func init() {
	var frontendPath string
	var backendPath string
	flag.StringVar(&frontendPath, "front", "", "指定前端配置文件")
	flag.StringVar(&backendPath, "back", ".", "指定后端配置文件所在目录")
	flag.Parse()

	config.Load(backendPath)
}

func main() {
	//初始化服务器
	api := router.InitRouter()
	api.TrustedPlatform = "" //todo:配置白名单
	server := &http.Server{Handler: api}
	//关闭服务器
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP)
	go shutdown(sigChan, server)

	//服务器开始监听并运行(堵塞状态)
	if err := server.ListenAndServe(); err != nil {
		log.Println("服务器运行失败...", err.Error())
	}

	//会阻塞等待信号到来，若信号已接收或管道已关闭，它将立即返回(使用goroutine才能达到此效果)
	<-sigChan
	//关闭数据库连接
	model.DisConnect()
}

func shutdown(sigChan chan os.Signal, server *http.Server) {
	defer close(sigChan)
	sig := <-sigChan
	log.Printf("接收到%s信号\n，关闭服务器...", sig)
	ctx := context.Background()
	if config.SysCfg.Period != 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(config.SysCfg.Period)*time.Second)
		defer cancel()
	}

	if err := server.Shutdown(ctx); err != nil {
		log.Println("服务器关闭失败orz...", err.Error())
	}
}
