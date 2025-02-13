package model

import "time"

type User struct {
	ID        uint   `gorm:"primary_key"`
	FName     string `gorm:"not null"`
	LName     string `gorm:"not null"`
	Phone     string
	Email     string `gorm:"not null"`
	Reserves  []Reserve `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Admin struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(fName, lName, phone, email string) *User {
	return &User{
		FName: fName,
		LName: lName,
		Phone: phone,
		Email: email,
	}
}

func NewAdmin(username, password string) *Admin {
	return &Admin{
		Username: username,
		Password: password,
	}
}
