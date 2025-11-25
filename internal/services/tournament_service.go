package services

import (
	"errors"
	"igaming-platform/internal/repositories"
	"igaming-platform/internal/shared/fallacies"

	"gorm.io/gorm"
)

type tournamentService struct {
	repository repositories.ITournamentRepository
}

type ITournamentService interface {
	DistributePrizes(tournamentId int) error
}

func NewTournamentService(repository repositories.ITournamentRepository) ITournamentService {
	return &tournamentService{repository}
}

func (service *tournamentService) DistributePrizes(tournamentId int) error {
	tournament, err := service.repository.GetByID(int64(tournamentId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fallacies.TournamentNotFound
		}

		return err
	}

	if tournament.PrizesDistributed {
		return fallacies.PrizesAlreadyDistributed
	}

	if err := service.repository.DistributePrizes(tournament.ID); err != nil {
		return err
	}

	return nil
}
