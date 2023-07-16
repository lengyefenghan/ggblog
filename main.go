package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"lengyefenghan.top/internal/router"
)

func main() {
	// 日志文件如果不存在就创建
	logfile, err := os.OpenFile("ggblog.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		//很有可能是权限问题哦
		fmt.Println(time.Now().UTC().Local(), "error:Open log or Create log file failed !!! ")
		os.Exit(1)
	}
	//设置日志输出目录
	log.SetOutput(logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	//启动路由
	log.Println("router.Start")
	router.Start()

}
