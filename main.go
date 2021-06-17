package main

import (
	"github.com/gin-gonic/gin"
	"github.com/razasayed/golang_rest_api_boilerplate/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.Run()
}
