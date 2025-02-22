package repository

import "movie/internal/model"

type MovieRepository interface {
	GetMovieByID(ID uint) (*model.Movie, error)
	GetMovies() ([]model.Movie, error)
	CreateMovie(movie *model.Movie) error
	UpdateMovieByID(ID uint, movie *model.Movie) error
	DeleteMovieByID(ID uint) error
	GetMoviesByTitle(title string) ([]*model.Movie, error)
}

type GenreRepository interface {
	GetGenreByID(ID uint) (*model.Genre, error)
	GetGenres() ([]model.Genre, error)
	CreateGenre(genre *model.Genre) error
	UpdateGenreByID(ID uint, genre *model.Genre) error
	DeleteGenreByID(ID uint) error
	GetGenresByName(name string) ([]*model.Genre, error)
}
