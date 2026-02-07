package auth

import (
	"errors"
	"testing"

	"businessense/internal/domain"
	"businessense/pkg/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// MockUserRepository is a mock implementation of repository.UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByID(id uint) (*domain.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) FindByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestRegister_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	cfg := &config.Config{JWTSecret: "test-secret"}
	service := NewService(mockRepo, cfg)

	email := "test@example.com"
	password := "password123"
	firstName := "John"
	lastName := "Doe"

	// User does not exist
	mockRepo.On("FindByEmail", email).Return(nil, nil)
	// Create user succeeds
	mockRepo.On("Create", mock.AnythingOfType("*domain.User")).Return(nil)

	user, token, err := service.Register(email, password, firstName, lastName)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, firstName, user.FirstName)
	assert.Equal(t, lastName, user.LastName)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestRegister_UserExists(t *testing.T) {
	mockRepo := new(MockUserRepository)
	cfg := &config.Config{JWTSecret: "test-secret"}
	service := NewService(mockRepo, cfg)

	email := "existing@example.com"
	existingUser := &domain.User{
		ID:    1,
		Email: email,
	}

	// User already exists
	mockRepo.On("FindByEmail", email).Return(existingUser, nil)

	user, token, err := service.Register(email, "password123", "John", "Doe")

	assert.Error(t, err)
	assert.Equal(t, ErrUserExists, err)
	assert.Nil(t, user)
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestLogin_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	cfg := &config.Config{JWTSecret: "test-secret"}
	service := NewService(mockRepo, cfg)

	email := "test@example.com"
	password := "password123"

	// Hash the password like the service would
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	existingUser := &domain.User{
		ID:        1,
		Email:     email,
		Password:  string(hashedPassword),
		FirstName: "John",
		LastName:  "Doe",
	}

	mockRepo.On("FindByEmail", email).Return(existingUser, nil)

	user, token, err := service.Login(email, password)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, email, user.Email)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestLogin_WrongPassword(t *testing.T) {
	mockRepo := new(MockUserRepository)
	cfg := &config.Config{JWTSecret: "test-secret"}
	service := NewService(mockRepo, cfg)

	email := "test@example.com"
	correctPassword := "password123"
	wrongPassword := "wrongpassword"

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(correctPassword), bcrypt.DefaultCost)

	existingUser := &domain.User{
		ID:        1,
		Email:     email,
		Password:  string(hashedPassword),
		FirstName: "John",
		LastName:  "Doe",
	}

	mockRepo.On("FindByEmail", email).Return(existingUser, nil)

	user, token, err := service.Login(email, wrongPassword)

	assert.Error(t, err)
	assert.Equal(t, ErrInvalidCredentials, err)
	assert.Nil(t, user)
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestLogin_UserNotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	cfg := &config.Config{JWTSecret: "test-secret"}
	service := NewService(mockRepo, cfg)

	email := "notfound@example.com"

	// User not found - return nil user and an error
	mockRepo.On("FindByEmail", email).Return(nil, errors.New("user not found"))

	user, token, err := service.Login(email, "password123")

	assert.Error(t, err)
	assert.Equal(t, ErrInvalidCredentials, err)
	assert.Nil(t, user)
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}
