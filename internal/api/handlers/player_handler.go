package handlers

import (
	"igaming-platform/internal/services"
	"igaming-platform/internal/shared/responses"
	"net/http"
)

type playerHandler struct {
	service services.IPlayerService
}

type IPlayerHandler interface {
	PlayersRankingHandler(w http.ResponseWriter, r *http.Request)
}

func NewPlayerHandler(service services.IPlayerService) IPlayerHandler {
	return &playerHandler{service: service}
}

// PlayersRankingHandler returns the global ranking of all players
// @Summary Get global player ranking
// @Description Retrieve a ranking list of all players based on their account balances.
// @Tags Players
// @Success 200 {object} responses.SuccessResponse{data=[]domain.PlayerRanking} "Player ranking fetched successfully"
// @Failure 500 {object} responses.ErrorResponse "Failed to fetch player ranking"
// @Router /players/ranking [get]
func (handler *playerHandler) PlayersRankingHandler(w http.ResponseWriter, r *http.Request) {
	rankings, err := handler.service.GetPlayersRanking()
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Success(w, "Tournament ranking fetched successfully.", rankings)
}
