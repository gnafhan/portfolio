package projects

import (
	connect_firebase "backend/firebase"
	models "backend/model"

	"github.com/gin-gonic/gin"
)

func PostProject(c *gin.Context) {
	client := connect_firebase.Connection()
	defer client.Close()

	// post data to firebase with Socme struct
	var projects models.Project
	err := c.BindJSON(&projects)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	_, _, err = client.Collection("projects").Add(c, map[string]interface{}{
		"title":        projects.Title,
		"decription":   projects.Description,
		"image":        projects.Image,
		"url":          projects.URL,
		"technologies": projects.Technologies,
	})

	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"message": "Post project successfully",
	})
}
