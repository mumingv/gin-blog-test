package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mumingv/gin-blog/models"
	"github.com/mumingv/gin-blog/util"
)

// AuthMiddleware 认证登录中间件
func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		fmt.Println("AuthMiddleware uri = ", uri)

		if uri == "/admin/login" || uri == "/admin/logout" {
			return
		}

		// 获取session
		// 判断nil，请登录
		sess := util.GetSess(c, "user")
		if sess == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errmsg": "登录状态失败，请重新登录，地址：http://ip:port/admin/login",
			})
			c.Abort()
			return
		}

		// 解析session的用户信息，println
		var member models.User
		json.Unmarshal(sess.([]byte), &member)
		fmt.Println("models.User = ", member.Username)
		c.Next()
	}
}
