package projects

import (
	connect_firebase "backend/firebase"
	models "backend/model"
	"log"

	"github.com/gin-gonic/gin"
)

func GetProject(c *gin.Context) {
	// get data
	client := connect_firebase.Connection()
	defer client.Close()

	docs, err := client.Collection("projects").Documents(c).GetAll()
	if err != nil {
		log.Println(err)
	}

	datas := []models.Project{}

	var projects models.Project

	for _, doc := range docs {
		// generate id from firebase
		projects.ID = doc.Ref.ID
		projects.CreatedAt = doc.CreateTime
		projects.UpdatedAt = doc.UpdateTime
		doc.DataTo(&projects)
		datas = append(datas, projects)
	}

	c.JSON(200, gin.H{
		"message": datas,
	})
}

func GetProjectById(c *gin.Context) {
	client := connect_firebase.Connection()
	defer client.Close()

	id := c.Param("id")
	is_any := false

	// get data from firebase
	docs, err := client.Collection("projects").Documents(c).GetAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}

	var projects models.Project
	for _, doc := range docs {
		if doc.Ref.ID == id {
			// generate id from firebase
			is_any = true
			projects.ID = doc.Ref.ID
			projects.CreatedAt = doc.CreateTime
			projects.UpdatedAt = doc.UpdateTime
			doc.DataTo(&projects)
		}
	}

	if is_any == false {
		c.JSON(404, gin.H{
			"message": "Not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": projects,
	})
}
