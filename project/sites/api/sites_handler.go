package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sites/pkg/site"
	"strings"
)

func PostSitesHandler(siteRepo site.Repository) func(c *gin.Context) {
	type Body struct {
		Title   string `json:"title"`
		Domains string `json:"domains"`
	}

	return func(c *gin.Context) {
		userID, ok := c.Get("UserID")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "invalid token",
			})

			return
		}

		var body Body
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})

			return
		}

		domains := strings.Split(body.Domains, ",")
		newSite := site.NewSite(userID.(string), body.Title, domains)
		err := siteRepo.Store(newSite)
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
			"message": "",
			"data": gin.H{
				"id": newSite.ID.GetValue(),
			},
		})
	}
}

func GetSitesHandler(siteRepo site.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		userID, ok := c.Get("UserID")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "invalid token",
			})

			return
		}

		sites := siteRepo.FindByUserID(userID.(string))

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"data":    sites,
		})
	}
}
