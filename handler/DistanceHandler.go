package handler

import (
	"context"
	"encoding/json"
	"googlemaps.github.io/maps"
	"log"
	"main/structs"
	"net/http"
	"os"
)

type DistanceHandler struct {
}

func NewDistanceHandler() *DistanceHandler {
	return &DistanceHandler{}
}

func (d *DistanceHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")

	if request.Method == http.MethodPost {
		d.getDistance(writer, request)
		return
	}

	writer.WriteHeader(http.StatusMethodNotAllowed)
}

func (d *DistanceHandler) getDistance(writer http.ResponseWriter, request *http.Request) {

	type mainResponse struct {
		OriginAddresses      string `json:"origin_addresses"`
		DestinationAddresses string `json:"destination_addresses"`
		DistanceKm           string `json:"distance_km"`
		DistanceMeters       int    `json:"distance_meters"`
		Success              bool   `json:"success"`
	}

	ds := structs.DistanceStruct{}
	err := ds.JsonToObject(request.Body)

	if err != nil {
		http.Error(writer, "Origin or Destination is wrong", http.StatusBadRequest)
		return
	}

	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	r := &maps.DistanceMatrixRequest{
		Origins:      []string{ds.Origin.String()},
		Destinations: []string{ds.Destination.String()},
	}

	distanceResponseResult, err := c.DistanceMatrix(context.Background(), r)

	finalResponse := mainResponse{}

	if err != nil || distanceResponseResult.Rows[0].Elements[0].Status == "ZERO_RESULTS" {
		writer.WriteHeader(http.StatusBadRequest)
		finalResponse.Success = false
	} else {
		writer.WriteHeader(http.StatusOK)
		finalResponse.OriginAddresses = distanceResponseResult.OriginAddresses[0]
		finalResponse.DestinationAddresses = distanceResponseResult.DestinationAddresses[0]
		finalResponse.DistanceKm = distanceResponseResult.Rows[0].Elements[0].Distance.HumanReadable
		finalResponse.DistanceMeters = distanceResponseResult.Rows[0].Elements[0].Distance.Meters
		finalResponse.Success = true
	}

	mainResponseJson, _ := json.Marshal(finalResponse)
	_, _ = writer.Write(mainResponseJson)
}
