package web

import (
	"github.com/gin-gonic/gin"
	"hw06/user/domain/session"
	"hw06/user/domain/user"
	"net/http"
)

type loginData struct {
	Login 		string `json:"login"`
	Password 	string `json:"password"`
}

// Login signs in client with login and password
func Login(userRepository user.Repository, sessionRepository session.Repository) func (c *gin.Context) {
	return func (c *gin.Context) {
		var data loginData
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		actualUser := userRepository.FindUserByLoginAndPassword(data.Login, data.Password)

		if actualUser == nil {
			c.JSON(http.StatusOK, gin.H{
				"error": "User not found",
			})

			return
		}

		newSession := session.CreateSession(
			actualUser.ID,
			c.GetHeader("user-agent"),
			c.ClientIP(),
		)
		err := sessionRepository.Store(newSession)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Couldn't create session",
			})

			return
		}

		token, err := session.CreateAccessToken(actualUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Couldn't create access token",
			})

			return
		}

		c.JSON(200, gin.H{
			"data": gin.H{
				"accessToken": token,
				"refreshToken": newSession.Token.Value,
			},
		})
	}
}
