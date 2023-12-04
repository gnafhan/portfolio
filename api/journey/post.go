package journey

import (
	connect_firebase "backend/firebase"
	models "backend/model"

	"github.com/gin-gonic/gin"
)

func PostJourney(c *gin.Context) {
	client := connect_firebase.Connection()
	defer client.Close()

	// post data to firebase with Socme struct
	var journey models.Journey
	err := c.BindJSON(&journey)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	_, _, err = client.Collection("journey").Add(c, map[string]interface{}{
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
		"message": "Post journey successfully",
	})
}
