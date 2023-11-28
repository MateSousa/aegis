package entity

import (
	"github.com/MateSousa/aegis/internal/domain/common"
	"github.com/google/uuid"
)

type Customer struct {
	common.Model
	UserID   uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	User     User      `gorm:"foreignKey:UserID" json:"user"`
	TenantID uuid.UUID `gorm:"type:uuid;not null" json:"tenant_id"`
	Tenant   Tenant    `gorm:"foreignKey:TenantID" json:"tenant"`
}
