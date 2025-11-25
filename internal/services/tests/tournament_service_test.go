package tests

import (
	"errors"
	"igaming-platform/internal/domain"
	"igaming-platform/internal/services"
	"igaming-platform/internal/shared/fallacies"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributePrizes_Success(t *testing.T) {
	mockRepo := &MockTournamentRepository{}

	service := services.NewTournamentService(mockRepo)

	tourn := &domain.Tournament{ID: 10}

	mockRepo.On("GetByID", int64(10)).Return(tourn, nil)
	mockRepo.On("DistributePrizes", int64(10)).Return(nil)

	err := service.DistributePrizes(10)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDistributePrizes_TournamentNotFound(t *testing.T) {
	mockRepo := &MockTournamentRepository{}

	service := services.NewTournamentService(mockRepo)

	mockRepo.On("GetByID", int64(10)).Return(nil, errors.New(fallacies.TournamentNotFound.Error()))

	err := service.DistributePrizes(10)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), fallacies.TournamentNotFound.Error())
	mockRepo.AssertExpectations(t)
}

func TestDistributePrizes_DistributeError(t *testing.T) {
	mockRepo := &MockTournamentRepository{}

	service := services.NewTournamentService(mockRepo)

	tourn := &domain.Tournament{ID: 10}

	mockRepo.On("GetByID", int64(10)).Return(tourn, nil)
	mockRepo.On("DistributePrizes", int64(10)).Return(errors.New("db error"))

	err := service.DistributePrizes(10)

	assert.Error(t, err)
	assert.Equal(t, "db error", err.Error())
	mockRepo.AssertExpectations(t)
}
