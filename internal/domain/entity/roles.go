package entity

import "github.com/MateSousa/internal/domain/common"

type Role struct {
	common.Model
	Name string `json:"name"`
}
