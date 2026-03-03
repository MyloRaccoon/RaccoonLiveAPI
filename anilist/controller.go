package anilist

import (
	"encoding/json"
	"net/http"
	"os"
)

func ControllerGetLastActivity(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserID(os.Getenv("ANILIST_USERNAME"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	activity, err := GetLastActivity(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activity)
}

func ControllerGetFavoriteAnimes(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("ANILIST_USERNAME")
	animes, err := GetFavoritesAnime(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animes)
}

func ControllerGetFavoriteMangas(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("ANILIST_USERNAME")
	mangas, err := GetFavoritesManga(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mangas)
}

func ControllerGetFavoriteCharacters(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("ANILIST_USERNAME")
	characters, err := GetFavoritesCharacters(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(characters)
}

func ControllerGetFavoriteStaff(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("ANILIST_USERNAME")
	staff, err := GetFavoritesStaff(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(staff)
}

func ControllerGetFavoriteStudio(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("ANILIST_USERNAME")
	studios, err := GetFavoritesStudio(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studios)
}