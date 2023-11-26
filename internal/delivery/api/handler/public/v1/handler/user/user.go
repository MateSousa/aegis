package handler

import (
	"net/http"

	"github.com/MateSousa/aegis/internal/domain/entity"
	userUseCase "github.com/MateSousa/aegis/internal/usecase/user"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUseCase userUseCase.IUserUsecase
}

func NewUserHandler(userUseCase userUseCase.IUserUsecase) *UserHandler {
	return &UserHandler{
		userUseCase,
	}
}

func (h *UserHandler) InitRoutes(e *echo.Group) {
	e.GET("/users", h.GetUsers())
	e.GET("/users/:id", h.GetUserByID())
	e.POST("/users", h.CreateUser())
	e.PUT("/users", h.UpdateUser())
	e.DELETE("/users/:id", h.DeleteUser())
}

func (h *UserHandler) GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := h.UserUseCase.GetUsers()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, users)
	}
}

func (h *UserHandler) GetUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return err
		}
		user, err := h.UserUseCase.GetUserByID(id)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user)
	}
}

func (h *UserHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(entity.User)
		if err := c.Bind(user); err != nil {
			return err
		}
		user, err := h.UserUseCase.CreateUser(user)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user)
	}
}

func (h *UserHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(entity.User)
		if err := c.Bind(user); err != nil {
			return err
		}
		user, err := h.UserUseCase.UpdateUser(user)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user)
	}
}

func (h *UserHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return err
		}
		err = h.UserUseCase.DeleteUser(id)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, nil)
	}
}
