package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string `gorm:"not null"`
	Apikey    string `gorm:"not null"`
}

type UserFeed struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string `gorm:"not null"`
	Url       string `gorm:"not null"`
	UserRefer uuid.UUID
	User      User `gorm:"foreignKey:UserRefer"`
}

type FeedFollows struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	UserRefer uuid.UUID
	FeedRefer uuid.UUID
	User      User     `gorm:"foreignKey:UserRefer,uniqueIndex:Feed"`
	Feed      UserFeed `gorm:"foreignKey:FeedRefer,uniqueIndex:User"`
}
