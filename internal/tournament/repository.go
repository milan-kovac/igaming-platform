package tournament

import "github.com/stretchr/testify/mock"

type MockTournamentRepository struct {
	mock.Mock
}

func (repository *tournamentRepository) GetByID(id int64) (*Tournament, error) {
	var t Tournament
	return &t, repository.db.First(&t, id).Error
}

func (repository *tournamentRepository) DistributePrizes(id int64) error {
	return repository.db.Exec("CALL distribute_prizes(?)", id).Error
}
