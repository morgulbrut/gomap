package leaflet

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/morgulbrut/colorlog"
)

// Icon represents an marker icon
// MarkerColor can be "red", "blue", "green", "purple", "orange", "darkred", "lightred", "beige", "darkblue",
//"darkgreen", "cadetblue", "darkpurple", "white", "pink", "lightblue", "lightgreen", "gray", "black", "lightgray" or "random"
type Icon struct {
	Symbol      string
	IconColor   string
	MarkerColor string
}

// Popup represents a  marker popup
type Popup struct {
	Title   string
	Link    string
	Address string
	Times   string
	Desc    string
}

// Popup represents a fontawsome marker on the map
type Marker struct {
	Name     string
	Position Coordinate
	Icon     Icon
	Popup    Popup
}

// Syms is a map representing a collection of fontawesome symbols mapped to keywords
var Syms map[string]string

var Colors = []string{"red", "blue", "green", "purple", "orange", "darkred", "lightred", "beige", "darkblue",
	"darkgreen", "cadetblue", "darkpurple", "white", "pink", "lightblue", "lightgreen", "gray", "black", "lightgray"}

// MarkerString returns a chunk of JS representing a marker
func MarkerString(m Marker, ma Mapdata) []byte {
	var marker strings.Builder

	marker.WriteString("var marker = L.marker([{{.Position.Lat}}, {{.Position.Lon}}],{icon: new L.Icon.Default()}).addTo(marker_cluster);")
	marker.WriteString("var icon = L.AwesomeMarkers.icon({icon: '{{.Icon.Symbol}}', iconColor: '{{.Icon.IconColor}}', markerColor: '{{.Icon.MarkerColor}}', prefix: 'fa', extraClasses: 'fa-rotate-0'});")
	marker.WriteString("marker.setIcon(icon);")

	marker.WriteString("var popup = L.popup({maxWidth: '300'});")
	marker.WriteString(`var html = $('<div id="html" style="width: 100.0%%; height: 100.0%%;"><a target="_blank" href="{{.Popup.Link}}"><b>{{.Popup.Title}}</b></a><br/>{{.Popup.Desc}}<br/><br/><b>{{.Popup.Address}}</b><br/><br/>{{.Popup.Times}}</div>')[0];`)
	marker.WriteString("popup.setContent(html);")
	marker.WriteString("marker.bindPopup(popup);")

	tmpl, err := template.New("marker").Parse(marker.String())
	if err != nil {
		colorlog.Fatal(err.Error())
	}
	var ret bytes.Buffer
	err = tmpl.Execute(&ret, m)
	if err != nil {
		colorlog.Fatal(err.Error())
	}

	return ret.Bytes()
}

// Adds a marker to a map
func AddMarker(m Marker, ma *Mapdata) {
	ma.Markers = append(ma.Markers, m)
}
