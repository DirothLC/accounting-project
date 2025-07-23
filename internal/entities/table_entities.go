package entities

import (
	"github.com/google/uuid"
	"time"
)

type Terminal struct {
	ID           int64  `gorm:"primaryKey"`
	ClientID     string `gorm:"unique"`
	ClientSecret string
	UUID         uuid.UUID
}

type TransactionStatus struct {
	Code        string `gorm:"primaryKey"`
	Description string
}

type Transaction struct {
	ID            int64 `gorm:"primaryKey"`
	TerminalID    int64
	OrderID       string
	Amount        float64
	Status        string
	CreatedAt     time.Time
	StatusChanged time.Time
	Code          string
	Message       *string

	Terminal  Terminal          `gorm:"foreignKey:TerminalID"`
	StatusRef TransactionStatus `gorm:"foreignKey:Status;references:Code"`
}
