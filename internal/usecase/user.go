package usecase

import (
	"log/slog"

	"github.com/CAMELNINJA/bot-bet.git/internal/models"
)

type userRepo interface {
	// GetByID returns user by id
	GetByID(id int) (*models.User, error)
	// GetByTelegramID returns user by telegram id
	GetByTelegramID(id int) (*models.User, error)
	// Create creates new user
	Create(user *models.User) error
	// Update updates user
	Update(user *models.User) error
	// Delete deletes user
	Delete(id int) error
}

type UserUsecase struct {
	log  *slog.Logger
	repo userRepo
}

func NewUserUsecase(log *slog.Logger, repo userRepo) *UserUsecase {
	return &UserUsecase{
		log:  log,
		repo: repo,
	}
}

// GetByID returns user by id
func (u *UserUsecase) GetByID(id int) (*models.User, error) {
	return u.repo.GetByID(id)
}

// GetByTelegramID returns user by telegram id
func (u *UserUsecase) GetByTelegramID(id int) (*models.User, error) {
	return u.repo.GetByTelegramID(id)
}

// Create creates new user
func (u *UserUsecase) Create(user *models.User) error {
	return u.repo.Create(user)
}

// Update updates user
func (u *UserUsecase) Update(user *models.User) error {
	return u.repo.Update(user)
}

// Delete deletes user
func (u *UserUsecase) Delete(id int) error {
	return u.repo.Delete(id)
}
