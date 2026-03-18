package music

import (
	"encoding/json"
	"net/http"
	"raccoonlive-api/logger"

	"github.com/gorilla/mux"
)

func GetMusicsController(w http.ResponseWriter, r *http.Request) {
	musics, err := getMusics()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Log("ERROR - GET Musics: " + err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(musics)
	logger.Log("SUCCESS - GET Musics")
}

func GetMusicByIDController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	music, err := getMusicById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Log("ERROR - GET Music: " + err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(music)
	logger.Log("SUCCESS - GET Music")
}

func PutMusicController(w http.ResponseWriter, r *http.Request) {
	var m Music
	json.NewDecoder(r.Body).Decode(&m)

	err := putMusic(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Log("ERROR - PUT Music: " + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	logger.Log("SUCCESS - PUT Music")
}

func DeleteMusicController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	music, err := deleteMusicById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Log("ERROR - DELETE Music: " + err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(music)
	logger.Log("SUCCESS - DELETE Music")
}

func PatchMusicController(w http.ResponseWriter, r *http.Request) {
	var m Music
	json.NewDecoder(r.Body).Decode(&m)

	err := patchMusic(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Log("ERROR - PATCH Music: " + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	logger.Log("SUCCESS - PATCH Music")
}