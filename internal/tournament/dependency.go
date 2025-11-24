package tournament

import (
	"net/http"

	"gorm.io/gorm"
)

// repository
type tournamentRepository struct {
	db *gorm.DB
}

type ITournamentRepository interface {
	GetByID(id int64) (*Tournament, error)
	DistributePrizes(id int64) error
}

func NewTournamentRepository(db *gorm.DB) ITournamentRepository {
	return &tournamentRepository{db: db}
}

// service
type tournamentService struct {
	repository ITournamentRepository
}

type ITournamentService interface {
	DistributePrizes(tournamentId int) error
}

func NewTournamentService(repository ITournamentRepository) ITournamentService {
	return &tournamentService{repository}
}

// handler
type ITournamentHandler interface {
	DistributePrizesHandler(w http.ResponseWriter, r *http.Request)
}

type tournamentHandler struct {
	service ITournamentService
}

func NewTournamentHandler(tournamentService ITournamentService) ITournamentHandler {
	return &tournamentHandler{tournamentService}
}
