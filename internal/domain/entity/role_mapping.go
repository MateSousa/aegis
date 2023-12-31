package entity

import (
	"github.com/MateSousa/aegis/internal/domain/common"
	"github.com/google/uuid"
)

type RoleMapping struct {
	common.Model
	UserID   uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	RoleID   uuid.UUID `gorm:"type:uuid;not null" json:"role_id"`
	TenantID uuid.UUID `gorm:"type:uuid;not null" json:"tenant_id"`
}
