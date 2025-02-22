package model

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	ID          uint    `gorm:"primary_key"`
	Title       string  `gorm:"not null"`
	Year        uint    `gorm:"not null"`
	Duration    uint    `gorm:"not null"`
	Description string  `gorm:"not null"`
	Genres      []Genre `gorm:"many2many:movie_genres;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Deleted     gorm.DeletedAt
}

type Genre struct {
	ID        uint    `gorm:"primary_key"`
	Name      string  `gorm:"not null"`
	Movies    []Movie `gorm:"many2many:movie_genres;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   gorm.DeletedAt
}

func NewMovie(title string, year, duration uint, description string) *Movie {
	return &Movie{
		Title:       title,
		Year:        year,
		Duration:    duration,
		Description: description,
	}
}

func NewGenre(name string) *Genre {
	return &Genre{
		Name: name,
	}
}
