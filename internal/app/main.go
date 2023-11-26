package app

import (
	"os"

	"github.com/MateSousa/aegis/internal/delivery/api"
	"github.com/MateSousa/aegis/internal/domain/migrations"
	"github.com/MateSousa/aegis/internal/driver/database"
	"github.com/MateSousa/aegis/internal/driver/logs"
	"gorm.io/gorm"
)

func StartApi() {
	logs.InitLogrus()

	dbRead := os.Getenv("DB_CONNECTION_READ")
	dbWrite := os.Getenv("DB_CONNECTION_WRITE")

	connections := map[string]*gorm.DB{
		"wr": database.New(dbWrite),
		"rd": database.New(dbRead),
	}

	migrations.Setup(connections["wr"], connections["rd"])

	api.InitRoutes(connections)
}
