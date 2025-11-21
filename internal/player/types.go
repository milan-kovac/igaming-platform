package player

type Player struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	AccountBalance float64 `json:"account_balance"`
}
