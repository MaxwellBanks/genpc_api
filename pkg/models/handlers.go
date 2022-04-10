package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

//Logs and continues in case of nonfatal error
//TODO: Replace with more robust logging
func handleNonFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetSoilMappings(path string) []SoilMapping {
	var soilmappings []SoilMapping
	mappingFile, err := os.Open(path)
	defer mappingFile.Close()
	handleNonFatal(err)
	byteArray, _ := ioutil.ReadAll(mappingFile)
	json.Unmarshal(byteArray, &soilmappings)
	return soilmappings
}
