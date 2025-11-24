package player_test

import (
	"igaming-platform/internal/player"

	"github.com/stretchr/testify/mock"
)

type MockPlayerRepository struct {
	mock.Mock
}

func (m *MockPlayerRepository) GetPlayersRanking() ([]player.PlayerRanking, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]player.PlayerRanking), args.Error(1)
	}
	return nil, args.Error(1)
}
