package tournament

import "github.com/gorilla/mux"

func RegiterTournamentRouter(router *mux.Router) {
	tournamentRouter := router.PathPrefix("/tournaments").Subrouter()

	tournamentRouter.HandleFunc("/{id}/distribute", DistributePrizesHandler).Methods("POST")
}
