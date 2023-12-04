package main

import (
	"backend/api/about_me"
	"backend/api/certificates"
	"backend/api/socmed"
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
	router.GET("/about_me/:id", about_me.GetAboutMeById)
	router.POST("/about_me", about_me.PostAboutMe)
	router.PUT("/about_me/:id", about_me.EditAboutMe)
	router.DELETE("/about_me/:id", about_me.DeleteAboutMe)

	router.GET("/certificates", certificates.GetCertificates)
	router.GET("/certificates/:id", certificates.GetCertificateById)
	router.POST("/certificates", certificates.PostCertificate)
	router.PUT("/certificates/:id", certificates.EditCertificate)
	router.DELETE("/certificates/:id", certificates.DeleteCertificate)

	router.GET("/socmed", socmed.GeSocmed)
	router.GET("/socmed/:id", socmed.GetSocmedById)
	router.POST("/socmed", socmed.PostSocmed)
	router.PUT("/socmed/:id", socmed.EditSocmed)
	router.DELETE("/socmed/:id", socmed.DeleteSocmed)

	router.Run("localhost:8080")
}
