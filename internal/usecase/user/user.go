package usecase

import (
	"github.com/MateSousa/aegis/internal/domain/entity"
	"github.com/MateSousa/aegis/internal/domain/repository/user"
	"github.com/google/uuid"
)

type IUserUsecase interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUserByID(id uuid.UUID) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(id uuid.UUID) error
	GetUsers() ([]*entity.User, error)
}

type userUseCase struct {
	userRepo user.IUserRepository
}

func NewUserUsecase(userRepo user.IUserRepository) IUserUsecase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (u *userUseCase) CreateUser(user *entity.User) (*entity.User, error) {
	return u.userRepo.CreateUser(user)
}

func (u *userUseCase) GetUserByID(id uuid.UUID) (*entity.User, error) {
	return u.userRepo.GetUserByID(id)
}

func (u *userUseCase) GetUserByEmail(email string) (*entity.User, error) {
	return u.userRepo.GetUserByEmail(email)
}

func (u *userUseCase) UpdateUser(user *entity.User) (*entity.User, error) {
	return u.userRepo.UpdateUser(user)
}

func (u *userUseCase) DeleteUser(id uuid.UUID) error {
	return u.userRepo.DeleteUser(id)
}

func (u *userUseCase) GetUsers() ([]*entity.User, error) {
	return u.userRepo.GetUsers()
}
