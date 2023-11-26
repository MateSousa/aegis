package entity

import "github.com/MateSousa/aegis/internal/domain/common"

type Role struct {
	common.Model
	Name string `json:"name"`
}
