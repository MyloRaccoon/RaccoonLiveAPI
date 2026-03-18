package github

import (
	"encoding/json"
	"net/http"
	"os"
	"raccoonlive-api/logger"
)

func ProfileController(w http.ResponseWriter, r *http.Request) {
	profile, err := getUser(os.Getenv("GITHUB_LOGIN"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Log("ERROR - GET Github Profile: " + err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
	logger.Log("SUCCESS - GET Github Profile")
}

func RepoController(w http.ResponseWriter, r *http.Request) {
	profile, err := getLastRepo(os.Getenv("GITHUB_LOGIN"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Log("ERROR - GET Github Repo: " + err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
	logger.Log("SUCCESS - GET Github Repo")
}