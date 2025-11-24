package player

import "fmt"

func (service *playerService) GetPlayersRanking() ([]PlayerRanking, error) {
	rankings, err := service.repository.GetPlayersRanking()
	if err != nil {
		return nil, fmt.Errorf("Failed to get players ranking: %w", err)
	}
	return rankings, nil
}
