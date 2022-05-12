package api

import "github.com/gin-gonic/gin"

func RootHandler() func (c *gin.Context) {
	return func (c *gin.Context) {
		c.JSON(200, "Hello to payment service!")
	}
}
