package main

import (
	"fmt"
	"serverGo/foundation/logger"
	"serverGo/foundation/mysql"
	"serverGo/foundation/routers"
	"serverGo/foundation/settings"
)

func main() {
	//utils.TestLocalTime()
	//加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("加载配置文件失败！ :%v\n", err)
		return
	}
	// 日志配置
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("初始化日志失败！ : %v\n", err)
		return
	}
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("mysql初始化失败！ ： %v\n", err)
		return
	}

	r := routers.SetupRouter(settings.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
