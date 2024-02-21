package handler

import (
	"gorm.io/gorm"
)

type ApiConfig struct {
	DBPointer *gorm.DB
}
