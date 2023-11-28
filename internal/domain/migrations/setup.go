package migrations

import (
	"github.com/MateSousa/aegis/internal/domain/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Setup(wr *gorm.DB) {
	if err := wr.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		logrus.Error("Error on CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";", err)
	}

	err := wr.AutoMigrate(
		&entity.User{},
		&entity.Role{},
		&entity.RoleMapping{},
		&entity.Tenant{},
		&entity.Customer{},
	)
	if err != nil {
		logrus.Error("Error on Auto migrate WRITE", err)
	}

	roleNames := []string{"admin", "member", "user", "customer"}
	for _, roleName := range roleNames {
		var role entity.Role
		if err := wr.Where("name = ?", roleName).First(&role).Error; err != nil {
			wr.Create(&entity.Role{Name: roleName})
		}
	}
}
