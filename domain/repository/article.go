package repository

import (
	"blog-api/domain/entity"
)

type ArticleRepository interface {
	FindAll() ([]entity.Article, error)
	Create(article *entity.Article) (*entity.Article, error)
	Find(id int) (*entity.Article, error)
	Update(article *entity.Article) (*entity.Article, error)
	Delete(id int) error
}
