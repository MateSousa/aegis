package handler

import (
	"fmt"
	"net/http"

	"github.com/MateSousa/aegis/internal/repository/health"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HealthHandler struct {
	HealthDBConnections map[string]*gorm.DB
}

func (c HealthHandler) InitRoutes(e *echo.Group) {
	e.GET("/healthz", c.Live)
}

func (c HealthHandler) Live(ctx echo.Context) {
	for connectionName, connection := range c.HealthDBConnections {
		healthRepositorys := health.NewRepository(connection)
		if err := healthRepositorys.Check(); err != nil {
			logrus.Error("Error on health check for connection: ", connectionName, err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("Error on health check for connection: %s", connectionName))
		}
		return
	}

	return ctx.JSON(http.StatusOK, "API is live")
}

func NewHealthHandler(connections map[string]*gorm.DB) *HealthHandler {
	return &HealthHandler{
		HealthDBConnections: connections,
	}
}

func InitHealthHandler(connections map[string]*gorm.DB) *HealthHandler {
	return NewHealthHandler(connections)
}
