package tournament_test

import (
	"errors"
	"igaming-platform/internal/tournament"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributePrizes_Success(t *testing.T) {
	mockRepo := &MockTournamentRepository{}

	service := tournament.NewTournamentService(mockRepo)

	tourn := &tournament.Tournament{ID: 10}

	mockRepo.On("GetByID", int64(10)).Return(tourn, nil)
	mockRepo.On("DistributePrizes", int64(10)).Return(nil)

	err := service.DistributePrizes(10)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDistributePrizes_TournamentNotFound(t *testing.T) {
	mockRepo := &MockTournamentRepository{}

	service := tournament.NewTournamentService(mockRepo)

	mockRepo.On("GetByID", int64(10)).Return(nil, errors.New("not found"))

	err := service.DistributePrizes(10)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Tournament not found")
	mockRepo.AssertExpectations(t)
}

func TestDistributePrizes_DistributeError(t *testing.T) {
	mockRepo := &MockTournamentRepository{}

	service := tournament.NewTournamentService(mockRepo)

	tourn := &tournament.Tournament{ID: 10}

	mockRepo.On("GetByID", int64(10)).Return(tourn, nil)
	mockRepo.On("DistributePrizes", int64(10)).Return(errors.New("db error"))

	err := service.DistributePrizes(10)

	assert.Error(t, err)
	assert.Equal(t, "db error", err.Error())
	mockRepo.AssertExpectations(t)
}
