package handler

import (
	"fmt"
	"net/http"

	"github.com/MateSousa/aegis/internal/domain/entity"
	roleUseCase "github.com/MateSousa/aegis/internal/usecase/role"
	roleMappingUseCase "github.com/MateSousa/aegis/internal/usecase/rolemapping"
	userUseCase "github.com/MateSousa/aegis/internal/usecase/user"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

type UserHandler struct {
	UserUseCase        userUseCase.IUserUsecase
	RoleMappingUseCase roleMappingUseCase.IRoleMappingUsecase
	RoleUseCase        roleUseCase.IRoleUsecase
}

func NewUserHandler(userUseCase userUseCase.IUserUsecase, roleMappingUseCase roleMappingUseCase.IRoleMappingUsecase, roleUseCase roleUseCase.IRoleUsecase) *UserHandler {
	return &UserHandler{
		userUseCase,
		roleMappingUseCase,
		roleUseCase,
	}
}

func (h *UserHandler) InitRoutes(e *echo.Group) {
	e.GET("/users", h.GetUsers())
	e.GET("/users/:id", h.GetUserByID())
	e.POST("/users", h.CreateUser())
	e.PUT("/users", h.UpdateUser())
	e.DELETE("/users/:id", h.DeleteUser())
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
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

		if err := c.Validate(user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		user, err := h.UserUseCase.CreateUser(user)
		if err != nil {
			if err == entity.ErrEmailExists {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return err
		}

		role := new(entity.Role)
		role, err = h.RoleUseCase.GetRoleByName("admin")
		if err != nil {
			fmt.Println("Error getting role by name")
			return err
		}

		roleMapping := new(entity.RoleMapping)
		roleMapping = &entity.RoleMapping{
			UserID: user.ID,
			RoleID: role.ID,
		}

		_, err = h.RoleMappingUseCase.CreateRoleMapping(roleMapping)
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

		if err := c.Validate(user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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
