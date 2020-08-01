package structs

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"googlemaps.github.io/maps"
	"io"
	"time"
)

type DistanceStruct struct {
	Origin      maps.LatLng `json:"origin" bson:"origin"`
	Destination maps.LatLng `json:"destination" bson:"origin"`
}

func (ds *DistanceStruct) JsonToObject(r io.Reader) error {
	return json.NewDecoder(r).Decode(ds)
}

func (ds *DistanceStruct) LogToDB(client *mongo.Client) {

	database := client.Database("log")
	collection := database.Collection("distance_request_log")
	mongoContext, _ := context.WithTimeout(context.Background(), 1*time.Second)

	_, _ = collection.InsertOne(mongoContext, bson.D{
		{Key: "origin", Value: ds.Origin},
		{Key: "destination", Value: ds.Destination},
	})

}
