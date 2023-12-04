package socmed

import (
	connect_firebase "backend/firebase"
	models "backend/model"
	"log"

	"github.com/gin-gonic/gin"
)

func GeSocmed(c *gin.Context) {
	// get data
	client := connect_firebase.Connection()
	defer client.Close()

	docs, err := client.Collection("socmed").Documents(c).GetAll()
	if err != nil {
		log.Println(err)
	}

	datas := []models.SocialMedia{}

	var socmed models.SocialMedia

	for _, doc := range docs {
		// generate id from firebase
		socmed.ID = doc.Ref.ID
		socmed.CreatedAt = doc.CreateTime
		socmed.UpdatedAt = doc.UpdateTime
		doc.DataTo(&socmed)
		datas = append(datas, socmed)
	}

	c.JSON(200, gin.H{
		"message": datas,
	})
}

func GetSocmedById(c *gin.Context) {
	client := connect_firebase.Connection()
	defer client.Close()

	id := c.Param("id")
	is_any := false

	// get data from firebase
	docs, err := client.Collection("socmed").Documents(c).GetAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	var socmed models.SocialMedia
	for _, doc := range docs {
		if doc.Ref.ID == id {
			// generate id from firebase
			is_any = true
			socmed.ID = doc.Ref.ID
			socmed.CreatedAt = doc.CreateTime
			socmed.UpdatedAt = doc.UpdateTime
			doc.DataTo(&socmed)
		}
	}

	if is_any == false {
		c.JSON(404, gin.H{
			"message": "Not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": socmed,
	})
}
