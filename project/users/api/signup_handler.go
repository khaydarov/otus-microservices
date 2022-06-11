package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"users/pkg/user"
)

func SignUpHandler() func(c *gin.Context) {
	type Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Type     int    `json:"type"`
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

		u := user.NewUser(body.Email, body.Password, body.Type)
		repo := user.NewRepository()
		err := repo.Save(u)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
				"data":    gin.H{},
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"id": u.ID.GetValue(),
			},
		})
	}
}
