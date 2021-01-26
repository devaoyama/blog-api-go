package handler

import (
	"blog-api/domain/entity"
	"blog-api/usecase"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ArticleHandler struct {
	au usecase.ArticleUseCase
}

func NewArticleHandler (articleUseCase usecase.ArticleUseCase) ArticleHandler {
	return ArticleHandler{
		au: articleUseCase,
	}
}

func (h *ArticleHandler) GetAllArticle(c echo.Context) error {
	articles, err := h.au.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, articles)
}

func (h *ArticleHandler) GetArticleById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	article, err := h.au.GetById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, article)
}

func (h *ArticleHandler) PostArticle(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(float64)

	article := &entity.Article{}
	err := c.Bind(article)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	article, err = h.au.PostArticle(int(userId), article)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, article)
}

func (h *ArticleHandler) UpdateArticle(c echo.Context) error {
	articleId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	article := &entity.Article{}
	err = c.Bind(article)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(float64)

	article, err = h.au.EditArticle(int(userId), articleId, article)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, article)
}
