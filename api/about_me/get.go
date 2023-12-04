package about_me

import (
	connect_firebase "backend/firebase"
	models "backend/model"

	"github.com/gin-gonic/gin"
)

func GetAboutMe(c *gin.Context) {
	client := connect_firebase.Connection()
	defer client.Close()

	// get data from firebase
	docs, err := client.Collection("about_me").Documents(c).GetAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	datas := []models.AboutMe{}

	var about_me models.AboutMe

	for _, doc := range docs {
		// generate id from firebase
		about_me.ID = doc.Ref.ID
		about_me.CreatedAt = doc.CreateTime
		about_me.UpdatedAt = doc.UpdateTime
		doc.DataTo(&about_me)
		datas = append(datas, about_me)
	}

	c.JSON(200, gin.H{
		"message": datas,
	})
}

func GetAboutMeById(c *gin.Context) {
	client := connect_firebase.Connection()
	defer client.Close()

	id := c.Param("id")

	// get data from firebase
	docs, err := client.Collection("about_me").Documents(c).GetAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	var about_me models.AboutMe
	for _, doc := range docs {
		if doc.Ref.ID == id {
			// generate id from firebase
			about_me.ID = doc.Ref.ID
			about_me.CreatedAt = doc.CreateTime
			about_me.UpdatedAt = doc.UpdateTime
			doc.DataTo(&about_me)
		}
	}

	c.JSON(200, gin.H{
		"message": about_me,
	})
}
