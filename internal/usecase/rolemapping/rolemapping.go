package rolemapping

import (
	"github.com/MateSousa/aegis/internal/domain/entity"
	"github.com/MateSousa/aegis/internal/repository/rolemapping"
	"github.com/google/uuid"
)

type IRoleMappingUsecase interface {
	CreateRoleMapping(roleMapping *entity.RoleMapping) (*entity.RoleMapping, error)
	GetRoleMappingByID(id uuid.UUID) (*entity.RoleMapping, error)
	GetRoleMappingByUserID(userID uuid.UUID) (*entity.RoleMapping, error)
}

type RoleMappingUsecase struct {
	roleMappingRepository rolemapping.IRoleMappingRepository
}

func NewRoleMappingUsecase(roleMappingRepository rolemapping.IRoleMappingRepository) IRoleMappingUsecase {
	return &RoleMappingUsecase{
		roleMappingRepository: roleMappingRepository,
	}
}

func (u *RoleMappingUsecase) CreateRoleMapping(roleMapping *entity.RoleMapping) (*entity.RoleMapping, error) {
	return u.roleMappingRepository.CreateRoleMapping(roleMapping)
}

func (u *RoleMappingUsecase) GetRoleMappingByID(id uuid.UUID) (*entity.RoleMapping, error) {
	return u.roleMappingRepository.GetRoleMappingByID(id)
}

func (u *RoleMappingUsecase) GetRoleMappingByUserID(userID uuid.UUID) (*entity.RoleMapping, error) {
	return u.roleMappingRepository.GetRoleMappingByUserID(userID)
}
