package usecase

import (
	"blog-api/domain/entity"
	"blog-api/domain/repository"
	"errors"
)

type ArticleUseCase struct {
	articleRepository repository.ArticleRepository
}

func NewArticleUseCase(ar repository.ArticleRepository) ArticleUseCase {
	return ArticleUseCase{
		articleRepository: ar,
	}
}

func (au *ArticleUseCase) GetAll() ([]entity.Article, error) {
	return au.articleRepository.FindAll()
}

func (au *ArticleUseCase) GetById(id int) (*entity.Article, error) {
	return au.articleRepository.Find(id)
}

func (au *ArticleUseCase) PostArticle(userId int, article *entity.Article) (*entity.Article, error) {
	article.UserId = userId
	return au.articleRepository.Create(article)
}

func (au *ArticleUseCase) EditArticle(userId, articleId int, newArticle *entity.Article) (*entity.Article, error) {
	article, err := au.articleRepository.Find(articleId)
	if err != nil {
		return nil, err
	}

	if userId != article.UserId {
		return nil, errors.New("unauthorized")
	}

	article.Title = newArticle.Title
	article.Content = newArticle.Content

	return au.articleRepository.Update(article)
}
