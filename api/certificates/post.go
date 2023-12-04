package certificates

import (
	connect_firebase "backend/firebase"
	models "backend/model"

	"github.com/gin-gonic/gin"
)

func PostCertificate(c *gin.Context) {
	client := connect_firebase.Connection()
	defer client.Close()

	// post data to firebase with Certificate struct
	var certificate models.Certificate
	err := c.BindJSON(&certificate)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	_, _, err = client.Collection("certificates").Add(c, map[string]interface{}{
		"title":       certificate.Title,
		"authority":   certificate.Authority,
		"credential":  certificate.Credential,
		"url":         certificate.URL,
		"date_earned": certificate.DateEarned,
		"image":       certificate.Image,
	})

	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"message": "Post certificate successfully",
	})
}
