package v1

import (
	"github.com/MateSousa/aegis/internal/delivery/api/handler/private/v1/handler"
	userHandler "github.com/MateSousa/aegis/internal/delivery/api/handler/public/v1/handler/user"
	roleRepository "github.com/MateSousa/aegis/internal/repository/role"
	roleMappingRepository "github.com/MateSousa/aegis/internal/repository/rolemapping"
	userRepository "github.com/MateSousa/aegis/internal/repository/user"
	roleUseCase "github.com/MateSousa/aegis/internal/usecase/role"
	roleMappingUseCase "github.com/MateSousa/aegis/internal/usecase/rolemapping"
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

	userHandler := userHandler.NewUserHandler(userUseCase, roleMappingUseCase, roleUseCase)

	userHandler.InitRoutes(r)

	health := handler.NewHealthHandler(connection)

	health.InitRoutes(r)
}
