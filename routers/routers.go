package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mumingv/gin-blog/controller"
	"github.com/mumingv/gin-blog/logger"
	"github.com/mumingv/gin-blog/settings"
	"github.com/mumingv/gin-blog/util"
)

func helloHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"mode": settings.Conf.Mode,
		"host": settings.Conf.MySQLConfig.Host,
	})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 告诉gin框架模板文件引用的静态文件
	r.Static("/static", "static")
	// 告诉gin框架模板文件
	r.LoadHTMLGlob("templates/*/*")

	// 中间件注册
	r.Use(logger.GinLogger())
	// 集成Session
	util.InitSession(r)
	// 函数注册
	r.GET("/hello", helloHandler)

	// 前端系统

	// 后端系统
	v2Group := r.Group("admin")
	v2Group.Use(controller.AuthMiddleware())
	admin := controller.AdminController{}
	{
		// 主页面登录
		// localhost:9002/admin/login
		v2Group.GET("/login", admin.Login)
		v2Group.POST("/login", admin.Login)
		v2Group.GET("/logout", admin.Logout)

		// 主页
		// localhost:9002/admin/main
		v2Group.GET("/main", admin.Main)

		// 系统配置
		////页面展示
		v2Group.GET("/config", admin.Config)
		////提交更新
		v2Group.POST("/addconfig", admin.AddConfig)
		// 博文列表
		v2Group.GET("/index", admin.Index)

		// 博文添加
		//// 显示
		v2Group.GET("/article", admin.Article)
		//// 文章保存
		v2Group.POST("/save", admin.Save)
		//// 文章删除
		v2Group.GET("/delete", admin.PostDel)

		// 类目主页
		v2Group.GET("/category", admin.Category)
		// 类目增加
		v2Group.GET("/categoryadd", admin.CategoryAdd)
		// 类目保存
		v2Group.POST("/categorysave", admin.CategorySave)
		// 类目删除
		v2Group.GET("/categorydel", admin.CategoryDel)
	}
	return r
}
