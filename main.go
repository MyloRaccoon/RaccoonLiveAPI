package main

import (
	"context"
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
	fmt.Println("Starting server...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	bot, err := discord.Run()
	if err != nil {
		fmt.Println("Error starting Discord bot: ", err)
		return
	}
	defer bot.Close()

	router := mux.NewRouter()
	
	router.HandleFunc("/discord", discord.Controller).Methods("GET")

	router.HandleFunc("/anilist", anilist.GetLastActivityController).Methods("GET")

	router.HandleFunc("/anilist/animes", anilist.GetFavoriteAnimesController).Methods("GET")

	router.HandleFunc("/anilist/mangas", anilist.GetFavoriteMangasController).Methods("GET")

	router.HandleFunc("/anilist/characters", anilist.GetFavoriteCharactersController).Methods("GET")

	router.HandleFunc("/anilist/staff", anilist.GetFavoriteStaffController).Methods("GET")

	router.HandleFunc("/anilist/studios", anilist.GetFavoriteStudioController).Methods("GET")

	router.HandleFunc("/youtube", youtube.Controller).Methods("GET")

	router.HandleFunc("/github/profile", github.ProfileController).Methods("GET")

	router.HandleFunc("/github/repo", github.RepoController).Methods("GET")

	router.HandleFunc("/musics", music.GetMusicsController).Methods("GET")

	router.HandleFunc("/music/{id}", music.GetMusicByIDController).Methods("GET")

	router.Handle("/music", 
		apiKeyMiddleware(http.HandlerFunc(music.PutMusicController))).
		Methods("PUT")

	router.Handle("/music/{id}", 
		apiKeyMiddleware(http.HandlerFunc(music.DeleteMusicController))).
		Methods("DELETE")

	router.Handle("/music/{id}",
		apiKeyMiddleware(http.HandlerFunc(music.PatchMusicController))).
		Methods("PATCH")

	server := &http.Server{Addr: ":8080", Handler: router}

	go server.ListenAndServe()
	fmt.Println("Server started")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	fmt.Println("Stopping...")
	server.Shutdown(context.Background())
}