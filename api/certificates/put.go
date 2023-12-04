package certificates

import (
	connect_firebase "backend/firebase"
	models "backend/model"

	"github.com/gin-gonic/gin"
)

func EditCertificate(c *gin.Context) {
	// edit data
	client := connect_firebase.Connection()
	defer client.Close()

	id := c.Param("id")

	var certificate models.Certificate
	err := c.BindJSON(&certificate)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	_, err = client.Collection("certificates").Doc(id).Set(c, map[string]interface{}{
		"title":       certificate.Title,
		"authority":   certificate.Authority,
		"credential":  certificate.Credential,
		"url":         certificate.URL,
		"date_earned": certificate.DateEarned,
	})

	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"message": "Edit certificate successfully",
	})
}
