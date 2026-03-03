package github

import (
	"encoding/json"
	"net/http"
	"os"
)

func ControllerProfile(w http.ResponseWriter, r *http.Request) {
	profile, err := GetUser(os.Getenv("GITHUB_LOGIN"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func ControllerRepo(w http.ResponseWriter, r *http.Request) {
	profile, err := GetLastRepo(os.Getenv("GITHUB_LOGIN"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}