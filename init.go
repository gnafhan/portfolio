package main

import (
	"backend/api/about_me"
	connect_firebase "backend/firebase"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		// post demo data to firebase
		client := connect_firebase.Connection()
		defer client.Close()

		_, _, err := client.Collection("about_me").Add(context.Background(), map[string]interface{}{
			"name":        "Nguyen Van A",
			"title":       "Software Engineer",
			"description": "I am a software engineer",
			"skills":      []string{"Golang", "Python", "Java"},
		})

		if err != nil {
			log.Fatalf("Failed adding alovelace: %v", err)
		}

		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	router.GET("/about_me", about_me.GetAboutMe)
	router.POST("/about_me", about_me.PostAboutMe)
	router.GET("/about_me/:id", about_me.GetAboutMeById)
	router.Run("localhost:8080")
}