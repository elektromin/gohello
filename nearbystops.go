package main

import (
	"net/http"
	"encoding/json"
	"strconv"
)

type NearbyStops struct {
	LocationList struct {
		StopLocations []StopLocation `json:"StopLocation"`
	}
}

type StopLocation struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Lat  string `json:"lat"`
	Lon  string `json:"lon"`
	Dist string `json:"dist"`
}

const apiKey = "28b7e5e14d3049e8afa45f48aa8630a8"
const maxresults = 5
const radius = 2000

func nearbyStops(lat float64, long float64) (NearbyStops, error) {
	url := "http://api.sl.se/api2/nearbystops.json?key=" + apiKey +
		"&maxresults=" + strconv.Itoa(maxresults) +
		"&originCoordLat=" + strconv.FormatFloat(lat, 'f', -1, 64) +
		"&originCoordLong=" + strconv.FormatFloat(long, 'f', -1, 64) +
		"&radius=" + strconv.Itoa(radius)
	Info.Println("url: " + url)

	resp, err := http.Get(url)
	if err != nil {
		return NearbyStops{}, err
	}

	defer resp.Body.Close()

	var d NearbyStops

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return NearbyStops{}, err
	}

	Info.Println("response: ", d)

	return d, nil
}
