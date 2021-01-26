package handler

import (
	"blog-api/domain/entity"
	"blog-api/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase) UserHandler {
	return UserHandler{userUseCase: userUseCase}
}

func (h *UserHandler) SignUp(c echo.Context) error {
	var user entity.User
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := h.userUseCase.SignUp(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, token)
}

func (h *UserHandler) Login(c echo.Context) error {
	token, err := h.userUseCase.Login(c.FormValue("email"), c.FormValue("password"))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
	return c.String(http.StatusOK, token)
}

func (h *UserHandler) GetAllUser(c echo.Context) error {
	users, err := h.userUseCase.GetAllUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	user, err := h.userUseCase.GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
