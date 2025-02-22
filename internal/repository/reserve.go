package repository

import "movie/internal/model"

type ReserveRepository interface {
	GetReserveByID(ID uint) (*model.Reserve, error)
	GetReserves() ([]model.Reserve, error)
	CreateReserve(reserve *model.Reserve) error
	DeleteReserveByID(ID uint) error
}
