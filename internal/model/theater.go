package model

import (
	"time"

	"gorm.io/gorm"
)

type Theater struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"not null"`
	Address   string `gorm:"not null"`
	Halls     []Hall `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Screening []Screening `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   gorm.DeletedAt
}

func NewTheater(name, address string) *Theater {
	return &Theater{
		Name:    name,
		Address: address,
	}
}
