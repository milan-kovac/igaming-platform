package tournament

import (
	"igaming-platform/internal/common"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// DistributePrizeHandler distributes prizes for a given tournament
// @Summary Distribute prizes
// @Description Distribute prizes based on tournament ranking
// @Tags Tournaments
// @Param id path int true "Tournament ID"
// @Success 200 {object} common.SuccessResponse{data=nil} "Prizes successfully distributed"
// @Failure 400 {object} common.ErrorResponse "Bad request"
// @Router /tournaments/{id}/distribute [post]
func (handler *tournamentHandler) DistributePrizesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	tournamentId, parseErr := strconv.Atoi(vars["id"])
	if parseErr != nil {
		common.Error(w, http.StatusBadRequest, "Invalid tournament ID", "ID must be a number")
		return
	}

	distributeErr := handler.service.DistributePrizes(tournamentId)
	if distributeErr != nil {
		common.Error(w, http.StatusNotFound, "Failed to distribute prizes", distributeErr.Error())
		return
	}

	common.Success(w, "Prizes successfully distributed.", nil)
}
