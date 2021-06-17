package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/razasayed/golang_rest_api_boilerplate/models"
)

type CreateMovieInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateMovieInput struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// GET /movies
// Get all movies
func FindMovies(c *gin.Context) {
	var movies []models.Movie
	models.DB.Find(&movies)

	c.JSON(http.StatusOK, gin.H{"data": movies})
}

// POST /movies
// Create a new movie
func CreateMovie(c *gin.Context) {
	// Validate user input
	var input CreateMovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create movie
	movie := models.Movie{Title: input.Title, Description: input.Description}
	models.DB.Create(&movie)

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// GET /movies/:id
// Find a movie
func FindMovie(c *gin.Context) {
	var movie models.Movie
	if err := models.DB.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// PATCH /movies/:id
// Update a movie
func UpdateMovie(c *gin.Context) {
	// Get model if exist
	var movie models.Movie
	if err := models.DB.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate user input
	var input UpdateMovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&movie).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// DELETE /movies/:id
// Delete a movie
func DeleteMovie(c *gin.Context) {
	// Get model if exist
	var movie models.Movie
	if err := models.DB.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&movie)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
