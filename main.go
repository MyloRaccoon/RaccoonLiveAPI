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

func apiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-Api-Key")
		if key != os.Getenv("API_KEY") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

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
	
	router.HandleFunc("/discord", discord.Controller).Methods("GET")

	router.HandleFunc("/anilist", anilist.ControllerGetLastActivity).Methods("GET")

	router.HandleFunc("/anilist/animes", anilist.ControllerGetFavoriteAnimes).Methods("GET")

	router.HandleFunc("/anilist/mangas", anilist.ControllerGetFavoriteMangas).Methods("GET")

	router.HandleFunc("/anilist/characters", anilist.ControllerGetFavoriteCharacters).Methods("GET")

	router.HandleFunc("/anilist/staff", anilist.ControllerGetFavoriteStaff).Methods("GET")

	router.HandleFunc("/anilist/studios", anilist.ControllerGetFavoriteStudio).Methods("GET")

	router.HandleFunc("/youtube", youtube.Controller).Methods("GET")

	router.HandleFunc("/github/profile", github.ControllerProfile).Methods("GET")

	router.HandleFunc("/github/repo", github.ControllerRepo).Methods("GET")

	router.HandleFunc("/musics", music.ControllerGetMusics).Methods("GET")

	router.HandleFunc("/music/{id}", music.ControllerGetMusicByID).Methods("GET")

	router.Handle("/music", 
		apiKeyMiddleware(http.HandlerFunc(music.ControllerPostMusic))).
		Methods("POST")

	router.Handle("/music/{id}", 
		apiKeyMiddleware(http.HandlerFunc(music.ControllerDeleteMusic))).
		Methods("DELETE")

	server := &http.Server{Addr: ":8080", Handler: router}

	go server.ListenAndServe()
	fmt.Println("Server started")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	fmt.Println("Stopping...")
	server.Shutdown(context.Background())
}