package anilist

import (
	"encoding/json"
	"net/http"
	"os"
	"raccoonlive-api/logger"

)

func GetProfileController(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("ANILIST_USERNAME")
	profile, err := getProfile(username)
	if err != nil {
		logger.Log("Error GET Anilist Profile: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
	logger.Log("SUCCESS - GET Anilist Profile")
}

func GetLastActivityController(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserID(os.Getenv("ANILIST_USERNAME"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	activity, err := getLastActivity(userId)
	if err != nil {
		logger.Log("Error GET Anilist Activity: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activity)
	logger.Log("SUCCESS - GET Anilist Activity")
}

func GetFavoriteAnimesController(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("ANILIST_USERNAME")
	animes, err := getFavoritesAnime(username)
	if err != nil {
		logger.Log("Error GET Anilist Fav Animes: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animes)
	logger.Log("SUCCESS - GET Anilist Fav Animes")
}

func GetFavoriteMangasController(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("ANILIST_USERNAME")
	mangas, err := getFavoritesManga(username)
	if err != nil {
		logger.Log("Error GET Anilist Fav Mangas: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mangas)
	logger.Log("SUCCESS - GET Anilist Fav Mangas")
}

func GetFavoriteCharactersController(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("ANILIST_USERNAME")
	characters, err := getFavoritesCharacters(username)
	if err != nil {
		logger.Log("Error GET Anilist Fav Characters: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(characters)
	logger.Log("SUCCESS - GET Anilist Fav Characters")
}

func GetFavoriteStaffController(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("ANILIST_USERNAME")
	staff, err := getFavoritesStaff(username)
	if err != nil {
		logger.Log("Error GET Anilist Fav Staff " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(staff)
	logger.Log("SUCCESS - GET Anilist Fav Saff")
}

func GetFavoriteStudiosController(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("ANILIST_USERNAME")
	studios, err := getFavoritesStudios(username)
	if err != nil {
		logger.Log("Error GET Anilist Fav Studios " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studios)
	logger.Log("SUCCESS - GET Anilist Fav Studios")
}