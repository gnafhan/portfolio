package socmed

import (
	connect_firebase "backend/firebase"
	models "backend/model"

	"github.com/gin-gonic/gin"
)

func EditSocmed(c *gin.Context) {
	// edit data
	client := connect_firebase.Connection()
	defer client.Close()

	id := c.Param("id")

	var socmed models.SocialMedia
	err := c.BindJSON(&socmed)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	_, err = client.Collection("socmed").Doc(id).Set(c, map[string]interface{}{
		"username": socmed.Username,
		"platform": socmed.Platform,
		"icon":     socmed.Icon,
		"url":      socmed.URL,
	})

	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"message": "Edit social media successfully",
	})
}
