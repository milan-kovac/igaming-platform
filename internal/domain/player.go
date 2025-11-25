package domain

type Player struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	AccountBalance float64 `json:"account_balance"`
}

type PlayerRanking struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	AccountBalance float64 `json:"account_balance"`
	PlayerRank     int     `json:"player_rank"`
}
