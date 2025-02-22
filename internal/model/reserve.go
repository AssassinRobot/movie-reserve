package model

import (
	"time"

	"gorm.io/gorm"
)

type Reserve struct {
	ID        uint `gorm:"primary_key"`
	SeatID    uint `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID    uint `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	Deleted   gorm.DeletedAt
}

func NewReserve(seatID, userID uint) *Reserve {
	return &Reserve{
		SeatID: seatID,
		UserID: userID,
	}
}
