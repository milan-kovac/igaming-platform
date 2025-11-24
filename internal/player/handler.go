package player

import (
	"igaming-platform/internal/common"
	"net/http"
)

// PlayersRankingHandler returns the global ranking of all players
// @Summary Get global player ranking
// @Description Retrieve a ranking list of all players based on their account balances.
// @Tags Players
// @Success 200 {object} common.SuccessResponse{data=[]PlayerRanking} "Player ranking fetched successfully"
// @Failure 500 {object} common.ErrorResponse "Failed to fetch player ranking"
// @Router /players/ranking [get]
func (handler *playerHandler) PlayersRankingHandler(w http.ResponseWriter, r *http.Request) {
	rankings, err := handler.service.GetPlayersRanking()
	if err != nil {
		common.Error(w, http.StatusInternalServerError, "Failed to fetch tournament ranking", err.Error())
		return
	}

	common.Success(w, "Tournament ranking fetched successfully.", rankings)
}
