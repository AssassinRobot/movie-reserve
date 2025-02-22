package model

import (
	"time"

	"gorm.io/gorm"
)

type Screening struct {
	ID        uint   `gorm:"primary_key"`
	MovieID   uint   `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	HallID    uint   `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TheaterID uint   `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StartTime string `gorm:"not null"`
	EndTime   string `gorm:"not null"`
	Status    string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   gorm.DeletedAt
}

func NewScreening(movieID, hallID, theaterID uint, startTime, endTime, status string) *Screening {
	return &Screening{
		MovieID:   movieID,
		HallID:    hallID,
		TheaterID: theaterID,
		StartTime: startTime,
		EndTime:   endTime,
		Status:    status,
	}
}
