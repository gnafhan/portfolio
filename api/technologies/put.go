package technologies

import (
	connect_firebase "backend/firebase"
	models "backend/model"

	"github.com/gin-gonic/gin"
)

func EditTechnology(c *gin.Context) {
	// edit data
	client := connect_firebase.Connection()
	defer client.Close()

	id := c.Param("id")

	var technologies models.Technology
	err := c.BindJSON(&technologies)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	_, err = client.Collection("technologies").Doc(id).Set(c, map[string]interface{}{
		"name":  technologies.Name,
		"icon":  technologies.Icon,
		"skill": technologies.Skill,
	})

	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"message": "Edit technology successfully",
	})
}
