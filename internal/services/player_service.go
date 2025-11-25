package services

import (
	"igaming-platform/internal/domain"
	"igaming-platform/internal/repositories"
)

type IPlayerService interface {
	GetPlayersRanking() ([]domain.PlayerRanking, error)
}

type playerService struct {
	repository repositories.IPlayerRepository
}

func NewPlayerService(repository repositories.IPlayerRepository) IPlayerService {
	return &playerService{repository: repository}
}

func (service *playerService) GetPlayersRanking() ([]domain.PlayerRanking, error) {
	return service.repository.GetPlayersRanking()
}
