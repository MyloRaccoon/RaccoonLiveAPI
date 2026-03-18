package youtube

import (
	"encoding/json"
	"net/http"
	"raccoonlive-api/logger"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	video, err := getLastVideo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Log("ERROR - GET Youtube: " + err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(video)
	logger.Log("SUCCESS - GET Youtube")
}