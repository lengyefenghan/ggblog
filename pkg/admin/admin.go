package admin

import "github.com/gin-gonic/gin"

// 进行登录校验
func Adminlogin(c *gin.Context) {
	c.String(200, "login")

}

// 登出程序服务端收回token权限
func Adminlogout(c *gin.Context) {
	c.String(200, "logout")
}

// 获取资源会对资源进行安全过滤保证安全
func AdminGetAssets(c *gin.Context) {
	c.String(200, "assets")
}

// 多个接口的实现
func AdminApi(c *gin.Context) {
	c.String(200, "api")
}
