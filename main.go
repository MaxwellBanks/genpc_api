package main

import (
	"net/http"

	"github.com/MaxwellBanks/genpc_api/pkg/models"
	"github.com/gin-gonic/gin"
)

func main() {
	soilmappings := models.GetSoilMappings("pkg/models/soil_mapping.json")
	router := gin.New()

	router.GET("/soil", func(c *gin.Context) {
		c.JSON(http.StatusOK, soilmappings)
	})
	router.Run()
}
