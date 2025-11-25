package routers

import (
	"igaming-platform/internal/api/handlers"
	"igaming-platform/internal/database"
	"igaming-platform/internal/repositories"
	"igaming-platform/internal/services"

	"github.com/gorilla/mux"
)

func RegiterPlayersRouter(router *mux.Router) {
	repository := repositories.NewPlayerRepository(database.GetDB())
	service := services.NewPlayerService(repository)
	handler := handlers.NewPlayerHandler(service)

	tournamentRouter := router.PathPrefix("/players").Subrouter()

	tournamentRouter.HandleFunc("/ranking", handler.PlayersRankingHandler).Methods("GET")
}
