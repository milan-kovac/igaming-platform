package repositories

import (
	"igaming-platform/internal/domain"

	"gorm.io/gorm"
)

type playerRepository struct {
	db *gorm.DB
}

type IPlayerRepository interface {
	GetPlayersRanking() ([]domain.PlayerRanking, error)
}

func NewPlayerRepository(db *gorm.DB) IPlayerRepository {
	return &playerRepository{db: db}
}

func (repository *playerRepository) GetPlayersRanking() ([]domain.PlayerRanking, error) {
	var rankings []domain.PlayerRanking

	return rankings, repository.db.Raw(`
       WITH ranked_players AS (
            SELECT
                id,
                name,
                account_balance,
                RANK() OVER (ORDER BY account_balance DESC) AS player_rank
            FROM players
        )
        SELECT *
        FROM ranked_players
`).Scan(&rankings).Error
}
