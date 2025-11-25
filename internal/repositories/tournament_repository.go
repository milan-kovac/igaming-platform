package repositories

import (
	"igaming-platform/internal/domain"

	"gorm.io/gorm"
)

type tournamentRepository struct {
	db *gorm.DB
}

type ITournamentRepository interface {
	GetByID(id int64) (*domain.Tournament, error)
	DistributePrizes(id int64) error
}

func NewTournamentRepository(db *gorm.DB) ITournamentRepository {
	return &tournamentRepository{db: db}
}

func (repository *tournamentRepository) GetByID(id int64) (*domain.Tournament, error) {
	var tournament domain.Tournament
	result := repository.db.Raw("SELECT * FROM tournaments WHERE id = ?", id).Scan(&tournament)
	if err := result.Error; err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &tournament, nil
}

func (repository *tournamentRepository) DistributePrizes(id int64) error {
	return repository.db.Exec("CALL distribute_prizes(?)", id).Error
}
