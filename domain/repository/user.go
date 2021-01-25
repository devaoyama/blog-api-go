package repository

import "blog-api/domain"

type UserRepository interface {
	FindAll() ([]domain.User, error)
	Create(user *domain.User) (*domain.User, error)
	Find(id int) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	Delete(id int) error
}
