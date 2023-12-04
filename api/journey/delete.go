package journey

import (
	connect_firebase "backend/firebase"

	"github.com/gin-gonic/gin"
)

func DeleteJourney(c *gin.Context) {
	// delete data
	client := connect_firebase.Connection()
	defer client.Close()

	id := c.Param("id")

	_, err := client.Collection("journey").Doc(id).Delete(c)

	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"message": "Delete journey successfully",
	})
}
