package role

import (
	"github.com/MateSousa/aegis/internal/domain/entity"
	"github.com/MateSousa/aegis/internal/repository/role"
	"github.com/google/uuid"
)

type IRoleUsecase interface {
	CreateRole(role *entity.Role) (*entity.Role, error)
	GetRoleByID(id uuid.UUID) (*entity.Role, error)
	GetRoleByName(name string) (*entity.Role, error)
}

type RoleUsecase struct {
	roleRepository role.IRoleRepository
}

func NewRoleUsecase(roleRepository role.IRoleRepository) IRoleUsecase {
	return &RoleUsecase{
		roleRepository: roleRepository,
	}
}

func (u *RoleUsecase) CreateRole(role *entity.Role) (*entity.Role, error) {
	return u.roleRepository.CreateRole(role)
}

func (u *RoleUsecase) GetRoleByID(id uuid.UUID) (*entity.Role, error) {
	return u.roleRepository.GetRoleByID(id)
}

func (u *RoleUsecase) GetRoleByName(name string) (*entity.Role, error) {
	return u.roleRepository.GetRoleByName(name)
}
