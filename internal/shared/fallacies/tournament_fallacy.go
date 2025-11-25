package fallacies

import "errors"

var (
	TournamentNotFound       = errors.New("Tournament not found.")
	PrizesAlreadyDistributed = errors.New("Prizes already distributed")
)
