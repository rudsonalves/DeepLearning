package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// citiBikeURL provides the station statuses of CitiBike bike sharing stations
const citiBikeURL = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"

// stationData is used to unmarshal the JSON document returned from citiBikeURL
type stationData struct {
	LastUpdated int `json:"last_updated"`
	TTL         int `json:"ttl"`
	Data        struct {
		Stations []station `json:"stations"`
	} `json:"data"`
}

// station is used to unmarshal each of the station documents in stationData
type station struct {
	ID                string `json:"station_id"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumBikesDisabled  int    `json:"num_bike_disabled"`
	NumDocksAvailable int    `json:"num_docks_available"`
	NumDocksDisabled  int    `json:"num_docks_disabled"`
	IsInstalled       int    `json:"is_installed"`
	IsRenting         int    `json:"is_renting"`
	IsReturning       int    `json:"is_returning"`
	LastReported      int    `json:"last_reported"`
	HasAvailableKeys  bool   `json:"eightd_has_available_keys"`
}

func main() {
	os.Chdir("/home/rudson/Documents/Estudos/DeepLearning/json/citibike")

	// Get the JSON response from URL
	response, err := http.Get(citiBikeURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Read the body of the response into []byte
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Declare a variable of type stationData
	var sd stationData

	// Unmarshal the JSON data into the variable
	if err := json.Unmarshal(body, &sd); err != nil {
		log.Fatal(err)
	}

	// Print the first station
	fmt.Printf("%+v\n\n", sd.Data.Stations[0])

	// Marshal the data
	outputData, err := json.Marshal(sd)
	if err != nil {
		log.Fatal(err)
	}

	// Save the marshalled data to a file
	if err := ioutil.WriteFile("citibike.json", outputData, 0644); err != nil {
		log.Fatal(err)
	}
}
