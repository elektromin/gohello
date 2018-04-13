package main

import (
	"strconv"
	"net/http"
	"encoding/json"
)

type Station struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Distance  int64 `json:"distance"`
}

func toStation(sl StopLocation) (Station, error) {
	lat, err := strconv.ParseFloat(sl.Lat, 64)
	if err != nil {
		Error.Println("error: ", err)
		return Station{}, err
	}

	long, err := strconv.ParseFloat(sl.Lon, 64)
	if err != nil {
		Error.Println("error: ", err)
		return Station{}, err
	}

	dist, err := strconv.ParseInt(sl.Dist, 10, 64)
	if err != nil {
		Error.Println("error: ", err)
		return Station{}, err
	}

	return Station{
		Id:        sl.Id,
		Name:      sl.Name,
		Latitude:  lat,
		Longitude: long,
		Distance:  dist,
	}, nil
}

func closestStation(w http.ResponseWriter, r *http.Request) {
	lat, err := strconv.ParseFloat(r.URL.Query().Get("latitude"), 64)
	if err != nil {
		Error.Println("error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	long, err := strconv.ParseFloat(r.URL.Query().Get("longitude"), 64)
	if err != nil {
		Error.Println("error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ns, err := nearbyStops(lat, long)
	if err != nil {
		Error.Println("error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(ns.LocationList.StopLocations) == 0 {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	station, err := toStation(ns.LocationList.StopLocations[0])
	if err != nil {
		Error.Println("error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Info.Println("response: ", station)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(station)
}
