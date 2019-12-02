package geocode

import (
	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/morgulbrut/gomap/leaflet"
)

var geocoder = openstreetmap.Geocoder()

// GetCoord returns a leaflet.Coordinate of a given address
func GetCoord(addr string) leaflet.Coordinate {
	loc, _ := geocoder.Geocode(addr)
	return leaflet.Coordinate{Lat: loc.Lat, Lon: loc.Lng}
}
