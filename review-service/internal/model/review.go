package model

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID  uint
	MovieID uint
	Rating  float32
	Content string
	Status  string
}

func (Review) TableName() string {
	return "reviews"
}
