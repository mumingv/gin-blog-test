package main

import (
	"fmt"
	"github.com/mumingv/gin-blog/dao"
	"github.com/mumingv/gin-blog/logger"
	"github.com/mumingv/gin-blog/models"
	"github.com/mumingv/gin-blog/routers"
	"github.com/mumingv/gin-blog/settings"
)

func main() {
	// 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("Loading config failed, err: %v\n", err)
	}

	// 初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err: %v\n", err)
		return
	}

	// MySQL
	// 连接数据库
	err := dao.InitMySQL(settings.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer dao.Close() // 程序退出关闭数据库连接
	// 模型绑定
	dao.DB.AutoMigrate(new(models.User),
		new(models.Category),
		new(models.Post),
		new(models.Config),
		new(models.Comment))

	// 注册路由
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err: %v\n", err)
	}
}
