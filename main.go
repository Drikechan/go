package main

import (
	"test-todolist/config"
	"test-todolist/respository/db/dao"
	"test-todolist/routes"
)

func main() {
	r := routes.Router()

	// 加载配置
	initConfig()

	r.Run(config.HttpPort)
}

func initConfig() {
	config.Init()
	dao.InitMysql()
}
