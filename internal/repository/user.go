package repository

import "movie/internal/model"

type UserRepository interface {
	GetUserByID(ID uint) (*model.User, error)
	GetUsers() ([]model.User, error)
	CreateUser(user *model.User) error
	UpdateUserByID(ID uint, user *model.User) error
	DeleteUserByID(ID uint) error
}

type AdminRepository interface {
	GetAdminByID(ID uint) (*model.Admin, error)
	GetAdmins() ([]model.Admin, error)
	CreateAdmin(admin *model.Admin) error
	UpdateAdminByID(ID uint, admin *model.Admin) error
	DeleteAdminByID(ID uint) error
}

