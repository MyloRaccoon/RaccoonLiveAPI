package youtube

import (
	"encoding/json"
	"net/http"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	video, err := getLastVideo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(video)
}