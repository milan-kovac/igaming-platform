package player

import (
	"fmt"
	"igaming-platform/internal/database"
)

func GetPlayersRanking() ([]PlayerRanking, error) {
	var rankings []PlayerRanking

	err := database.DB.Raw(`
		SELECT
			RANK() OVER (ORDER BY account_balance DESC) AS player_rank,
			id,
			name,
			account_balance
		FROM players
	`).Scan(&rankings).Error

	if err != nil {
		return nil, fmt.Errorf("Failed to fetch player ranking: %w", err)
	}

	return rankings, nil
}
