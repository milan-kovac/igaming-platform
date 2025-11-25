package tests

import (
	"igaming-platform/internal/domain"

	"github.com/stretchr/testify/mock"
)

type MockPlayerRepository struct {
	mock.Mock
}

func (m *MockPlayerRepository) GetPlayersRanking() ([]domain.PlayerRanking, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]domain.PlayerRanking), args.Error(1)
	}
	return nil, args.Error(1)
}
