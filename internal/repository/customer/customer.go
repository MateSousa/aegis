package customer

import (
	"time"

	"github.com/MateSousa/aegis/internal/domain/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(connection *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		db: connection,
	}
}

func (r *CustomerRepository) Create(customer *entity.Customer) (*entity.Customer, error) {
	if err := r.db.Create(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *CustomerRepository) Update(customer *entity.Customer) (*entity.Customer, error) {
	if err := r.db.Save(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *CustomerRepository) Delete(id uuid.UUID) error {
	if err := r.db.Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepository) FindById(id uuid.UUID) (*entity.Customer, error) {
	customer := &entity.Customer{}
	if err := r.db.Preload("User").First(customer, id).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *CustomerRepository) FindByUserId(userId uuid.UUID) (*entity.Customer, error) {
	customer := &entity.Customer{}
	if err := r.db.Preload("User").Where("user_id = ?", userId).First(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *CustomerRepository) FindByTenantId(tenantId uuid.UUID) ([]*entity.Customer, error) {
	customers := []*entity.Customer{}
	if err := r.db.Preload("User").Where("tenant_id = ?", tenantId).Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}
