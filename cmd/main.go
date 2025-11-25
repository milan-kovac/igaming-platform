package main

import (
	"igaming-platform/internal/api/routers"
	"igaming-platform/internal/config"
	"igaming-platform/internal/database"
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

	// database
	database.Initialize(env)
	defer database.Close()

	router := mux.NewRouter()

	// tournaments router
	routers.RegiterTournamentsRouter(router)

	// playerss router
	routers.RegiterPlayersRouter(router)

	// swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler())

	err := http.ListenAndServe(":"+env.Server.Port, router)
	if err != nil {
		log.Println("Server failed to start:", err)
	}

}
