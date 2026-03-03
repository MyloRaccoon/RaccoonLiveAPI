package discord

import (
	"encoding/json"
	"net/http"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	Status.MU.RLock()
	defer Status.MU.RUnlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Status)
}