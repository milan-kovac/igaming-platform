package tournament

import "fmt"

func (service *tournamentService) DistributePrizes(tournamentId int) error {
	tournament, err := service.repository.GetByID(int64(tournamentId))
	if err != nil {
		return fmt.Errorf("Tournament not found: %w", err)
	}

	if err := service.repository.DistributePrizes(tournament.ID); err != nil {
		return err
	}

	return nil
}
