package v1

import (
	"github.com/MateSousa/aegis/internal/delivery/api/handler/private/v1/handler"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Group, connections map[string]*gorm.DB) {
	health := handler.InitHealthHandler(connections)

	health.InitRoutes(e)
}
