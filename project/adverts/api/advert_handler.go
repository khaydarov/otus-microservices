package api

import (
	"adverts/pkg/advert"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func PostAdvertHandler(advertRepo advert.Repository) func(c *gin.Context) {
	type Body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Link        string `json:"link"`
		Image       string `json:"image"`
		Dates       string `json:"dates"`
		Devices     string `json:"devices"`
		Hits        int    `json:"hits"`
		Cost        int    `json:"cost"`
	}

	return func(c *gin.Context) {
		var body Body
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
				"data":    gin.H{},
			})

			return
		}

		dates := strings.Split(body.Dates, ",")
		devices := strings.Split(body.Devices, ",")

		newAdvert := advert.NewAdvert(
			body.Title,
			body.Description,
			body.Link,
			body.Image,
			devices,
			dates,
			body.Hits,
			body.Cost,
		)

		err := advertRepo.Store(newAdvert)
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
				"id": newAdvert.ID.GetValue(),
			},
		})
	}
}
