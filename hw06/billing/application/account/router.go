package account

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"hw06/billing/middlewares"
	"net/http"
)

func RegisterRoutes(route *gin.RouterGroup, db *pgx.Conn) {
	repository := NewRepository(db)

	route.POST("/accounts", CreateAccount(repository))
	route.GET("/myAccount", GetAccount(repository))
	route.POST("/depositAccount", DepositAccount(repository))
	route.POST("/withdrawAccount", WithdrawAccount(repository))
}

type createRequest struct {
	OwnerID string `json:"ownerId"`
}

func CreateAccount(repository Repository) func (c *gin.Context) {
	return func (c *gin.Context) {
		var data createRequest
		var err error
		if err = c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		newAccount := NewAccount(data.OwnerID)
		err = repository.Store(newAccount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"data": gin.H{
				"id": newAccount.ID.Value,
			},
		})
	}
}

func GetAccount(repository Repository) func (c *gin.Context) {
	return func (c *gin.Context) {
		credentials, ok := c.Get("user")
		if !ok {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		user := credentials.(middlewares.User)
		account := repository.GetByOwnerID(OwnerID{
			Value: user.ID,
		})

		if account == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "account not found",
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"id": account.ID.Value,
				"balance": account.Balance,
			},
		})
	}
}

type depositRequest struct {
	Amount 	int `json:"amount"`
}

func DepositAccount(repository Repository) func (c *gin.Context) {
	return func (c *gin.Context) {
		credentials, ok := c.Get("user")
		if !ok {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		var data depositRequest
		var err error
		if err = c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		user := credentials.(middlewares.User)
		account := repository.GetByOwnerID(OwnerID{
			Value: user.ID,
		})

		if account == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "account not found",
			})

			return
		}

		err = account.Deposit(data.Amount)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		err = repository.Update(account)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"id": account.ID.Value,
				"balance": account.Balance,
			},
		})
	}
}

type withdrawRequest struct {
	Amount	int `json:"amount"`
}

func WithdrawAccount(repository Repository) func (c *gin.Context) {
	return func (c *gin.Context) {
		credentials, ok := c.Get("user")
		if !ok {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		var data withdrawRequest
		var err error
		if err = c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		user := credentials.(middlewares.User)
		account := repository.GetByOwnerID(OwnerID{
			Value: user.ID,
		})

		if account == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "account not found",
			})

			return
		}

		err = account.Withdraw(data.Amount)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		err = repository.Update(account)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"id": account.ID.Value,
				"balance": account.Balance,
			},
		})
	}
}