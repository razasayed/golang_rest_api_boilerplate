package main

import (
	"github.com/gin-gonic/gin"
	"github.com/razasayed/golang_rest_api_boilerplate/controllers"
	"github.com/razasayed/golang_rest_api_boilerplate/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/movies", controllers.FindMovies)
	r.POST("/movies", controllers.CreateMovie)
	r.GET("/movies/:id", controllers.FindMovie)
	r.PATCH("movies/:id", controllers.UpdateMovie)
	r.DELETE("movies/:id", controllers.DeleteMovie)

	r.Run()
}
