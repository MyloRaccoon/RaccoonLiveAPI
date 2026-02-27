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
	})

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
	})

	router.HandleFunc("/anilist/anime", func(w http.ResponseWriter, r *http.Request) {
		username := os.Getenv("ANILIST_USERNAME")
		animes, err := anilist.GetFavoritesAnime(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(animes)
	})

	router.HandleFunc("/youtube", func(w http.ResponseWriter, r *http.Request) {
		video, err := youtube.GetLastVideo()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(video)
	})

	router.HandleFunc("/github/profile", func(w http.ResponseWriter, r *http.Request) {
		profile, err := github.GetUser(os.Getenv("GITHUB_LOGIN"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(profile)
	})

	router.HandleFunc("/github/repo", func(w http.ResponseWriter, r *http.Request) {
		profile, err := github.GetLastRepo(os.Getenv("GITHUB_LOGIN"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(profile)
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