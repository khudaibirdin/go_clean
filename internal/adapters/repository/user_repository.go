package repository

import (
	"app/internal/domain/entity"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	GetOne(id string) (*entity.User, error)
	GetAll() (*[]entity.User, error)
	Delete(entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

// Создание нового репозитория
// Не знаю, нужно оно здесь или нет
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Delete(user entity.User) error {
	return r.db.Delete(&user).Error
} 

func (r *userRepository) GetOne(id string) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.First(&user, id).Error; err != nil {
		return &entity.User{}, err
	}
	return user, nil
}
func (r *userRepository) GetAll() (*[]entity.User, error) {
	users := new([]entity.User)
	if err := r.db.Find(&users).Error; err != nil {
		return &[]entity.User{}, err
	}
	return users, nil
}