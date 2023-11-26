package user

import (
	"github.com/MateSousa/aegis/internal/domain/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(connection *gorm.DB) *UserRepository {
	return &UserRepository{
		db: connection,
	}
}

func (r *UserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByID(id uuid.UUID) (*entity.User, error) {
	user := &entity.User{}
	if err := r.db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user *entity.User) (*entity.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) DeleteUser(id uuid.UUID) error {
	if err := r.db.Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUsers() ([]*entity.User, error) {
	users := []*entity.User{}
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) EmailExists(email string) (bool, error) {
	var count int64
	if err := r.db.Model(&entity.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
