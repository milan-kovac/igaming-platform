package tournament

import (
	"errors"
	"igaming-platform/internal/database"

	"gorm.io/gorm"
)

func DistributePrizes(tournamentId int) error {
	var tournament Tournament

	if err := database.DB.First(&tournament, tournamentId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Tournament not found.")
		}
		return err
	}

	if err := database.DB.Exec("CALL distribute_prizes(?)", tournamentId).Error; err != nil {
		return err
	}

	return nil
}
