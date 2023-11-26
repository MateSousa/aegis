package rolemapping

import (
	"github.com/MateSousa/aegis/internal/domain/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleMappingRepository struct {
	db *gorm.DB
}

func NewRoleMappingRepository(connection *gorm.DB) *RoleMappingRepository {
	return &RoleMappingRepository{
		db: connection,
	}
}

func (r *RoleMappingRepository) CreateRoleMapping(roleMapping *entity.RoleMapping) (*entity.RoleMapping, error) {
	if err := r.db.Create(roleMapping).Error; err != nil {
		return nil, err
	}
	return roleMapping, nil
}

func (r *RoleMappingRepository) GetRoleMappingByID(id uuid.UUID) (*entity.RoleMapping, error) {
	roleMapping := &entity.RoleMapping{}
	if err := r.db.Where("id = ?", id).First(roleMapping).Error; err != nil {
		return nil, err
	}
	return roleMapping, nil
}

func (r *RoleMappingRepository) GetRoleMappingByUserID(userID uuid.UUID) (*entity.RoleMapping, error) {
	roleMapping := &entity.RoleMapping{}
	if err := r.db.Where("user_id = ?", userID).First(roleMapping).Error; err != nil {
		return nil, err
	}
	return roleMapping, nil
}
