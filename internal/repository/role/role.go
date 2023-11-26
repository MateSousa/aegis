package role

import (
	"github.com/MateSousa/aegis/internal/domain/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(connection *gorm.DB) *RoleRepository {
	return &RoleRepository{
		db: connection,
	}
}

func (r *RoleRepository) CreateRole(role *entity.Role) (*entity.Role, error) {
	if err := r.db.Create(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *RoleRepository) GetRoleByID(id uuid.UUID) (*entity.Role, error) {
	role := &entity.Role{}
	if err := r.db.Where("id = ?", id).First(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *RoleRepository) GetRoleByName(name string) (*entity.Role, error) {
	role := &entity.Role{}
	if err := r.db.Where("name = ?", name).First(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}
