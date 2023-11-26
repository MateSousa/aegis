package api

import (
	"os"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/MateSousa/aegis/docs"
	privateV1 "github.com/MateSousa/aegis/internal/delivery/api/handler/private/v1"
	publicV1 "github.com/MateSousa/aegis/internal/delivery/api/handler/public/v1"
	"gorm.io/gorm"

	"github.com/MateSousa/aegis/internal/driver/http"
	"github.com/sirupsen/logrus"
)

func InitRoutes(connection *gorm.DB) {
	logrus.Info("Initializing Api...")
	httpServer := http.New()

	// Use httpServer.Router to define your routes
	privateV1Group := httpServer.Router.Group("/api/v1")
	privateV1.InitRoutes(privateV1Group, connection)

	publicV1Group := httpServer.Router.Group("/api/v1")
	publicV1.InitRoutes(publicV1Group, connection)

	// Initialize Swagger
	InitSwagger(httpServer.Router)

	go httpServer.Run()
}

func InitSwagger(e *echo.Echo) {
	// Swagger configuration
	docs.SwaggerInfo.Title = os.Getenv("SWAGGER_TITLE")
	docs.SwaggerInfo.Description = "This document provides the documentation for a REST API designed."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{os.Getenv("SWAGGER_SCHEMES")}
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	docs.SwaggerInfo.BasePath = os.Getenv("SWAGGER_BASEPATH")
	e.GET("/docs/*", echoSwagger.WrapHandler)
}
