package technologies

import (
	connect_firebase "backend/firebase"
	models "backend/model"

	"github.com/gin-gonic/gin"
)

func PostTechnology(c *gin.Context) {
	client := connect_firebase.Connection()
	defer client.Close()

	// post data to firebase with Socme struct
	var technologies models.Technology
	err := c.BindJSON(&technologies)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	_, _, err = client.Collection("technologies").Add(c, map[string]interface{}{
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
		"message": "Post technology successfully",
	})
}
