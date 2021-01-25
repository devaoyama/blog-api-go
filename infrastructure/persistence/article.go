package persistence

import (
	"blog-api/domain/entity"
	"blog-api/domain/repository"
	"gorm.io/gorm"
)

type articlePersistence struct {
	Db *gorm.DB
}

func NewArticlePersistence(db *gorm.DB) repository.ArticleRepository {
	return &articlePersistence{Db: db}
}

func (ap *articlePersistence) FindAll() ([]entity.Article, error) {
	var articles []entity.Article
	result := ap.Db.Find(&articles)
	if err := result.Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (ap *articlePersistence) Create(article *entity.Article) (*entity.Article, error) {
	result := ap.Db.Create(article)
	if err := result.Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (ap *articlePersistence) Find(id int) (*entity.Article, error) {
	article := &entity.Article{}
	result := ap.Db.First(article, id)
	if err := result.Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (ap *articlePersistence) Update(article *entity.Article) (*entity.Article, error) {
	result := ap.Db.Save(article)
	if err := result.Error; err != nil {
		return nil, result.Error
	}
	return article, nil
}

func (ap *articlePersistence) Delete(id int) error {
	result := ap.Db.Delete(&entity.Article{}, id)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
