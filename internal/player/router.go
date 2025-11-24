package player

import "github.com/gorilla/mux"

func RegiterPlayersRouter(router *mux.Router) {
	tournamentRouter := router.PathPrefix("/players").Subrouter()

	tournamentRouter.HandleFunc("/ranking", PlayersRankingHandler).Methods("GET")
}
