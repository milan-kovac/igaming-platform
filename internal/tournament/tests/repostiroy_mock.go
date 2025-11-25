package tournament_test

import (
	"igaming-platform/internal/tournament"

	"github.com/stretchr/testify/mock"
)

type MockTournamentRepository struct {
	mock.Mock
}

func (m *MockTournamentRepository) GetByID(id int64) (*tournament.Tournament, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*tournament.Tournament), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTournamentRepository) DistributePrizes(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}
