package main

import (
	"GameMall/dal"
	"GameMall/logs"
	"GameMall/router"
)

func main() {
	//初始化日志打印配置
	logs.InitLogger()

	//初始化数据库配置及连接
	err := dal.InitDB()
	if err != nil {
		panic(err)
	}

	//初始化路由配置
	r := router.InitRouter()

	//启动
	_ = r.Run()
}