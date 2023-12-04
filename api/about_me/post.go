package about_me

import (
	connect_firebase "backend/firebase"
	models "backend/model"
	"time"

	"github.com/gin-gonic/gin"
)

func PostAboutMe(c *gin.Context) {
	client := connect_firebase.Connection()
	defer client.Close()

	// post data to firebase with AboutMe struct
	var about_me models.AboutMe
	err := c.BindJSON(&about_me)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	_, _, err = client.Collection("about_me").Add(c, map[string]interface{}{
		"name":        about_me.Name,
		"title":       about_me.Title,
		"description": about_me.Description,
		"skills":      about_me.Skills,
		// is selected if null then false
		"is_selected": about_me.IsSelected,
		"created_at":  time.Now(),
	})

	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"message": "Post about me successfully",
	})
}
