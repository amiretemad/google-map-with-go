package structs

import (
	"encoding/json"
	"googlemaps.github.io/maps"
	"io"
)

type DistanceStruct struct {
	Origin      maps.LatLng `json:"origin"`
	Destination maps.LatLng `json:"destination"`
}

func (ds *DistanceStruct) JsonToObject(r io.Reader) error {
	return json.NewDecoder(r).Decode(ds)
}
