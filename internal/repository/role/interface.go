package role

import (
	"github.com/MateSousa/aegis/internal/domain/entity"
	"github.com/google/uuid"
)

type IRoleRepository interface {
	CreateRole(role *entity.Role) (*entity.Role, error)
	GetRoleByID(id uuid.UUID) (*entity.Role, error)
	GetRoleByName(name string) (*entity.Role, error)
}
