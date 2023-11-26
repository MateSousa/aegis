package migrations

import (
	"github.com/MateSousa/aegis/internal/domain/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Setup(wr *gorm.DB) {
	err := wr.AutoMigrate(
		&entity.User{},
		&entity.Role{},
		&entity.RoleMapping{},
	)
	if err != nil {
		logrus.Error("Error on Auto migrate WRITE", err)
	}
}
