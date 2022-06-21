package api

import (
	"adverts/pkg/advert"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRelevantAdvertHandler(advertRepo advert.Repository, advertSelector advert.Selector) func(c *gin.Context) {
	return func(c *gin.Context) {
		dates := []string{"21-06-2022"}
		devices := []string{"ios"}
		advertIds := advertSelector.MatchAdvert(dates, devices)

		relevantAdvert, err := advertRepo.FindByID(advert.WithValue(advertIds[0]))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
				"data":    gin.H{},
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"data":    relevantAdvert,
		})
	}
}
