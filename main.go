package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"raccoonlive-api/anilist"
	"raccoonlive-api/discord"
	"raccoonlive-api/github"
	"raccoonlive-api/music"
	"raccoonlive-api/youtube"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	bot := discord.Run()
	if bot == nil {
		return
	}
	defer bot.Close()

	router := mux.NewRouter()
	
	router.HandleFunc("/discord", func(w http.ResponseWriter, r *http.Request) {
		discord.Status.MU.RLock()
		defer discord.Status.MU.RUnlock()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&discord.Status)
	}).Methods("GET")

	router.HandleFunc("/anilist", func(w http.ResponseWriter, r *http.Request) {
		userId, err := anilist.GetUserID(os.Getenv("ANILIST_USERNAME"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		activity, err := anilist.GetLastActivity(userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(activity)
	}).Methods("GET")

	router.HandleFunc("/anilist/animes", func(w http.ResponseWriter, r *http.Request) {
		username := os.Getenv("ANILIST_USERNAME")
		animes, err := anilist.GetFavoritesAnime(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(animes)
	}).Methods("GET")

	router.HandleFunc("/anilist/mangas", func(w http.ResponseWriter, r *http.Request) {
		username := os.Getenv("ANILIST_USERNAME")
		mangas, err := anilist.GetFavoritesManga(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mangas)
	}).Methods("GET")

	router.HandleFunc("/anilist/characters", func(w http.ResponseWriter, r *http.Request) {
		username := os.Getenv("ANILIST_USERNAME")
		characters, err := anilist.GetFavoritesCharacters(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(characters)
	}).Methods("GET")

	router.HandleFunc("/anilist/staff", func(w http.ResponseWriter, r *http.Request) {
		username := os.Getenv("ANILIST_USERNAME")
		staff, err := anilist.GetFavoritesStaff(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(staff)
	}).Methods("GET")

	router.HandleFunc("/anilist/studios", func(w http.ResponseWriter, r *http.Request) {
		username := os.Getenv("ANILIST_USERNAME")
		studios, err := anilist.GetFavoritesStudio(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(studios)
	}).Methods("GET")

	router.HandleFunc("/youtube", func(w http.ResponseWriter, r *http.Request) {
		video, err := youtube.GetLastVideo()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(video)
	}).Methods("GET")

	router.HandleFunc("/github/profile", func(w http.ResponseWriter, r *http.Request) {
		profile, err := github.GetUser(os.Getenv("GITHUB_LOGIN"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(profile)
	}).Methods("GET")

	router.HandleFunc("/github/repo", func(w http.ResponseWriter, r *http.Request) {
		profile, err := github.GetLastRepo(os.Getenv("GITHUB_LOGIN"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(profile)
	}).Methods("GET")

	router.HandleFunc("/musics", func(w http.ResponseWriter, r *http.Request) {
		musics, err := music.GetMusics()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(musics)
	}).Methods("GET")

	router.HandleFunc("/music/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		music, err := music.GetMusicById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(music)
	}).Methods("GET")

	router.HandleFunc("/music", func(w http.ResponseWriter, r *http.Request) {
		var m music.Music
		json.NewDecoder(r.Body).Decode(&m)

		err := music.PostMusic(m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	server := &http.Server{Addr: ":8080", Handler: router}

	go server.ListenAndServe()
	fmt.Println("Server started")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	fmt.Println("Stopping...")
	server.Shutdown(context.Background())
}