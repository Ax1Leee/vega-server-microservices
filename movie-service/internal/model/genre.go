package model

import (
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	Name string
}

func (Genre) TableName() string {
	return "genres"
}
