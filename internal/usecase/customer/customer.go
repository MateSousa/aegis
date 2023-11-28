package customer

import (
	"github.com/MateSousa/aegis/internal/domain/entity"
	"github.com/MateSousa/aegis/internal/repository/customer"
	"github.com/google/uuid"
)

type ICustomerUsecase interface {
	Create(customer *entity.Customer) (*entity.Customer, error)
	Update(customer *entity.Customer) (*entity.Customer, error)
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.Customer, error)
	FindByUserId(userId uuid.UUID) (*entity.Customer, error)
	FindByTenantId(tenantId uuid.UUID) ([]*entity.Customer, error)
}

type CustomerUsecase struct {
	customerRepository customer.ICustomerRepository
}

func NewCustomerUsecase(customerRepository customer.ICustomerRepository) ICustomerUsecase {
	return &CustomerUsecase{
		customerRepository: customerRepository,
	}
}

func (u *CustomerUsecase) Create(customer *entity.Customer) (*entity.Customer, error) {
	return u.customerRepository.Create(customer)
}

func (u *CustomerUsecase) Update(customer *entity.Customer) (*entity.Customer, error) {
	return u.customerRepository.Update(customer)
}

func (u *CustomerUsecase) Delete(id uuid.UUID) error {
	return u.customerRepository.Delete(id)
}

func (u *CustomerUsecase) FindById(id uuid.UUID) (*entity.Customer, error) {
	return u.customerRepository.FindById(id)
}

func (u *CustomerUsecase) FindByUserId(userId uuid.UUID) (*entity.Customer, error) {
	return u.customerRepository.FindByUserId(userId)
}

func (u *CustomerUsecase) FindByTenantId(tenantId uuid.UUID) ([]*entity.Customer, error) {
	return u.customerRepository.FindByTenantId(tenantId)
}
