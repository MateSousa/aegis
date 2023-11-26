package user

import (
	"github.com/MateSousa/aegis/internal/domain/entity"
	"github.com/google/uuid"
)

type IUserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUserByID(id uuid.UUID) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(id uuid.UUID) error
	GetUsers() ([]*entity.User, error)
	EmailExists(email string) (bool, error)
}
