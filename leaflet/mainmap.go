package leaflet

import (
	"bytes"
	"text/template"
)

// Coordinate represents a coordinate in 41.32409, 7.90039 format
type Coordinate struct {
	Lat float64
	Lon float64
}

// Mapdata represents the settings a map
// Style can be "toner", "watercolor" or "terrain"
type Mapdata struct {
	Center       string
	CenterCoords Coordinate
	Style        string
	Name         string
	Zoom         int
	MaxZoom      int
	MinZoom      int
	Markers      []Marker
	IconColor    string
	MarkerColor  string
	Symbols      []string
}

// MapString returns a chunk of JS defining a map without markers
func MapString(data Mapdata) []byte {

	mapTemplate := `
	var bounds = null;
	var map= L.map(
		'map', {
		center: [{{.CenterCoords.Lat}}, {{.CenterCoords.Lon}}],
		zoom: {{.Zoom}},
		maxBounds: bounds,
		layers: [],
		worldCopyJump: false,
		crs: L.CRS.EPSG3857,
		zoomControl: true,
		});

	var tile_layer= L.tileLayer(
		'https://stamen-tiles-{s}.a.ssl.fastly.net/{{.Style}}/{z}/{x}/{y}.png', // toner, watercolor, terrain
		{
		"attribution": null,
		"detectRetina": false,
		"maxNativeZoom": {{.MaxZoom}},
		"maxZoom": {{.MaxZoom}},
		"minZoom": {{.MinZoom}},
		"noWrap": false,
		"subdomains": "abc"
	}).addTo(map);

	var marker_cluster = L.markerClusterGroup({});
	map.addLayer(marker_cluster);`

	tmpl, err := template.New("map").Parse(mapTemplate)
	if err != nil {
		panic(err)
	}
	var ret bytes.Buffer
	err = tmpl.Execute(&ret, data)
	if err != nil {
		panic(err)
	}
	return ret.Bytes()
}

// GenerateMap returns a chunk of JS defining and it's marker
// Usually that's the function you want to call
func GenerateMap(data Mapdata) string {
	m := MapString(data)
	for _, marker := range data.Markers {
		m = append(m, MarkerString(marker, data)...)
	}
	return string(m)
}
