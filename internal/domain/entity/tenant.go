package entity

import (
	"github.com/MateSousa/aegis/internal/domain/common"
	"github.com/google/uuid"
)

type Tenant struct {
	common.Model
	Name   string    `json:"name" validate:"required,alphanum,min=3,max=255"`
	UserId uuid.UUID `json:"user_id" validate:"required"`
	User   User      `json:"user"`
}
