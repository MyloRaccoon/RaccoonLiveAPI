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
		json.NewEncoder(w).Encode(&discord.Status)
	})

	router.HandleFunc("/anilist", func(w http.ResponseWriter, r *http.Request) {
		userId, err := anilist.GetUserID(os.Getenv("ANILIST_USERNAME"))
		if err != nil {
			return
		}
		activity, err := anilist.GetLastActivity(userId)
		if err != nil {
			return
		}
		json.NewEncoder(w).Encode(activity)
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