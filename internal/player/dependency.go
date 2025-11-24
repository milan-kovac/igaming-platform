package player

import (
	"net/http"

	"gorm.io/gorm"
)

// repository
type IPlayerRepository interface {
	GetPlayersRanking() ([]PlayerRanking, error)
}

type playerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) IPlayerRepository {
	return &playerRepository{db: db}
}

// service

type IPlayerService interface {
	GetPlayersRanking() ([]PlayerRanking, error)
}

type playerService struct {
	repository IPlayerRepository
}

func NewPlayerService(repository IPlayerRepository) IPlayerService {
	return &playerService{repository: repository}
}

// handler

type IPlayerHandler interface {
	PlayersRankingHandler(w http.ResponseWriter, r *http.Request)
}

type playerHandler struct {
	service IPlayerService
}

func NewPlayerHandler(service IPlayerService) IPlayerHandler {
	return &playerHandler{service: service}
}
