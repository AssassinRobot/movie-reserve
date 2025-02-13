package model

import "time"

type Hall struct {
	ID          uint `gorm:"primary_key"`
	TheaterID   uint `gorm:"not null"`
	SeatsNumber uint `gorm:"not null"`
	Seats       []Seat
	IsFull      bool `gorm:"not null"`
	CreatedAt   time.Time
	UpdateAt    time.Time
}

func NewHall(theaterID,seatID uint, seatsNumber uint, isFull bool) *Hall {
	return &Hall{
		TheaterID:   theaterID,
		SeatsNumber: seatsNumber,
		IsFull:      isFull,
	}
}