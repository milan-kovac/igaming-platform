package domain

import "time"

type Tournament struct {
	ID                int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name              string    `json:"name"`
	PrizePool         float64   `json:"prize_pool"`
	StartDate         time.Time `json:"start_date"`
	EndDate           time.Time `json:"end_date"`
	PrizesDistributed bool      `json:"prizes_distributed" gorm:"default:false"`
}
