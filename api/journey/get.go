package journey

import (
	connect_firebase "backend/firebase"
	models "backend/model"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func GetJourney(c *gin.Context) {
	// get data
	client := connect_firebase.Connection()
	defer client.Close()

	docs, err := client.Collection("journey").Documents(c).GetAll()
	if err != nil {
		log.Println(err)
	}

	datas := []models.Journey{}

	var journey models.Journey

	for _, doc := range docs {
		// generate id from firebase
		journey.ID = doc.Ref.ID
		journey.CreatedAt = doc.CreateTime
		journey.UpdatedAt = doc.UpdateTime
		date_earned := doc.Data()["date"]
		journey.Date = date_earned.(time.Time)
		doc.DataTo(&journey)
		datas = append(datas, journey)
	}

	c.JSON(200, gin.H{
		"message": datas,
	})
}

func GetJourneyById(c *gin.Context) {
	client := connect_firebase.Connection()
	defer client.Close()

	id := c.Param("id")
	is_any := false

	// get data from firebase
	docs, err := client.Collection("journey").Documents(c).GetAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	var journey models.Journey
	for _, doc := range docs {
		if doc.Ref.ID == id {
			// generate id from firebase
			is_any = true
			journey.ID = doc.Ref.ID
			journey.CreatedAt = doc.CreateTime
			journey.UpdatedAt = doc.UpdateTime
			date_earned := doc.Data()["date"]
			journey.Date = date_earned.(time.Time)
			doc.DataTo(&journey)
		}
	}

	if is_any == false {
		c.JSON(404, gin.H{
			"message": "Not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": journey,
	})
}
