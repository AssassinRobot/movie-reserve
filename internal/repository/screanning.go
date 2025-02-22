package repository

import "movie/internal/model"

type ScreeningRepository interface {
	GetScreeningByID(ID uint) (*model.Screening, error)
	GetScreenings() ([]model.Screening, error)
	CreateScreening(screening *model.Screening) error
	UpdateScreeningByID(ID uint, screening *model.Screening) error
	DeleteScreeningByID(ID uint) error
}