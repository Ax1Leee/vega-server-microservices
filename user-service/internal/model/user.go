package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string
	HashedPassword string
	Avatar         string
	Name           string
	Gender         string
	BirthDate      string
	Location       string
	Bio            string
}

func (User) TableName() string {
	return "users"
}
