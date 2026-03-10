package anilist

import (
	"encoding/json"
	"net/http"
	"os"
)

func GetProfileController(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("ANILIST_USERNAME")
	profile, err := getProfile(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func GetLastActivityController(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserID(os.Getenv("ANILIST_USERNAME"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	activity, err := getLastActivity(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activity)
}

func GetFavoriteAnimesController(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("ANILIST_USERNAME")
	animes, err := getFavoritesAnime(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animes)
}

func GetFavoriteMangasController(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("ANILIST_USERNAME")
	mangas, err := getFavoritesManga(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mangas)
}

func GetFavoriteCharactersController(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("ANILIST_USERNAME")
	characters, err := getFavoritesCharacters(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(characters)
}

func GetFavoriteStaffController(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("ANILIST_USERNAME")
	staff, err := getFavoritesStaff(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(staff)
}

func GetFavoriteStudioController(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("ANILIST_USERNAME")
	studios, err := getFavoritesStudio(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studios)
}