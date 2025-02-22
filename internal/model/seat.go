package model

import (
	"time"

	"gorm.io/gorm"
)

type Seat struct {
	ID         uint      `gorm:"primary_key"`
	TheaterID  uint      `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	HallID     uint      `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsReserved bool      `gorm:"not null"`
	Price      float64   `gorm:"not null"`
	Reserve    []Reserve `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Deleted    gorm.DeletedAt
}

func NewSeat(theaterID, hallID uint, isReserved bool, price float64) *Seat {
	return &Seat{
		TheaterID:  theaterID,
		HallID:     hallID,
		IsReserved: isReserved,
		Price:      price,
	}
}
