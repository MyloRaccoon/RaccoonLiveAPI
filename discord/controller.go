package discord

import (
	"encoding/json"
	"net/http"
	"raccoonlive-api/logger"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	Status.MU.RLock()
	defer Status.MU.RUnlock()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(&Status)
	logger.Log("SUCCESS - GET discord")
}