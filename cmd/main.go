package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type MappingProbabilities struct {
	Clay            float64 `json:"Clay"`
	Clay_loam       float64 `json:"Clay Loam"`
	Loam            float64 `json:"Loam"`
	Loamy_sand      float64 `json:"Loamy Sand"`
	Sand            float64 `json:"Sand"`
	Sandy_clay      float64 `json:"Sandy Clay"`
	Sandy_clay_loam float64 `json:"Sandy Clay Loam"`
	Sandy_loam      float64 `json:"Sandy Loam"`
	Silt            float64 `json:"Silt"`
	Silt_loam       float64 `json:"Silt Loam"`
	Silty_clay      float64 `json:"Silty Clay"`
	Silty_clay_loam float64 `json:"Silty Clay Loam"`
}

type SoilMapping struct {
	Red           int                  `json:"Red"`
	Green         int                  `json:"Green"`
	Blue          int                  `json:"Blue"`
	Description   string               `json:"Description"`
	Probabilities MappingProbabilities `json:"Probabilities"`
}

func loadSoilMappings(path string) []SoilMapping {
	var soilmappings []SoilMapping
	mappingFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	byteArray, _ := ioutil.ReadAll(mappingFile)
	json.Unmarshal(byteArray, &soilmappings)
	defer mappingFile.Close()
	return soilmappings
}

func main() {
	soilmappings := loadSoilMappings("assets/soil_mapping.json")
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/soil", func(c *gin.Context) {
		c.JSON(http.StatusOK, soilmappings)
	})
	router.Run()
}
