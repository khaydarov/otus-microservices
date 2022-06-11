package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"users/pkg/session"
	"users/pkg/user"
)

func LoginHandler() func(c *gin.Context) {
	type Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(c *gin.Context) {
		body := Body{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
				"data":    gin.H{},
			})

			return
		}

		userRepo := user.NewRepository()
		u, err := userRepo.FindByEmailAndPassword(user.NewEmail(body.Email), body.Password)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "user not found",
				"data":    gin.H{},
			})
		}

		newSession := session.CreateSession(
			u.ID,
			c.GetHeader("user-agent"),
			c.ClientIP(),
		)

		//sessionRepo := session.NewRepository()
		//err := sessionRepo.Store(newSession)
		//if err != nil {
		//	c.JSON(http.StatusInternalServerError, gin.H{
		//		"error": "Couldn't create session",
		//	})
		//
		//	return
		//}

		token, err := session.CreateAccessToken(u)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Couldn't create access token",
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"accessToken":  token,
				"refreshToken": newSession.Token.Value,
			},
		})
	}
}
