package projects

import (
	connect_firebase "backend/firebase"
	models "backend/model"

	"github.com/gin-gonic/gin"
)

func EditProject(c *gin.Context) {
	// edit data
	client := connect_firebase.Connection()
	defer client.Close()

	id := c.Param("id")

	var projects models.Project
	err := c.BindJSON(&projects)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	_, err = client.Collection("projects").Doc(id).Set(c, map[string]interface{}{
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
		"message": "Edit project successfully",
	})
}
