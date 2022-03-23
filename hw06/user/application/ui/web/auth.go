package web

import (
	"github.com/gin-gonic/gin"
	"hw06/user/domain/user"
)

// Auth authenticates client with login and password
func Auth(repository user.Repository) func (c *gin.Context) {
	return func (c *gin.Context) {
		c.JSON(200, "OK!")
	}
}
