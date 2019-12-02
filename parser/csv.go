package parser

import (
	"os"
	"strconv"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/morgulbrut/colorlog"
	"github.com/morgulbrut/gomap/geocode"
	"github.com/morgulbrut/gomap/leaflet"
)

type Entry struct { // Our example struct, you can use "-" to ignore a field
	Name    string `csv:"name"`
	Address string `csv:"address"`
	URL     string `csv:"url"`
	Desc    string `csv:"desc"`
	Times   string `csv:"opening_times"`
	Type    string `csv:"type"`
}

func ReadCSV(path string, data *leaflet.Mapdata) []leaflet.Marker {
	dataFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		colorlog.Fatal("%s", err.Error())
	}
	defer dataFile.Close()

	colorlog.Debug("Generating map icons")
	for _, s := range data.Symbols {
		sym := strings.Split(s, ":")
		leaflet.Syms[sym[0]] = sym[1]
	}

	colorlog.Info("Reading %s", path)
	var markers []leaflet.Marker
	var mr leaflet.Marker
	var pos leaflet.Coordinate
	var indx = 0

	entries := []*Entry{}

	if err := gocsv.UnmarshalFile(dataFile, &entries); err != nil { // Load clients from file
		colorlog.Fatal("%s", err.Error())
	}

	for _, entry := range entries {
		colorlog.Debug("Adding marker for %s", entry.Name)
		pos = geocode.GetCoord(entry.Address)

		mr = leaflet.Marker{
			Name:     strconv.Itoa(indx),
			Position: pos,
			Icon: leaflet.Icon{
				Symbol:      leaflet.Syms[entry.Type],
				IconColor:   data.IconColor,
				MarkerColor: data.MarkerColor,
			},
			Popup: leaflet.Popup{
				Link:    strings.ReplaceAll(entry.URL, "'", `\'`),
				Title:   strings.ReplaceAll(entry.Name, "'", `\'`),
				Address: strings.ReplaceAll(entry.Address, "'", `\'`),
				Times:   strings.ReplaceAll(entry.Times, "'", `\'`),
				Desc:    strings.ReplaceAll(entry.Desc, "'", `\'`),
			},
		}
		markers = append(markers, mr)
	}
	return markers
}
