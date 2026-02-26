package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"

	"raccoonlive-api/discord"
)

func main() {
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

	server := &http.Server{Addr: ":8080", Handler: router}

	go server.ListenAndServe()
	fmt.Println("Server started")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	fmt.Println("Stopping...")
	server.Shutdown(context.Background())
}