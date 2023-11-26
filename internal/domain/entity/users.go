package entity

import (
	"errors"

	"github.com/MateSousa/aegis/internal/domain/common"
)

type User struct {
	common.Model
	Name            string `json:"name" validate:"required,alphanum,min=3,max=255"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6,max=255"`
	IsEmailVerified bool   `json:"is_email_verified"`
	LastLogin       string `json:"last_login"`
}

var ErrEmailExists = errors.New("email already exists")
