package persistence

import (
	"blog-api/domain"
	"blog-api/domain/repository"
	"gorm.io/gorm"
)

type articlePersistence struct {
	Db *gorm.DB
}

func NewArticlePersistence(db *gorm.DB) repository.ArticleRepository {
	return &articlePersistence{Db: db}
}

func (ap *articlePersistence) FindAll() ([]domain.Article, error) {
	var articles []domain.Article
	result := ap.Db.Find(&articles)
	if err := result.Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (ap *articlePersistence) Create(article *domain.Article) (*domain.Article, error) {
	result := ap.Db.Create(article)
	if err := result.Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (ap *articlePersistence) Find(id int) (*domain.Article, error) {
	article := &domain.Article{}
	result := ap.Db.First(article, id)
	if err := result.Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (ap *articlePersistence) Update(article *domain.Article) (*domain.Article, error) {
	result := ap.Db.Save(article)
	if err := result.Error; err != nil {
		return nil, result.Error
	}
	return article, nil
}

func (ap *articlePersistence) Delete(id int) error {
	result := ap.Db.Delete(&domain.Article{}, id)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
