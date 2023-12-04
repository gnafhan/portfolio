package journey

import (
	connect_firebase "backend/firebase"
	models "backend/model"

	"github.com/gin-gonic/gin"
)

func EditJourney(c *gin.Context) {
	// edit data
	client := connect_firebase.Connection()
	defer client.Close()

	id := c.Param("id")

	var journey models.Journey
	err := c.BindJSON(&journey)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	_, err = client.Collection("journey").Doc(id).Set(c, map[string]interface{}{
		"title":       journey.Title,
		"description": journey.Description,
		"date":        journey.Date,
	})

	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"message": "Edit journey successfully",
	})
}
