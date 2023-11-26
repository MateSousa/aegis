package api

import (
	"os"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"

	"github.com/MateSousa/aegis/docs"
	privateV1 "github.com/MateSousa/aegis/internal/delivery/api/handler/private/v1"
	publicV1 "github.com/MateSousa/aegis/internal/delivery/api/handler/public/v1"
	"gorm.io/gorm"

	"github.com/MateSousa/aegis/internal/driver/http"
	"github.com/sirupsen/logrus"

	"github.com/MateSousa/aegis/internal/driver/logs"
)

func InitRoutes(connections map[string]*gorm.DB) {
	logrus.Info("Initializing Api...")

	e := echo.New()

	InitSwagger(e)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Skip paths logic needs to be integrated into the logger middleware if required

	// Initialize Prometheus
	p := middleware.NewPrometheus("echo")
	p.Use(e)

	// Initialize Routes
	PrivateV1Group := e.Group("/api/private/v1")
	privateV1.InitRoutes(PrivateV1Group, connections)

	PublicV1 := e.Group("/api/public/v1")
	publicV1.InitRoutes(PublicV1)

	// Start server
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
}

func InitSwagger(e *echo.Echo) {
	docs.SwaggerInfo.Title = os.Getenv("SWAGGER_TITLE")
	docs.SwaggerInfo.Description = "This document provides the documentation for a REST API designed."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{os.Getenv("SWAGGER_SCHEMES")}
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	docs.SwaggerInfo.BasePath = os.Getenv("SWAGGER_BASEPATH")
	e.GET("/docs/*", echoSwagger.WrapHandler)
}

