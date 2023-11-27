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
		&entity.Tenant{},
	)
	if err != nil {
		logrus.Error("Error on Auto migrate WRITE", err)
	}

	if err := wr.FirstOrCreate(&entity.Role{Name: "admin"}, &entity.Role{Name: "member"}).Error; err != nil {
		logrus.Error("Error on create admin role", err)
	}
}
