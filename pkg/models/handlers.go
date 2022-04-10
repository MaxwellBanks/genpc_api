package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

//Logs and continues in case of nonfatal error
//TODO: Replace with more robust logging
func handeNonFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getSoilMappings(path string) []SoilMapping {
	var soilmappings []SoilMapping
	mappingFile, err := os.Open(path)
	defer mappingFile.Close()
	handleNonFatal(err)
	byteArray, _ := ioutil.ReadAll(mappingFile)
	json.Unmarshal(byteArray, &soilmappings)
	return soilmappings
}
