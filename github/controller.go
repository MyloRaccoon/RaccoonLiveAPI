package github

import (
	"encoding/json"
	"net/http"
	"os"
)

func ProfileController(w http.ResponseWriter, r *http.Request) {
	profile, err := getUser(os.Getenv("GITHUB_LOGIN"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func RepoController(w http.ResponseWriter, r *http.Request) {
	profile, err := getLastRepo(os.Getenv("GITHUB_LOGIN"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}