package router

import (
	"github.com/gin-gonic/gin"
	"lengyefenghan.top/pkg/admin"
	"lengyefenghan.top/pkg/blogs"
)

func Start() {

	route := gin.Default()
	ad := route.Group("/admin")
	{
		ad.GET("/login", admin.Adminlogin)
		ad.GET("/logout", admin.Adminlogout)
		ad.GET("/assets", admin.AdminGetAssets)
		ad.GET("/api", admin.AdminApi)
	}
	bg := route.Group("/blogs")
	{
		bg.GET("/", blogs.BlogIndex)
		bg.GET("/d/*date", blogs.BlogDate)
		bg.GET("/i/*id", blogs.BlogID)
		// bg.GET("/t/*title", blogs.BlogID)
	}

	route.Run(":8080")

}
