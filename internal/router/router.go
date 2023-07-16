package router

import (
	"github.com/gin-gonic/gin"
	"lengyefenghan.top/pkg/admin"
	"lengyefenghan.top/pkg/api"
)

func Start() {

	route := gin.Default()
	route.LoadHTMLGlob("web/theme/default/**/*")

	//后台系统
	ad := route.Group("/admin")
	{

		ad.GET("/login", admin.Adminlogin)
		ad.GET("/logout", admin.Adminlogout)
		ad.GET("/assets", admin.AdminGetAssets)
		ad.GET("/api", admin.AdminApi)
	}

	//实际渲染后的博客
	bg := route.Group("/public")
	{
		//正常渲染静态文件
		bg.GET("/*", func(ctx *gin.Context) {

		})
	}
	//接口信息
	ai := route.Group("/api")
	{
		ai.POST("/login", api.ApiLogin)

	}

	route.Run(":8080")

}
