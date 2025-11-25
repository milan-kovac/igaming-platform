package tests

import (
	"errors"
	"igaming-platform/internal/domain"
	"igaming-platform/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPlayersRanking_Success(t *testing.T) {
	mockRepo := new(MockPlayerRepository)
	service := services.NewPlayerService(mockRepo)

	mockRankings := []domain.PlayerRanking{
		{PlayerRank: 1, ID: 10, Name: "Alice", AccountBalance: 500},
		{PlayerRank: 2, ID: 11, Name: "Bob", AccountBalance: 300},
	}

	mockRepo.On("GetPlayersRanking").Return(mockRankings, nil)

	rankings, err := service.GetPlayersRanking()
	assert.NoError(t, err)
	assert.Equal(t, mockRankings, rankings)

	mockRepo.AssertExpectations(t)
}

func TestGetPlayersRanking_Error(t *testing.T) {
	mockRepo := new(MockPlayerRepository)
	service := services.NewPlayerService(mockRepo)

	mockRepo.On("GetPlayersRanking").Return(nil, errors.New("Failed to get players ranking."))

	rankings, err := service.GetPlayersRanking()
	assert.Nil(t, rankings)
	assert.EqualError(t, err, "Failed to get players ranking.")

	mockRepo.AssertExpectations(t)
}
