package v1

import (
	"github.com/MateSousa/aegis/internal/delivery/api/handler/private/v1/handler"
	userHandler "github.com/MateSousa/aegis/internal/delivery/api/handler/public/v1/user"
	userRepository "github.com/MateSousa/aegis/internal/repository/user"
	userUseCase "github.com/MateSousa/aegis/internal/usecase/user"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(r *echo.Group, connections map[string]*gorm.DB) {
	userRepository := userRepository.NewUserRepository(connections)
	userUseCase := userUseCase.NewUserUseCase(userRepository)
	userHandler := userHandler.NewUserHandler(userUseCase)

	userHandler.InitRoutes(r)

	health := handler.NewHealthHandler(connections)

	health.InitRoutes(r)
}
