package model

import (
	"time"
)

type Trnx struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Status     int       `json:"status"`
	Amount     float64   `json:"amount"`
	Vendor     string    `json:"vendor"`
	Created_at time.Time `gorm:"autoCreateTime"`
}
