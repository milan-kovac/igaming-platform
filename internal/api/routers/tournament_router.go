package routers

import (
	"igaming-platform/internal/api/handlers"
	"igaming-platform/internal/database"
	"igaming-platform/internal/repositories"
	"igaming-platform/internal/services"

	"github.com/gorilla/mux"
)

func RegiterTournamentsRouter(router *mux.Router) {
	repository := repositories.NewTournamentRepository(database.GetDB())
	service := services.NewTournamentService(repository)
	handler := handlers.NewTournamentHandler(service)

	tournamentRouter := router.PathPrefix("/tournaments").Subrouter()

	tournamentRouter.HandleFunc("/{id}/distribute", handler.DistributePrizesHandler).Methods("POST")
}
