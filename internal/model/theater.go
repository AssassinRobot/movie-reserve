package model

import "time"

type Theater struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"not null"`
	Address   string `gorm:"not null"`
	Halls     []Hall `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTheater(name, address string) *Theater {
	return &Theater{
		Name:    name,
		Address: address,
	}
}