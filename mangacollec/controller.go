package mangacollec

import (
	"encoding/json"
	"net/http"
	"raccoonlive-api/logger"
)

func GetSeriesController(w http.ResponseWriter, r *http.Request) {
	data, err := fetchData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Log("ERROR - GET MangaCollec Series")
		return
	}
	series := data.Series
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(series)
	logger.Log("SUCCESS - GET MangaCollec Series")
}

func GetEditionsController(w http.ResponseWriter, r *http.Request) {
	data, err := fetchData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Log("ERROR - GET MangaCollec Editions")
		return
	}
	editions := data.Editions
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(editions)
	logger.Log("SUCCESS - GET MangaCollec Editions")
}

func GetVolumesController(w http.ResponseWriter, r *http.Request) {
	data, err := fetchData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Log("ERROR - GET MangaCollec Volumes")
		return
	}
	volumes := data.Volumes
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(volumes)
	logger.Log("SUCCESS - GET MangaCollec Volumes")
}

func GetBoxEditionsController(w http.ResponseWriter, r *http.Request) {
	data, err := fetchData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Log("ERROR - GET MangaCollec BoxEditions")
		return
	}
	boxEditions := data.BoxEditions
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(boxEditions)
	logger.Log("SUCCESS - GET MangaCollec BoxEditions")
}

func GetBoxesController(w http.ResponseWriter, r *http.Request) {
	data, err := fetchData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Log("ERROR - GET MangaCollec Boxes")
		return
	}
	boxes := data.Boxes
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(boxes)
	logger.Log("SUCCESS - GET MangaCollec Boxes")
}

func GetBoxVolumesController(w http.ResponseWriter, r *http.Request) {
	data, err := fetchData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Log("ERROR - GET MangaCollec BoxVolumes")
		return
	}
	boxVolumes := data.BoxVolumes
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(boxVolumes)
	logger.Log("SUCCESS - GET MangaCollec BoxVolumes")
}

func GetVolumesInPossessionsController(w http.ResponseWriter, r *http.Request) {
	volumes, err := getVolumesInPossession()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Log("ERROR - GET MangaCollec VolumesInPossession")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(volumes)
	logger.Log("SUCCESS - GET MangaCollec VolumesInPossession")
}