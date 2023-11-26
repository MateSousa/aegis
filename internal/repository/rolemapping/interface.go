package rolemapping

import (
	"github.com/MateSousa/aegis/internal/domain/entity"
	"github.com/google/uuid"
)

type IRoleMappingRepository interface {
	CreateRoleMapping(roleMapping *entity.RoleMapping) (*entity.RoleMapping, error)
	GetRoleMappingByID(id uuid.UUID) (*entity.RoleMapping, error)
	GetRoleMappingByUserID(userID uuid.UUID) (*entity.RoleMapping, error)
}
