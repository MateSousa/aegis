package app

import (
	"os"

	"github.com/MateSousa/aegis/internal/delivery/api"
	"github.com/MateSousa/aegis/internal/domain/migrations"
	"github.com/MateSousa/aegis/internal/driver/database"
	"github.com/MateSousa/aegis/internal/driver/logs"
)

func StartApi() {
	logs.InitLogrus()

	dbRead := os.Getenv("DB_CONNECTION_READ")

	connection := database.New(dbRead)

	migrations.Setup(connection)

	api.InitRoutes(connection)
}
