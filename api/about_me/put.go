package about_me

import (
	connect_firebase "backend/firebase"
	models "backend/model"

	"github.com/gin-gonic/gin"
)

func EditAboutMe(c *gin.Context) {
	id := c.Param("id")

	client := connect_firebase.Connection()
	defer client.Close()

	// edit data

	var about_me models.AboutMe
	err := c.BindJSON(&about_me)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	_, err = client.Collection("about_me").Doc(id).Set(c, map[string]interface{}{
		"name":        about_me.Name,
		"title":       about_me.Title,
		"description": about_me.Description,
		"skills":      about_me.Skills,
		"is_selected": about_me.IsSelected,
	})

	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"message": "Edit about me successfully",
	})

}
