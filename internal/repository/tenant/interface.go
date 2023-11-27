package tenant

import (
	"github.com/MateSousa/aegis/internal/domain/entity"
	"github.com/google/uuid"
)

type ITenantRepository interface {
	Create(tenant *entity.Tenant) error
	Update(tenant *entity.Tenant) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.Tenant, error)
	FindAll() ([]*entity.Tenant, error)
	FindByUserId(userId uuid.UUID) ([]*entity.Tenant, error)
}
