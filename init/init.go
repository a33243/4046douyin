package init

import (
	"douyin/commom"
	"douyin/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Init() {
	config.InitConfig() // 加载配置文件
	commom.InitDB()     // 数据局初始化
	r := gin.Default()
	// 设置基本路径
	err := r.Run(fmt.Sprintf(":%d", config.Info.Port))
}
