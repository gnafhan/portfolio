package certificates

import (
	connect_firebase "backend/firebase"
	models "backend/model"
	"log"

	"github.com/gin-gonic/gin"
)

func GetCertificates(c *gin.Context) {
	// get data
	client := connect_firebase.Connection()
	defer client.Close()

	docs, err := client.Collection("certificates").Documents(c).GetAll()
	if err != nil {
		log.Println(err)
	}

	datas := []models.Certificate{}

	var certificate models.Certificate

	for _, doc := range docs {
		// generate id from firebase
		certificate.ID = doc.Ref.ID
		certificate.CreatedAt = doc.CreateTime
		certificate.UpdatedAt = doc.UpdateTime
		doc.DataTo(&certificate)
		datas = append(datas, certificate)
	}

	c.JSON(200, gin.H{
		"message": datas,
	})
}

func GetCertificateById(c *gin.Context) {
	client := connect_firebase.Connection()
	defer client.Close()

	id := c.Param("id")

	// get data from firebase
	docs, err := client.Collection("certificates").Documents(c).GetAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	var certificate models.Certificate
	for _, doc := range docs {
		if doc.Ref.ID == id {
			// generate id from firebase
			certificate.ID = doc.Ref.ID
			certificate.CreatedAt = doc.CreateTime
			certificate.UpdatedAt = doc.UpdateTime
			doc.DataTo(&certificate)
		}
	}

	c.JSON(200, gin.H{
		"message": certificate,
	})
}
