package entity

import (
	"github.com/MateSousa/aegis/internal/domain/common"
	"github.com/google/uuid"
)

type Tenant struct {
	common.Model
	Name         string    `json:"name" validate:"required,alphanum,min=3,max=255"`
	ClientId     uuid.UUID `gorm:"type:uuid;not null;uniqueIndex" json:"client_id"`
	ClientSecret string    `gorm:"not null;uniqueIndex" json:"client_secret"`
	OwnerId      uuid.UUID `gorm:"type:uuid;not null" json:"owner_id"`
	Owner        User      `gorm:"foreignKey:OwnerId" json:"owner"`
}
