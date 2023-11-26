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
	HealthDBConnections *gorm.DB
}

func (c HealthHandler) InitRoutes(e *echo.Group) {
	e.GET("/healthz", c.Live)
}

func (c HealthHandler) Live(ctx echo.Context) error {
	healthRepositorys := health.NewHealthRepository(c.HealthDBConnections)
	if err := healthRepositorys.GetHealthDB(ctx.Request().Context()); err != nil {
		logrus.Error("Error on health check for connection: ", err)
		ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("Error on health check for connection: %s", err))
	}
	return nil

	return ctx.JSON(http.StatusOK, "API is live")
}

func NewHealthHandler(connection *gorm.DB) *HealthHandler {
	return &HealthHandler{
		HealthDBConnections: connection,
	}
}

func InitHealthHandler(connection *gorm.DB) *HealthHandler {
	return NewHealthHandler(connection)
}
