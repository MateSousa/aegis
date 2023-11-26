package entity

import "github.com/MateSousa/internal/domain/common"

type User struct {
	common.Model
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"-"`
	IsEmailVerified bool   `json:"is_email_verified"`
	LastLogin       string `json:"last_login"`
}
