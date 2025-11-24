package tournament

import (
	"igaming-platform/internal/database"

	"github.com/gorilla/mux"
)

func RegiterTournamentsRouter(router *mux.Router) {
	repository := NewTournamentRepository(database.GetDB())
	service := NewTournamentService(repository)
	handler := NewTournamentHandler(service)

	tournamentRouter := router.PathPrefix("/tournaments").Subrouter()

	tournamentRouter.HandleFunc("/{id}/distribute", handler.DistributePrizesHandler).Methods("POST")
}
