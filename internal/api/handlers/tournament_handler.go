package handlers

import (
	"igaming-platform/internal/services"
	"igaming-platform/internal/shared/responses"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type tournamentHandler struct {
	service services.ITournamentService
}

type ITournamentHandler interface {
	DistributePrizesHandler(w http.ResponseWriter, r *http.Request)
}

func NewTournamentHandler(tournamentService services.ITournamentService) ITournamentHandler {
	return &tournamentHandler{tournamentService}
}

// DistributePrizeHandler distributes prizes for a given tournament
// @Summary Distribute prizes
// @Description Distribute prizes based on tournament ranking
// @Tags Tournaments
// @Param id path int true "Tournament id"
// @Success 200 {object} responses.SuccessResponse{data=nil} "Prizes successfully distributed"
// @Failure 400 {object} responses.ErrorResponse "Invalid tournament ID"
// @Failure 404 {object} responses.ErrorResponse "Tournament not found"
// @Failure 409 {object} responses.ErrorResponse "Prizes already distributed"
// @Router /tournaments/{id}/distribute [post]
func (handler *tournamentHandler) DistributePrizesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	tournamentId, parseErr := strconv.Atoi(vars["id"])
	if parseErr != nil {
		responses.Error(w, parseErr)
		return
	}

	distributeErr := handler.service.DistributePrizes(tournamentId)
	if distributeErr != nil {
		responses.Error(w, distributeErr)
		return
	}

	responses.Success(w, "Prizes successfully distributed.", nil)
}
