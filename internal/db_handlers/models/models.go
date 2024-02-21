package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"not null"`
}

type Method interface {
	Create() error
	Update() error
	Delete() error
}

func (u *User) Create(db *gorm.DB) error {
	return db.Create(u).Error
}
