package player

func (repository *playerRepository) GetPlayersRanking() ([]PlayerRanking, error) {
	var rankings []PlayerRanking

	return rankings, repository.db.Raw(`
    SELECT
        RANK() OVER (ORDER BY account_balance DESC) AS player_rank,
        id,
        name,
        account_balance
    FROM players
`).Scan(&rankings).Error
}
