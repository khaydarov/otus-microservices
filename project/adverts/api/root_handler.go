package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RootHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello from adverts")
	}
}
