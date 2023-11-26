package entity

import (
	"github.com/MateSousa/aegis/internal/domain/common"
	"github.com/google/uuid"
)

type RoleMapping struct {
	common.Model
	UserID uuid.UUID `json:"user_id"`
	RoleID uuid.UUID `json:"role_id"`
}
