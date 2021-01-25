package entity

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string
	Content string
	UserId  int
	User    User
}
