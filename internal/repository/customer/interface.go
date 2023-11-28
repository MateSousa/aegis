package customer

import (
	"github.com/MateSousa/aegis/internal/domain/entity"
	"github.com/google/uuid"
)

type ICustomerRepository interface {
	Create(customer *entity.Customer) (*entity.Customer, error)
	Update(customer *entity.Customer) (*entity.Customer, error)
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.Customer, error)
	FindByUserId(userId uuid.UUID) (*entity.Customer, error)
	FindByTenantId(tenantId uuid.UUID) ([]*entity.Customer, error)
}
