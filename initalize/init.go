package initalize

import (
	"io/ioutil"

	"github.com/morgulbrut/colorlog"
)

func Map() {
	colorlog.Debug("writing map.toml")
	err := ioutil.WriteFile("map.toml", mapToml(), 0644)
	if err != nil {
		colorlog.Fatal("%s", err.Error())
	}
	colorlog.Debug("writing template.html")
	err = ioutil.WriteFile("template.html", templateHtml(), 0644)
	if err != nil {
		colorlog.Fatal("%s", err.Error())
	}
	colorlog.Debug("writing data.csv")
	err = ioutil.WriteFile("data.csv", dataCsv(), 0644)
	if err != nil {
		colorlog.Fatal("%s", err.Error())
	}

}

func mapToml() []byte {
	return []byte(`Name="Test"
Style="toner" # terrain, watercolor, toner 
MarkerColor="darkgreen" # red, blue, green, purple, orange, darkred, lightred, beige, darkblue, darkgreen, cadetblue, darkpurple, white, pink, lightblue, lightgreen, gray, black, lightgray
IconColor="black"
Zoom=8
MinZoom=6
MaxZoom=18
Symbols=[
    #Klasse:fontawesome symbol
	"Shop:shopping-basket",
	"Event:calendar",
	"Restaurant:cutlery",
]
Center="London"
`)
}

func dataCsv() []byte {
	return []byte(`name, address, url, desc, opening_times, type
Reformhaus Seefeld,"Seefeldstrasse 187, 8008 Zürich",http://www.reformhaus-seefeld.ch,Ihre Bio-Oase in Zürich,Mo-Fr 08.00-18.30 / Sa 08.00-16.00,Shop`)
}

func templateHtml() []byte {
	return []byte(`<!DOCTYPE html>
	<head>
	<meta http-equiv="content-type" content="text/html; charset=UTF-8" />
	<script>L_PREFER_CANVAS=false; L_NO_TOUCH=false; L_DISABLE_3D=false;</script>
	<script src="https://cdn.jsdelivr.net/npm/leaflet@1.2.0/dist/leaflet.js"></script>
	<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js"></script>
	<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/Leaflet.awesome-markers/2.0.2/leaflet.awesome-markers.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/leaflet.markercluster/1.1.0/leaflet.markercluster.js"></script>
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/leaflet@1.2.0/dist/leaflet.css"/>
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta.2/css/bootstrap.min.css"/>
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta.2/css/bootstrap-theme.min.css"/>
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css"/>
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/Leaflet.awesome-markers/2.0.2/leaflet.awesome-markers.css"/>
	<link rel="stylesheet" href="https://rawgit.com/python-visualization/folium/master/folium/templates/leaflet.awesome.rotate.css"/>
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/leaflet.markercluster/1.1.0/MarkerCluster.css"/>
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/leaflet.markercluster/1.1.0/MarkerCluster.Default.css"/>
	<style>html, body {width: 100%;height: 100%;margin: 0;padding: 0;}</style>
	<style>#map {position: relative; width: 100.0%; height: 100.0%; left: 0.0%;top: 0.0%;}</style>
	<style>body {padding-top: 65px;}</style>
	<style>h1 {color: #718123; font-family: 'Kaushan Script', cursive;}</style>
	</head>
	<body>
	<link href="https://fonts.googleapis.com/css?family=Kaushan+Script&display=swap" rel="stylesheet">
	<!-- Navigation -->
	<nav class="navbar navbar-expand-lg navbar-dark bg-dark fixed-top">
	  <div class="container">
		<h1>
		  Test Map
		</h1>
		<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
		  <span class="navbar-toggler-icon">
		  </span>
		</button>
		<div class="collapse navbar-collapse" id="navbarResponsive">
		  <ul class="navbar-nav ml-auto">
			<li class="nav-item">
			  <button type="button" class="btn btn-sm align-middle btn-outline-secondary"  data-toggle="modal" data-target="#legendModal">
				Legende
			  </button>
			</li>
		  </ul>
		</div>
	  </div>
	</nav>
	<!-- Trigger the modal with a button -->
	<!-- Modal -->
	{{.Legend}}
	{{.Div}}
	</div>
	</body>
	<script>
	{{.Content}}	
	</script>
	`)
}
