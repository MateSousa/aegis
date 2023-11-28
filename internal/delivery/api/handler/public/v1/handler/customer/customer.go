package customer

import (
	"net/http"

	"github.com/MateSousa/aegis/internal/domain/entity"
	customerUseCase "github.com/MateSousa/aegis/internal/usecase/customer"
	roleUseCase "github.com/MateSousa/aegis/internal/usecase/role"
	roleMappingUseCase "github.com/MateSousa/aegis/internal/usecase/rolemapping"
	tenantUseCase "github.com/MateSousa/aegis/internal/usecase/tenant"
	userUseCase "github.com/MateSousa/aegis/internal/usecase/user"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

type CustomerHandler struct {
	CustomerUseCase    customerUseCase.ICustomerUsecase
	UserUseCase        userUseCase.IUserUsecase
	RoleMappingUseCase roleMappingUseCase.IRoleMappingUsecase
	RoleUseCase        roleUseCase.IRoleUsecase
	TenantUseCase      tenantUseCase.ITenantUsecase
}

func NewCustomerHandler(customerUseCase customerUseCase.ICustomerUsecase, userUseCase userUseCase.IUserUsecase, roleMappingUseCase roleMappingUseCase.IRoleMappingUsecase, roleUseCase roleUseCase.IRoleUsecase, tenantUseCase tenantUseCase.ITenantUsecase) *CustomerHandler {
	return &CustomerHandler{
		customerUseCase,
		userUseCase,
		roleMappingUseCase,
		roleUseCase,
		tenantUseCase,
	}
}

func (h *CustomerHandler) InitRoutes(e *echo.Group) {
	e.POST("/customers/signup", h.CreateCustomer())
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func (h *CustomerHandler) CreateCustomer() echo.HandlerFunc {
	return func(c echo.Context) error {
		tenantClientId, err := h.getTenantClientID(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		tenant, err := h.TenantUseCase.FindByClientId(tenantClientId)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		user := new(entity.User)
		if err := c.Bind(user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		user, err = h.UserUseCase.CreateUser(user)
		if err != nil {
			if err == entity.ErrEmailExists {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		role, err := h.RoleUseCase.GetRoleByName("customer")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		roleMapping := &entity.RoleMapping{
			UserID:   user.ID,
			RoleID:   role.ID,
			TenantID: tenant.ID,
		}
		_, err = h.RoleMappingUseCase.CreateRoleMapping(roleMapping)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		customer := new(entity.Customer)
		customer.UserID = user.ID
		customer.TenantID = tenant.ID
		customer, err = h.CustomerUseCase.Create(customer)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, customer)
	}
}

func (h *CustomerHandler) getTenantClientID(c echo.Context) (uuid.UUID, error) {
	clientID := c.Request().Header.Get("X-Tenant-Client-Id")
	if clientID == "" {
		return uuid.Nil, echo.NewHTTPError(http.StatusBadRequest, "X-Tenant-Client-Id header is required")
	}
	id, err := uuid.Parse(clientID)
	if err != nil {
		return uuid.Nil, echo.NewHTTPError(http.StatusBadRequest, "X-Tenant-Client-Id header is invalid")
	}
	return id, nil
}
