package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email         string
	Name          string
	Password      string
	PlainPassword string `gorm:"-"`
	Articles      []Article
}
