package usecase

import (
	"app/internal/domain/entity"
	"app/internal/entity"

	"golang.org/x/mod/sumdb/storage"
)

// Прописываем интерфейс для методов связанных с репозиторием
// Прописываем там, где используем
type UserStorage interface {
	Create(user entity.User) entity.User
	GetOne(id string) entity.User
	GetAll() []entity.User
	Delete(entity.User) error
}

type userUseCase struct {
	storage UserStorage
}

// Инициализация usecase, прокидываем ему интерфейс
func NewUserUseCase(storage UserStorage) *userUseCase {
	return &userUseCase{storage: storage}
}

// Прописываем методы которые нужны для userUseCase
// Используем в них методы репозитория

func (uc userUseCase) Create(user entity.User) entity.User {
	return uc.storage.Create(user)
}

func (uc userUseCase) Delete(user entity.User) error {
	return uc.storage.Delete(user)
}
