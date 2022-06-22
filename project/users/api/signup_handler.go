package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"users/pkg/user"
)

func SignUpHandler(userRepo user.Repository) func(c *gin.Context) {
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
		err := userRepo.Save(u)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
				"data":    gin.H{},
			})

			return
		}

		err = createAccount(u.ID)
		if err != nil {
			_ = userRepo.Delete(u)
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

func createAccount(userID user.ID) error {
	endpoint := fmt.Sprintf("%s/internal/billing/createAccount", os.Getenv("BILLING_HOST"))
	data := map[string]interface{}{
		"user_id": userID.GetValue(),
	}

	body, _ := json.Marshal(data)

	request, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.New("internal service error")
	}

	return nil
}
