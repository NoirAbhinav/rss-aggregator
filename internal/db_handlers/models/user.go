package models

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

func (u *User) Create(db *gorm.DB) error {
	return db.Create(u).Error
}

func (u *User) Update(db *gorm.DB) error {
	return db.Save(u).Error
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Delete(u).Error
}

func (u *User) Select(db *gorm.DB) (user *User, err error) {
	err = db.Model(&User{}).Take(&user, u).Error
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to select user")
	}
	return user, nil
}
