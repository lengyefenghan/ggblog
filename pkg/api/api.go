package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ApiLogin(c *gin.Context) {
	fmt.Println(c.Params)
}
