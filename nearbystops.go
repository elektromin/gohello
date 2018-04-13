package main

import (
	"net/http"
	"encoding/json"
	"strconv"
	"os"
	"fmt"
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

var apiKey = os.Getenv("NEARBYSTOPS_KEY")
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

	Info.Println("http response: ", resp.Status)
	if resp.StatusCode != 200 {
		return NearbyStops{}, fmt.Errorf("Response: ", resp.Status)
	}

	var d NearbyStops
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return NearbyStops{}, err
	}

	Info.Println("response: ", d)

	return d, nil
}
