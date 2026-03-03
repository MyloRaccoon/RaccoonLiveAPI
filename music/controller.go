package music

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func ControllerGetMusics(w http.ResponseWriter, r *http.Request) {
	musics, err := getMusics()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(musics)
}

func ControllerGetMusicByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	music, err := getMusicById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(music)
}

func ControllerPostMusic(w http.ResponseWriter, r *http.Request) {
	var m Music
	json.NewDecoder(r.Body).Decode(&m)

	err := postMusic(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func ControllerDeleteMusic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	music, err := deleteMusicById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(music)
}