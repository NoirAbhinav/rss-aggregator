package models

import (
	"gorm.io/gorm"
)

type Method interface {
	Create() error
	Update() error
	Delete() error
}

func (u *User) Create(db *gorm.DB) error {
	return db.Create(u).Error
}
