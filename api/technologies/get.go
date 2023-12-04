package technologies

import (
	connect_firebase "backend/firebase"
	models "backend/model"
	"log"

	"github.com/gin-gonic/gin"
)

func GetTechnology(c *gin.Context) {
	// get data
	client := connect_firebase.Connection()
	defer client.Close()

	docs, err := client.Collection("technologies").Documents(c).GetAll()
	if err != nil {
		log.Println(err)
	}

	datas := []models.Technology{}

	var technologies models.Technology

	for _, doc := range docs {
		// generate id from firebase
		technologies.ID = doc.Ref.ID
		technologies.CreatedAt = doc.CreateTime
		technologies.UpdatedAt = doc.UpdateTime
		doc.DataTo(&technologies)
		datas = append(datas, technologies)
	}

	c.JSON(200, gin.H{
		"message": datas,
	})
}

func GetTechnologyById(c *gin.Context) {
	client := connect_firebase.Connection()
	defer client.Close()

	id := c.Param("id")
	is_any := false

	// get data from firebase
	docs, err := client.Collection("technologies").Documents(c).GetAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	var technologies models.Technology
	for _, doc := range docs {
		if doc.Ref.ID == id {
			// generate id from firebase
			is_any = true
			technologies.ID = doc.Ref.ID
			technologies.CreatedAt = doc.CreateTime
			technologies.UpdatedAt = doc.UpdateTime
			doc.DataTo(&technologies)
		}
	}

	if is_any == false {
		c.JSON(404, gin.H{
			"message": "Not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": technologies,
	})
}
