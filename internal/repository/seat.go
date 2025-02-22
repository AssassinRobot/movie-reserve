package repository

import "movie/internal/model"

type SeatRepository interface {
	GetSeatByID(ID uint) (*model.Seat, error)
	GetSeats() ([]model.Seat, error)
	CreateSeat(seat *model.Seat) error
	UpdateSeatByID(ID uint, seat *model.Seat) error
	DeleteSeatByID(ID uint) error
}