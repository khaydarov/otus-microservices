package web

import (
	"github.com/gin-gonic/gin"
	"hw06/identity/domain/user"
	"net/http"
)

type registerData struct {
	Login 		string `json:"login"`
	FirstName 	string `json:"firstName"`
	LastName	string `json:"lastName"`
}

// Register Creates new user
func Register(repository user.Repository) func(c *gin.Context) {
	return func (c *gin.Context) {
		var data registerData
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"error": err.Error(),
			})

			return
		}

		newUser := user.NewUser(data.Login, data.FirstName, data.LastName)
		err := repository.Store(newUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusInternalServerError,
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": gin.H{
				"id": newUser.ID.Value,
				"password": newUser.Password.Value,
			},
		})
	}
}
