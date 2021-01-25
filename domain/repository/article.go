package repository

import (
	"blog-api/domain"
)

type ArticleRepository interface {
	FindAll() ([]domain.Article, error)
	Create(article *domain.Article) (*domain.Article, error)
	Find(id int) (*domain.Article, error)
	Update(article *domain.Article) (*domain.Article, error)
	Delete(id int) error
}
