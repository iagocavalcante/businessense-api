package postgres

import (
	"businessense/internal/domain"
	"businessense/internal/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
	model := &User{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
	}
	if err := r.db.Create(model).Error; err != nil {
		return err
	}
	user.ID = model.ID
	user.CreatedAt = model.CreatedAt
	user.UpdatedAt = model.UpdatedAt
	return nil
}

func (r *userRepository) FindByID(id uint) (*domain.User, error) {
	var model User
	if err := r.db.First(&model, id).Error; err != nil {
		return nil, err
	}
	return toDomainUser(&model), nil
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var model User
	if err := r.db.Where("email = ?", email).First(&model).Error; err != nil {
		return nil, err
	}
	return toDomainUser(&model), nil
}

func toDomainUser(m *User) *domain.User {
	return &domain.User{
		ID:        m.ID,
		Email:     m.Email,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Password:  m.Password,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
