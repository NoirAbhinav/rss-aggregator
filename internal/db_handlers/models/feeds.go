package models

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

func (u *UserFeed) Create(db *gorm.DB) error {
	return db.Create(u).Error
}

func (u *UserFeed) Update(db *gorm.DB) error {
	return db.Save(u).Error
}

func (u *UserFeed) Delete(db *gorm.DB) error {
	return db.Delete(u).Error
}

func (u *UserFeed) Select(db *gorm.DB) (userfeed *UserFeed, err error) {
	err = db.Model(&UserFeed{}).Take(&userfeed, u).Error
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to select user")
	}
	return userfeed, nil
}
