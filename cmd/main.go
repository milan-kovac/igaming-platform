package main

import (
	"igaming-platform/internal/config"
	"igaming-platform/internal/database"
	"igaming-platform/internal/player"
	"igaming-platform/internal/tournament"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "igaming-platform/cmd/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title iGaming Platform API
// @version 1.0
// @description API for managing players, tournaments and prizes.
func main() {
	env := config.LoadEnv()

	// database connection
	database.Initialize()

	router := mux.NewRouter()

	// tournaments router
	tournament.RegiterTournamentsRouter(router)

	// playerss router
	player.RegiterPlayersRouter(router)

	// swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler())

	err := http.ListenAndServe(":"+env.Server.Port, router)
	if err != nil {
		log.Println("Server failed to start:", err)
	}

	// close database connection
	database.CloseConnection()
}
