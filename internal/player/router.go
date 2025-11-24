package player

import (
	"igaming-platform/internal/database"

	"github.com/gorilla/mux"
)

func RegiterPlayersRouter(router *mux.Router) {
	repository := NewPlayerRepository(database.GetDB())
	service := NewPlayerService(repository)
	handler := NewPlayerHandler(service)

	tournamentRouter := router.PathPrefix("/players").Subrouter()

	tournamentRouter.HandleFunc("/ranking", handler.PlayersRankingHandler).Methods("GET")
}
