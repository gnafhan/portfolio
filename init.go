package main

import (
	"backend/api/about_me"
	"backend/api/certificates"
	"backend/api/journey"
	"backend/api/projects"
	"backend/api/socmed"
	"backend/api/technologies"
	"backend/middleware"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // You may replace "*" with specific origins
	router.Use(cors.New(config))
	router.Use(middleware.ApiKey)
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

	router.GET("/socmed", socmed.GetSocmed)
	router.GET("/socmed/:id", socmed.GetSocmedById)
	router.POST("/socmed", socmed.PostSocmed)
	router.PUT("/socmed/:id", socmed.EditSocmed)
	router.DELETE("/socmed/:id", socmed.DeleteSocmed)

	router.GET("/projects", projects.GetProject)
	router.GET("/projects/:id", projects.GetProjectById)
	router.POST("/projects", projects.PostProject)
	router.PUT("/projects/:id", projects.EditProject)
	router.DELETE("/projects/:id", projects.DeleteProject)

	router.GET("/journey", journey.GetJourney)
	router.GET("/journey/:id", journey.GetJourneyById)
	router.POST("/journey", journey.PostJourney)
	router.PUT("/journey/:id", journey.EditJourney)
	router.DELETE("/journey/:id", journey.DeleteJourney)

	router.GET("/technologies", technologies.GetTechnology)
	router.GET("/technologies/:id", technologies.GetTechnologyById)
	router.POST("/technologies", technologies.PostTechnology)
	router.PUT("/technologies/:id", technologies.EditTechnology)
	router.DELETE("/technologies/:id", technologies.DeleteTechnology)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)

}
