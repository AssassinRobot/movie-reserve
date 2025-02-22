package repository

import "movie/internal/model"

type TheaterRepository interface {
	GetTheaterByID(ID uint) (*model.Theater, error)
	GetTheaters() ([]model.Theater, error)
	GetTheatersByName(name string) ([]*model.Theater, error)
	CreateTheater(theater *model.Theater) error
	UpdateTheaterByID(ID uint, theater *model.Theater) error
	DeleteTheaterByID(ID uint) error
}
