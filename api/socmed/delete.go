package socmed

import (
	connect_firebase "backend/firebase"

	"github.com/gin-gonic/gin"
)

func DeleteSocmed(c *gin.Context) {
	// delete data
	client := connect_firebase.Connection()
	defer client.Close()

	id := c.Param("id")

	_, err := client.Collection("socmed").Doc(id).Delete(c)

	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"message": "Delete social media successfully",
	})
}
