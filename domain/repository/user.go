package repository

import (
	"blog-api/domain/entity"
)

type UserRepository interface {
	FindAll() ([]entity.User, error)
	Create(user *entity.User) (*entity.User, error)
	Find(id int) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	Delete(id int) error
}
