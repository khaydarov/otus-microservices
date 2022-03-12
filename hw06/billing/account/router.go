package account

import "github.com/gin-gonic/gin"

func RegisterRoutes(route *gin.RouterGroup) {
	route.GET("/:id", GetAccount())
	route.POST("/", CreateAccount())
}

func GetAccount() func (c *gin.Context) {
	return func (c *gin.Context) {
	}
}

func CreateAccount() func (c *gin.Context) {
	return func (c *gin.Context) {
	}
}
