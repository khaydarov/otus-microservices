package web

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"hw06/user/domain/user"
	"io/ioutil"
	"net/http"
	"os"
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
				"error": err.Error(),
			})

			return
		}

		newUser := user.NewUser(data.Login, data.FirstName, data.LastName)
		err := repository.Store(newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		err = createBillingAccount(newUser.ID.Value)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"data": gin.H{
				"id": newUser.ID.Value,
				"password": newUser.Password.Value,
			},
		})
	}
}

func createBillingAccount(ownerID string) error {
	endpoint := fmt.Sprintf("%s/accounts",
		os.Getenv("BILLING_SERVICE"),
	)

	var jsonData = []byte(fmt.Sprintf(`{"ownerID": "%s"}`, ownerID))
	response, _ := http.Post(endpoint, "application/json; charset=UTF-8", bytes.NewBuffer(jsonData))
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	if response.StatusCode != http.StatusCreated {
		return errors.New("account was not created")
	}

	return nil
}
