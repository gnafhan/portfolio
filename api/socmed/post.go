package socmed

import (
	connect_firebase "backend/firebase"
	models "backend/model"

	"github.com/gin-gonic/gin"
)

func PostSocmed(c *gin.Context) {
	client := connect_firebase.Connection()
	defer client.Close()

	// post data to firebase with Socme struct
	var socmed models.SocialMedia
	err := c.BindJSON(&socmed)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	_, _, err = client.Collection("socmed").Add(c, map[string]interface{}{
		"platform": socmed.Platform,
		"username": socmed.Username,
		"icon":     socmed.Icon,
		"url":      socmed.URL,
	})

	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"message": "Post social media successfully",
	})
}
