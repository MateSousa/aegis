package migrations

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
  "github.com/MateSousa/aegis/internal/domain/entity"
)

func Setup(wr *gorm.DB, rd *gorm.DB) {

	err := wr.AutoMigrate(
    &entity.User{},
    &entity.Role{},
    &entity.RoleMapping{}
  )

	if err != nil {
		logrus.Error("Error on Auto migrate WRITE", err)
	}

	err = rd.AutoMigrate(
    &entity.User{},
    &entity.Role{},
    &entity.RoleMapping{}
  )

	if err != nil {
		logrus.Error("Error on Auto migrate Read", err)
	}
}
