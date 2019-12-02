package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/morgulbrut/colorlog"
	"github.com/morgulbrut/gomap/geocode"
	"github.com/morgulbrut/gomap/initalize"
	"github.com/morgulbrut/toml"

	"github.com/akamensky/argparse"
	"github.com/morgulbrut/color256"
	"github.com/morgulbrut/gomap/leaflet"
	"github.com/morgulbrut/gomap/parser"
	"github.com/morgulbrut/gomap/site"
)

func logo() {
	color256.Init()

	color256.PrintRandom("\n  ██████╗  ██████╗ ███╗   ███╗ █████╗ ██████╗")
	color256.PrintRandom(" ██╔════╝ ██╔═══██╗████╗ ████║██╔══██╗██╔══██╗")
	color256.PrintRandom(" ██║  ███╗██║   ██║██╔████╔██║███████║██████╔╝")
	color256.PrintRandom(" ██║   ██║██║   ██║██║╚██╔╝██║██╔══██║██╔═══╝")
	color256.PrintRandom(" ╚██████╔╝╚██████╔╝██║ ╚═╝ ██║██║  ██║██║")
	color256.PrintRandom("  ╚═════╝  ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝ \n")
}

func readMapConfig() leaflet.Mapdata {
	var conf leaflet.Mapdata
	_, err := toml.DecodeFile("map.toml", &conf)
	if err != nil {
		colorlog.Fatal(err.Error())
		os.Exit(1)
	}
	conf.CenterCoords = geocode.GetCoord(conf.Center)
	return conf
}

func main() {
	logo()
	pr := argparse.NewParser("goMap", "Generates a map from a given input file")
	v := pr.Flag("v", "verbose", &argparse.Options{Required: false, Help: "Set verbosity"})
	i := pr.Flag("i", "init", &argparse.Options{Required: false, Help: "Initializes the a map project"})
	err := pr.Parse(os.Args)
	if err != nil {
		fmt.Print(pr.Usage(err))
	}
	if *v {
		colorlog.SetLogLevel(colorlog.DEBUG)
	} else {
		colorlog.SetLogLevel(colorlog.INFO)
	}
	if *i {
		initalize.Map()
		os.Exit(0)
	}

	leaflet.Syms = make(map[string]string)
	colorlog.Info("Reading config")
	mainMap := readMapConfig()
	markers := parser.ReadCSV("data.csv", &mainMap)

	for _, m := range markers {
		leaflet.AddMarker(m, &mainMap)
	}

	var s site.Site
	s.Legend = GenerateLegend()
	s.Div = `<div class="folium-map" id="map">`
	s.Content = leaflet.GenerateMap(mainMap)
	s.Name = mainMap.Name

	site.GenerateSite("index.html", s)
}

func GenerateLegend() string {
	var sb strings.Builder
	sb.WriteString(`<div id="legendModal" class="modal fade" role="dialog"><div class="modal-dialog"><div class="modal-content"><div class="modal-header"><h1 class="modal-title">Legende</h1></div><div class="modal-body">`)
	keys := reflect.ValueOf(leaflet.Syms).MapKeys()
	for _, k := range keys {
		sb.WriteString(`<p><i class="fa fa-` + leaflet.Syms[k.String()] + `" aria-hidden="true"></i>` + k.String() + `</p>`)
	}
	sb.WriteString(`</div><div class="modal-footer"><button type="button" class="btn btn-default" data-dismiss="modal">Schliessen</button></div></div></div></div>`)
	return sb.String()
}
