package database

import (
	"movie/internal/model"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	postgresDB *gorm.DB
	mu         = &sync.Mutex{}
)

func NewPostgresDB(dsn string) (*gorm.DB, error) {
	mu.Lock()
	defer mu.Unlock()

	if postgresDB == nil {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		postgresDB = db

		err = migrate()
		if err != nil {
			return nil, err
		}
	}

	return postgresDB, nil
}

func migrate() error {
	err := postgresDB.AutoMigrate(
		&model.User{}, &model.Admin{},
		&model.Theater{}, &model.Hall{},
		&model.Screening{}, &model.Movie{},
		&model.Seat{}, &model.Reserve{},
	)

	return err
}
