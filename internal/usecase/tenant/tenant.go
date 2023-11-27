package tenant

import (
	"math/rand"

	"github.com/MateSousa/aegis/internal/domain/entity"
	"github.com/MateSousa/aegis/internal/repository/tenant"
	"github.com/google/uuid"
)

type ITenantUsecase interface {
	Create(tenant *entity.Tenant) error
	Update(tenant *entity.Tenant) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.Tenant, error)
	FindAll() ([]*entity.Tenant, error)
	FindByUserId(userId uuid.UUID) ([]*entity.Tenant, error)
}

type TenantUsecase struct {
	tenantRepository tenant.ITenantRepository
}

func NewTenantUsecase(tenantRepository tenant.ITenantRepository) ITenantUsecase {
	return &TenantUsecase{
		tenantRepository: tenantRepository,
	}
}

func (u *TenantUsecase) Create(tenant *entity.Tenant) error {
	isFirstTenant := IsFirstTenantForUser(u, tenant.UserId)
	if isFirstTenant {
		tenant.Name = GenerateRandomTenantName()
	}
	return u.tenantRepository.Create(tenant)
}

func (u *TenantUsecase) Update(tenant *entity.Tenant) error {
	return u.tenantRepository.Update(tenant)
}

func (u *TenantUsecase) Delete(id uuid.UUID) error {
	return u.tenantRepository.Delete(id)
}

func (u *TenantUsecase) FindById(id uuid.UUID) (*entity.Tenant, error) {
	return u.tenantRepository.FindById(id)
}

func (u *TenantUsecase) FindAll() ([]*entity.Tenant, error) {
	return u.tenantRepository.FindAll()
}

func (u *TenantUsecase) FindByUserId(userId uuid.UUID) ([]*entity.Tenant, error) {
	return u.tenantRepository.FindByUserId(userId)
}

func IsFirstTenantForUser(tenantUsecase ITenantUsecase, userId uuid.UUID) bool {
	tenants, err := tenantUsecase.FindByUserId(userId)
	if err != nil {
		return false
	}
	return len(tenants) == 0
}

func GenerateRandomTenantName() string {
	// generate a random string with max length of 8
	// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
	return "development-" + randStringBytes(8)
}

func randStringBytes(n int) string {
	// generate a random string with max length of 8
	// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
