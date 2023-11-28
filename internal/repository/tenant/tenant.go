package tenant

import (
	"github.com/MateSousa/aegis/internal/domain/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) ITenantRepository {
	return &TenantRepository{db}
}

func (r *TenantRepository) Create(tenant *entity.Tenant) error {
	return r.db.Create(tenant).Error
}

func (r *TenantRepository) Update(tenant *entity.Tenant) error {
	return r.db.Save(tenant).Error
}

func (r *TenantRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entity.Tenant{}, id).Error
}

func (r *TenantRepository) FindById(id uuid.UUID) (*entity.Tenant, error) {
	tenant := &entity.Tenant{}
	err := r.db.Preload("Owner").First(tenant, id).Error
	if err != nil {
		return nil, err
	}
	return tenant, nil
}

func (r *TenantRepository) FindAll() ([]*entity.Tenant, error) {
	tenants := []*entity.Tenant{}
	err := r.db.Preload("Owner").Find(&tenants).Error
	if err != nil {
		return nil, err
	}
	return tenants, nil
}

func (r *TenantRepository) FindByUserId(userId uuid.UUID) ([]*entity.Tenant, error) {
	tenants := []*entity.Tenant{}
	err := r.db.Preload("Owner").Where("user_id = ?", userId).Find(&tenants).Error
	if err != nil {
		return nil, err
	}
	return tenants, nil
}

func (r *TenantRepository) FindByClientId(clientId uuid.UUID) (*entity.Tenant, error) {
	tenant := &entity.Tenant{}
	err := r.db.Preload("Owner").Where("client_id = ?", clientId).First(tenant).Error
	if err != nil {
		return nil, err
	}
	return tenant, nil
}
