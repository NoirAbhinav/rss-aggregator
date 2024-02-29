package models

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

func (u *FeedFollows) Create(db *gorm.DB) error {
	return db.Create(u).Error
}

func (u *FeedFollows) Update(db *gorm.DB) error {
	return db.Save(u).Error
}

func (u *FeedFollows) Delete(db *gorm.DB) error {
	return db.Delete(u).Error
}

func (u *FeedFollows) Select(db *gorm.DB) (feedfollow *FeedFollows, err error) {
	err = db.Model(&FeedFollows{}).Take(&feedfollow, u).Error
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to select user")
	}
	return feedfollow, nil
}
