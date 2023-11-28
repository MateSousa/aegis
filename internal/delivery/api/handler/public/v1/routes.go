package v1

import (
	"github.com/MateSousa/aegis/internal/delivery/api/handler/private/v1/handler"
	customerHandler "github.com/MateSousa/aegis/internal/delivery/api/handler/public/v1/handler/customer"
	userHandler "github.com/MateSousa/aegis/internal/delivery/api/handler/public/v1/handler/user"
	customerRepository "github.com/MateSousa/aegis/internal/repository/customer"
	roleRepository "github.com/MateSousa/aegis/internal/repository/role"
	roleMappingRepository "github.com/MateSousa/aegis/internal/repository/rolemapping"
	tenantRepository "github.com/MateSousa/aegis/internal/repository/tenant"
	userRepository "github.com/MateSousa/aegis/internal/repository/user"
	customerUseCase "github.com/MateSousa/aegis/internal/usecase/customer"
	roleUseCase "github.com/MateSousa/aegis/internal/usecase/role"
	roleMappingUseCase "github.com/MateSousa/aegis/internal/usecase/rolemapping"
	tenantUseCase "github.com/MateSousa/aegis/internal/usecase/tenant"
	userUseCase "github.com/MateSousa/aegis/internal/usecase/user"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(r *echo.Group, connection *gorm.DB) {
	userRepository := userRepository.NewUserRepository(connection)
	userUseCase := userUseCase.NewUserUsecase(userRepository)
	roleRepository := roleRepository.NewRoleRepository(connection)
	roleUseCase := roleUseCase.NewRoleUsecase(roleRepository)
	roleMappingRepository := roleMappingRepository.NewRoleMappingRepository(connection)
	roleMappingUseCase := roleMappingUseCase.NewRoleMappingUsecase(roleMappingRepository)
	tenantRepository := tenantRepository.NewTenantRepository(connection)
	tenantUseCase := tenantUseCase.NewTenantUsecase(tenantRepository)
	customerRepository := customerRepository.NewCustomerRepository(connection)
	customerUseCase := customerUseCase.NewCustomerUsecase(customerRepository)

	userHandler := userHandler.NewUserHandler(userUseCase, roleMappingUseCase, roleUseCase, tenantUseCase)
	customerHandler := customerHandler.NewCustomerHandler(customerUseCase, userUseCase, roleMappingUseCase, roleUseCase, tenantUseCase)

	userHandler.InitRoutes(r)
	customerHandler.InitRoutes(r)

	health := handler.NewHealthHandler(connection)

	health.InitRoutes(r)
}
