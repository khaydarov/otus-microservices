package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sites/pkg/site"
)

func PostSitesHandler(siteRepo site.Repository) func(c *gin.Context) {
	type Body struct {
		UserID  string `json:"user_id"`
		Title   string `json:"title"`
		Domains string `json:"domains"`
	}

	return func(c *gin.Context) {
		var body Body
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})

			return
		}

	}
}
