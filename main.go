package main

import (
	"douyin/commom"
	"douyin/config"
)

func main() {
	config.InitConfig() // 加载配置文件
	commom.InitDB()     // 数据局初始化
}
