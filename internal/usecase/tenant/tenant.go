package tenant

import (
	"encoding/base64"
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
	FindByClientId(clientId uuid.UUID) (*entity.Tenant, error)
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
	isFirstTenant := IsFirstTenantForUser(u, tenant.OwnerId)
	if isFirstTenant {
		tenant.Name = GenerateRandomTenantName()
	}
	tenant.ClientId = u.GenerateRandomClientId()
	tenant.ClientSecret = u.Base64Encode(u.GenerateRandomClientSecret())

	err := u.tenantRepository.Create(tenant)
	if err != nil {
		return err
	}
	return nil
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
	return "development-" + randStringBytes(8)
}

func randStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (u *TenantUsecase) GenerateRandomClientId() uuid.UUID {
	return uuid.New()
}

func (u *TenantUsecase) GenerateRandomClientSecret() string {
	return randStringBytes(32)
}

func (u *TenantUsecase) Base64Encode(secret string) string {
	return base64.StdEncoding.EncodeToString([]byte(secret))
}

func (u *TenantUsecase) FindByClientId(clientId uuid.UUID) (*entity.Tenant, error) {
	return u.tenantRepository.FindByClientId(clientId)
}
