package repository

import "movie/internal/model"

type HallRepository interface {
	GetHallByID(ID uint) (*model.Hall, error)
	GetHalls() ([]model.Hall, error)
	CreateHall(hall *model.Hall) error
	UpdateHallByID(ID uint, hall *model.Hall) error
	DeleteHallByID(ID uint) error
}