package main

import (
	"newsfeeder/models"

	"github.com/gin-gonic/gin"

	"net/http"
)

// GET /plants
// Get all plants
func FindPlants(c *gin.Context) {
	var plants []models.Plant
	models.DB.Find(&plants)

	c.JSON(http.StatusOK, gin.H{"data": plants})
}
