package model

import (
	"gorm.io/gorm"
)

type Star struct {
	gorm.Model
	Name string
}

func (Star) TableName() string {
	return "stars"
}
