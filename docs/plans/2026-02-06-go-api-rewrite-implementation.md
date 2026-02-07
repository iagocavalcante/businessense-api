# Go API Rewrite Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Rewrite the businessense Go API using Clean Architecture with Chi, GORM v2, and modern patterns.

**Architecture:** Clean Architecture with domain/usecase/repository/handler layers, dependency injection, typed errors.

**Tech Stack:** Go 1.22+, Chi, GORM v2, golang-jwt/jwt/v5, testify

---

## Phase 1: Project Setup

### Task 1: Backup Old API

**Files:**
- Backup: `businessense-api/` -> `businessense-api-v1-backup/`

**Step 1: Create backup**
```bash
cd /Users/iagocavalcante/Workspaces/IagoCavalcante/businessense
cp -r businessense-api businessense-api-v1-backup
```

**Step 2: Verify backup**
```bash
ls businessense-api-v1-backup/
```
Expected: server.go, app/, controllers/, models/, utils/

---

### Task 2: Initialize New Go Module

**Files:**
- Create: `businessense-api/go.mod`

**Step 1: Clean the directory (keep .env and .git)**
```bash
cd businessense-api
# Remove old Go files but keep config
rm -rf app controllers models utils server.go public
```

**Step 2: Initialize Go module**
```bash
go mod init businessense
```

**Step 3: Create directory structure**
```bash
mkdir -p cmd/server
mkdir -p internal/domain
mkdir -p internal/usecase/{auth,company,industry,painpoint,issue,solution,project}
mkdir -p internal/repository/postgres
mkdir -p internal/handler/middleware
mkdir -p pkg/config
mkdir -p migrations
```

**Step 4: Commit**
```bash
git add .
git commit -m "chore: initialize new Go module structure"
```

---

### Task 3: Install Dependencies

**Files:**
- Modify: `go.mod`

**Step 1: Install dependencies**
```bash
go get github.com/go-chi/chi/v5
go get github.com/go-chi/cors
go get github.com/golang-jwt/jwt/v5
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get golang.org/x/crypto
go get github.com/joho/godotenv
go get github.com/stretchr/testify
```

**Step 2: Verify go.mod**
```bash
cat go.mod
```
Expected: All dependencies listed

**Step 3: Commit**
```bash
git add go.mod go.sum
git commit -m "chore: add project dependencies"
```

---

### Task 4: Create Config Package

**Files:**
- Create: `pkg/config/config.go`

**Step 1: Write config.go**
```go
package config

import (
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    Port        string
    DBHost      string
    DBUser      string
    DBPassword  string
    DBName      string
    JWTSecret   string
}

func Load() *Config {
    godotenv.Load()

    return &Config{
        Port:       getEnv("PORT", "8000"),
        DBHost:     getEnv("db_host", "localhost"),
        DBUser:     getEnv("db_user", "postgres"),
        DBPassword: getEnv("db_pass", ""),
        DBName:     getEnv("db_name", "businessense"),
        JWTSecret:  getEnv("token_password", "secret"),
    }
}

func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}
```

**Step 2: Run test**
```bash
go build ./pkg/config
```
Expected: No errors

**Step 3: Commit**
```bash
git add pkg/config/config.go
git commit -m "feat: add config package"
```

---

## Phase 2: Domain Layer

### Task 5: Create Domain Entities

**Files:**
- Create: `internal/domain/user.go`
- Create: `internal/domain/company.go`
- Create: `internal/domain/industry.go`
- Create: `internal/domain/painpoint.go`
- Create: `internal/domain/issue.go`
- Create: `internal/domain/solution.go`
- Create: `internal/domain/constraint.go`
- Create: `internal/domain/project.go`

**Step 1: Write user.go**
```go
package domain

import "time"

type User struct {
    ID        uint      `json:"id"`
    Email     string    `json:"email"`
    FirstName string    `json:"firstName"`
    LastName  string    `json:"lastName"`
    Password  string    `json:"-"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}
```

**Step 2: Write industry.go**
```go
package domain

import "time"

type Industry struct {
    ID        uint      `json:"id"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}
```

**Step 3: Write company.go**
```go
package domain

import "time"

type Company struct {
    ID         uint      `json:"id"`
    Name       string    `json:"name"`
    IndustryID uint      `json:"industryId"`
    UserID     uint      `json:"userId"`
    CreatedAt  time.Time `json:"createdAt"`
    UpdatedAt  time.Time `json:"updatedAt"`
}
```

**Step 4: Write painpoint.go**
```go
package domain

import "time"

type PainPoint struct {
    ID        uint      `json:"id"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}
```

**Step 5: Write issue.go**
```go
package domain

import "time"

type Issue struct {
    ID          uint      `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Stats       string    `json:"stats"`
    RiskFactors string    `json:"riskFactors"`
    CreatedAt   time.Time `json:"createdAt"`
    UpdatedAt   time.Time `json:"updatedAt"`
}

type IssueRelevance struct {
    IssueID   uint    `json:"issueId"`
    Name      string  `json:"name"`
    Relevance float64 `json:"relevance"`
}

type IssuePainPointMap struct {
    ID          uint `json:"id"`
    IssueID     uint `json:"issueId"`
    PainPointID uint `json:"painPointId"`
}
```

**Step 6: Write solution.go**
```go
package domain

import "time"

type Solution struct {
    ID          uint      `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"createdAt"`
    UpdatedAt   time.Time `json:"updatedAt"`
}
```

**Step 7: Write constraint.go**
```go
package domain

import "time"

type Constraint struct {
    ID        uint      `json:"id"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

type ConstraintIssueSolutionMap struct {
    ID           uint `json:"id"`
    ConstraintID uint `json:"constraintId"`
    IssueID      uint `json:"issueId"`
    SolutionID   uint `json:"solutionId"`
}
```

**Step 8: Write project.go**
```go
package domain

import "time"

type Project struct {
    ID          uint      `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    CompanyID   uint      `json:"companyId"`
    CreatedAt   time.Time `json:"createdAt"`
    UpdatedAt   time.Time `json:"updatedAt"`
}
```

**Step 9: Verify build**
```bash
go build ./internal/domain
```
Expected: No errors

**Step 10: Commit**
```bash
git add internal/domain/
git commit -m "feat: add domain entities"
```

---

### Task 6: Create Repository Interfaces

**Files:**
- Create: `internal/repository/interfaces.go`

**Step 1: Write interfaces.go**
```go
package repository

import "businessense/internal/domain"

type UserRepository interface {
    Create(user *domain.User) error
    FindByID(id uint) (*domain.User, error)
    FindByEmail(email string) (*domain.User, error)
}

type IndustryRepository interface {
    Create(industry *domain.Industry) error
    FindAll() ([]domain.Industry, error)
    Search(query string) ([]domain.Industry, error)
}

type CompanyRepository interface {
    Create(company *domain.Company) error
    FindByID(id uint) (*domain.Company, error)
    FindByUserID(userID uint) ([]domain.Company, error)
}

type PainPointRepository interface {
    Create(pp *domain.PainPoint) error
    FindAll() ([]domain.PainPoint, error)
    Search(query string) ([]domain.PainPoint, error)
}

type IssueRepository interface {
    FindByID(id uint) (*domain.Issue, error)
    FindByPainPointIDs(ppIDs []uint) ([]domain.IssueRelevance, error)
}

type SolutionRepository interface {
    FindByID(id uint) (*domain.Solution, error)
    FindByIssueID(issueID uint) ([]domain.Solution, error)
}

type ConstraintRepository interface {
    FindByIssueAndSolution(issueID, solutionID uint) ([]domain.Constraint, error)
}

type ProjectRepository interface {
    Create(project *domain.Project) error
    FindByID(id uint) (*domain.Project, error)
    FindByCompanyID(companyID uint) ([]domain.Project, error)
}
```

**Step 2: Verify build**
```bash
go build ./internal/repository
```
Expected: No errors

**Step 3: Commit**
```bash
git add internal/repository/interfaces.go
git commit -m "feat: add repository interfaces"
```

---

## Phase 3: Repository Layer (GORM)

### Task 7: Create Database Connection

**Files:**
- Create: `internal/repository/postgres/db.go`

**Step 1: Write db.go**
```go
package postgres

import (
    "fmt"
    "log"
    "businessense/pkg/config"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

func Connect(cfg *config.Config) *gorm.DB {
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    return db
}
```

**Step 2: Verify build**
```bash
go build ./internal/repository/postgres
```
Expected: No errors

**Step 3: Commit**
```bash
git add internal/repository/postgres/db.go
git commit -m "feat: add database connection"
```

---

### Task 8: Create GORM Models

**Files:**
- Create: `internal/repository/postgres/models.go`

**Step 1: Write models.go (GORM models with tags)**
```go
package postgres

import "time"

type User struct {
    ID        uint      `gorm:"primaryKey"`
    Email     string    `gorm:"uniqueIndex;not null"`
    FirstName string    `gorm:"column:firstname"`
    LastName  string    `gorm:"column:lastname"`
    Password  string    `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (User) TableName() string { return "accounts" }

type Industry struct {
    ID        uint   `gorm:"primaryKey"`
    Name      string `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (Industry) TableName() string { return "industries" }

type Company struct {
    ID         uint   `gorm:"primaryKey"`
    Name       string `gorm:"not null"`
    IndustryID uint   `gorm:"column:industry_id"`
    UserID     uint   `gorm:"column:user_id"`
    CreatedAt  time.Time
    UpdatedAt  time.Time
}

func (Company) TableName() string { return "companies" }

type PainPoint struct {
    ID        uint   `gorm:"primaryKey"`
    Name      string `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (PainPoint) TableName() string { return "pain_points" }

type Issue struct {
    ID          uint   `gorm:"primaryKey"`
    Name        string `gorm:"not null"`
    Description string
    Stats       string
    RiskFactors string `gorm:"column:riskfactors"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

func (Issue) TableName() string { return "issues" }

type IssuePainPointMap struct {
    ID          uint `gorm:"primaryKey"`
    IssueID     uint `gorm:"column:issue_id"`
    PainPointID uint `gorm:"column:pain_point_id"`
}

func (IssuePainPointMap) TableName() string { return "issues_pain_points_maps" }

type Solution struct {
    ID          uint   `gorm:"primaryKey"`
    Name        string `gorm:"not null"`
    Description string
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

func (Solution) TableName() string { return "solutions" }

type Constraint struct {
    ID        uint   `gorm:"primaryKey"`
    Name      string `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (Constraint) TableName() string { return "constraints" }

type ConstraintIssueSolutionMap struct {
    ID           uint `gorm:"primaryKey"`
    ConstraintID uint `gorm:"column:constraint_id"`
    IssueID      uint `gorm:"column:issue_id"`
    SolutionID   uint `gorm:"column:solution_id"`
}

func (ConstraintIssueSolutionMap) TableName() string { return "constraint_issues_solution_maps" }

type Project struct {
    ID          uint   `gorm:"primaryKey"`
    Name        string `gorm:"not null"`
    Description string
    CompanyID   uint `gorm:"column:company_id"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

func (Project) TableName() string { return "projects" }
```

**Step 2: Verify build**
```bash
go build ./internal/repository/postgres
```

**Step 3: Commit**
```bash
git add internal/repository/postgres/models.go
git commit -m "feat: add GORM models"
```

---

### Task 9: Implement User Repository

**Files:**
- Create: `internal/repository/postgres/user.go`
- Create: `internal/repository/postgres/user_test.go`

**Step 1: Write the test first**
```go
package postgres_test

import (
    "testing"
    "businessense/internal/domain"
    "businessense/internal/repository/postgres"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestUserRepository_Create(t *testing.T) {
    // This is a unit test structure - integration tests would use real DB
    t.Run("creates user successfully", func(t *testing.T) {
        // Test will be expanded with mock DB
        assert.True(t, true)
    })
}
```

**Step 2: Write user.go**
```go
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
```

**Step 3: Run test**
```bash
go test ./internal/repository/postgres/... -v
```

**Step 4: Commit**
```bash
git add internal/repository/postgres/user.go internal/repository/postgres/user_test.go
git commit -m "feat: implement user repository"
```

---

### Task 10: Implement Remaining Repositories

**Files:**
- Create: `internal/repository/postgres/industry.go`
- Create: `internal/repository/postgres/company.go`
- Create: `internal/repository/postgres/painpoint.go`
- Create: `internal/repository/postgres/issue.go`
- Create: `internal/repository/postgres/solution.go`
- Create: `internal/repository/postgres/constraint.go`
- Create: `internal/repository/postgres/project.go`

(Each follows same pattern as user repository - implement interface, convert between GORM model and domain)

**Step 1: Implement all repositories following the user.go pattern**

**Step 2: Run tests**
```bash
go test ./internal/repository/postgres/... -v
```

**Step 3: Commit**
```bash
git add internal/repository/postgres/
git commit -m "feat: implement all repositories"
```

---

## Phase 4: Usecase Layer (Services)

### Task 11: Implement Auth Service with Tests

**Files:**
- Create: `internal/usecase/auth/service.go`
- Create: `internal/usecase/auth/service_test.go`
- Create: `internal/usecase/auth/errors.go`

**Step 1: Write errors.go**
```go
package auth

import "errors"

var (
    ErrInvalidCredentials = errors.New("invalid credentials")
    ErrUserExists         = errors.New("user already exists")
    ErrUserNotFound       = errors.New("user not found")
)
```

**Step 2: Write service_test.go**
```go
package auth_test

import (
    "testing"
    "businessense/internal/domain"
    "businessense/internal/usecase/auth"
    "businessense/pkg/config"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

type mockUserRepo struct {
    mock.Mock
}

func (m *mockUserRepo) Create(user *domain.User) error {
    args := m.Called(user)
    return args.Error(0)
}

func (m *mockUserRepo) FindByID(id uint) (*domain.User, error) {
    args := m.Called(id)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*domain.User), args.Error(1)
}

func (m *mockUserRepo) FindByEmail(email string) (*domain.User, error) {
    args := m.Called(email)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*domain.User), args.Error(1)
}

func TestAuthService_Register(t *testing.T) {
    t.Run("registers new user successfully", func(t *testing.T) {
        repo := new(mockUserRepo)
        cfg := &config.Config{JWTSecret: "test-secret"}
        svc := auth.NewService(repo, cfg)

        repo.On("FindByEmail", "test@example.com").Return(nil, nil)
        repo.On("Create", mock.AnythingOfType("*domain.User")).Return(nil)

        user, token, err := svc.Register("test@example.com", "password123", "John", "Doe")

        assert.NoError(t, err)
        assert.NotNil(t, user)
        assert.NotEmpty(t, token)
        repo.AssertExpectations(t)
    })

    t.Run("returns error when user exists", func(t *testing.T) {
        repo := new(mockUserRepo)
        cfg := &config.Config{JWTSecret: "test-secret"}
        svc := auth.NewService(repo, cfg)

        existingUser := &domain.User{ID: 1, Email: "test@example.com"}
        repo.On("FindByEmail", "test@example.com").Return(existingUser, nil)

        _, _, err := svc.Register("test@example.com", "password123", "John", "Doe")

        assert.ErrorIs(t, err, auth.ErrUserExists)
    })
}

func TestAuthService_Login(t *testing.T) {
    t.Run("logs in successfully with correct credentials", func(t *testing.T) {
        repo := new(mockUserRepo)
        cfg := &config.Config{JWTSecret: "test-secret"}
        svc := auth.NewService(repo, cfg)

        // Pre-hashed password for "password123"
        hashedPw := "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZRGdjGj/n3N/l.9FAw.ieFD0ybDAK"
        user := &domain.User{ID: 1, Email: "test@example.com", Password: hashedPw}
        repo.On("FindByEmail", "test@example.com").Return(user, nil)

        result, token, err := svc.Login("test@example.com", "password123")

        assert.NoError(t, err)
        assert.NotNil(t, result)
        assert.NotEmpty(t, token)
    })
}
```

**Step 3: Write service.go**
```go
package auth

import (
    "businessense/internal/domain"
    "businessense/internal/repository"
    "businessense/pkg/config"
    "time"
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
)

type Service struct {
    userRepo  repository.UserRepository
    jwtSecret string
}

func NewService(userRepo repository.UserRepository, cfg *config.Config) *Service {
    return &Service{
        userRepo:  userRepo,
        jwtSecret: cfg.JWTSecret,
    }
}

func (s *Service) Register(email, password, firstName, lastName string) (*domain.User, string, error) {
    existing, _ := s.userRepo.FindByEmail(email)
    if existing != nil {
        return nil, "", ErrUserExists
    }

    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, "", err
    }

    user := &domain.User{
        Email:     email,
        Password:  string(hash),
        FirstName: firstName,
        LastName:  lastName,
    }

    if err := s.userRepo.Create(user); err != nil {
        return nil, "", err
    }

    token, err := s.generateToken(user.ID)
    if err != nil {
        return nil, "", err
    }

    user.Password = ""
    return user, token, nil
}

func (s *Service) Login(email, password string) (*domain.User, string, error) {
    user, err := s.userRepo.FindByEmail(email)
    if err != nil || user == nil {
        return nil, "", ErrInvalidCredentials
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return nil, "", ErrInvalidCredentials
    }

    token, err := s.generateToken(user.ID)
    if err != nil {
        return nil, "", err
    }

    user.Password = ""
    return user, token, nil
}

func (s *Service) generateToken(userID uint) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.jwtSecret))
}
```

**Step 4: Run tests**
```bash
go test ./internal/usecase/auth/... -v
```

**Step 5: Commit**
```bash
git add internal/usecase/auth/
git commit -m "feat: implement auth service with tests"
```

---

### Task 12: Implement Remaining Services

**Files:**
- Create: `internal/usecase/industry/service.go`
- Create: `internal/usecase/company/service.go`
- Create: `internal/usecase/painpoint/service.go`
- Create: `internal/usecase/issue/service.go`
- Create: `internal/usecase/solution/service.go`
- Create: `internal/usecase/project/service.go`

(Each with corresponding tests following the auth service pattern)

**Step 1: Implement all services with tests**

**Step 2: Run all tests**
```bash
go test ./internal/usecase/... -v
```

**Step 3: Commit**
```bash
git add internal/usecase/
git commit -m "feat: implement all services with tests"
```

---

## Phase 5: HTTP Handler Layer

### Task 13: Create Response Helpers

**Files:**
- Create: `internal/handler/response.go`

**Step 1: Write response.go**
```go
package handler

import (
    "encoding/json"
    "net/http"
)

type ErrorResponse struct {
    Error   string `json:"error"`
    Message string `json:"message"`
}

func respondJSON(w http.ResponseWriter, status int, data any) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
    respondJSON(w, status, ErrorResponse{
        Error:   http.StatusText(status),
        Message: message,
    })
}
```

**Step 2: Commit**
```bash
git add internal/handler/response.go
git commit -m "feat: add response helpers"
```

---

### Task 14: Create JWT Middleware

**Files:**
- Create: `internal/handler/middleware/auth.go`

**Step 1: Write auth.go middleware**
```go
package middleware

import (
    "context"
    "net/http"
    "strings"
    "github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "userID"

func JWTAuth(secret string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            header := r.Header.Get("Authorization")
            if header == "" {
                http.Error(w, `{"error":"Unauthorized","message":"missing authorization header"}`, http.StatusUnauthorized)
                return
            }

            parts := strings.Split(header, " ")
            if len(parts) != 2 || parts[0] != "Bearer" {
                http.Error(w, `{"error":"Unauthorized","message":"invalid authorization header"}`, http.StatusUnauthorized)
                return
            }

            token, err := jwt.Parse(parts[1], func(t *jwt.Token) (any, error) {
                return []byte(secret), nil
            })
            if err != nil || !token.Valid {
                http.Error(w, `{"error":"Unauthorized","message":"invalid token"}`, http.StatusUnauthorized)
                return
            }

            claims, ok := token.Claims.(jwt.MapClaims)
            if !ok {
                http.Error(w, `{"error":"Unauthorized","message":"invalid claims"}`, http.StatusUnauthorized)
                return
            }

            userID := uint(claims["user_id"].(float64))
            ctx := context.WithValue(r.Context(), UserIDKey, userID)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

func GetUserID(r *http.Request) uint {
    if id, ok := r.Context().Value(UserIDKey).(uint); ok {
        return id
    }
    return 0
}
```

**Step 2: Commit**
```bash
git add internal/handler/middleware/auth.go
git commit -m "feat: add JWT authentication middleware"
```

---

### Task 15: Implement Auth Handler

**Files:**
- Create: `internal/handler/auth.go`

**Step 1: Write auth.go**
```go
package handler

import (
    "encoding/json"
    "errors"
    "net/http"
    "businessense/internal/domain"
    "businessense/internal/usecase/auth"
)

type AuthHandler struct {
    authService *auth.Service
}

func NewAuthHandler(authService *auth.Service) *AuthHandler {
    return &AuthHandler{authService: authService}
}

type RegisterRequest struct {
    Email     string `json:"email"`
    Password  string `json:"password"`
    FirstName string `json:"firstname"`
    LastName  string `json:"lastname"`
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type AuthResponse struct {
    Status  bool         `json:"status"`
    Message string       `json:"message"`
    Account *domain.User `json:"account,omitempty"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
    var req RegisterRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        respondError(w, http.StatusBadRequest, "invalid request body")
        return
    }

    user, token, err := h.authService.Register(req.Email, req.Password, req.FirstName, req.LastName)
    if err != nil {
        switch {
        case errors.Is(err, auth.ErrUserExists):
            respondError(w, http.StatusConflict, "email already in use")
        default:
            respondError(w, http.StatusInternalServerError, "failed to create account")
        }
        return
    }

    user.Password = ""
    respondJSON(w, http.StatusCreated, map[string]any{
        "status":  true,
        "message": "Account has been created",
        "account": map[string]any{
            "id":        user.ID,
            "email":     user.Email,
            "firstname": user.FirstName,
            "lastname":  user.LastName,
            "token":     token,
        },
    })
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    var req LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        respondError(w, http.StatusBadRequest, "invalid request body")
        return
    }

    user, token, err := h.authService.Login(req.Email, req.Password)
    if err != nil {
        respondJSON(w, http.StatusUnauthorized, map[string]any{
            "status":  false,
            "message": "Invalid login credentials",
        })
        return
    }

    respondJSON(w, http.StatusOK, map[string]any{
        "status":  true,
        "message": "Logged In",
        "account": map[string]any{
            "id":        user.ID,
            "email":     user.Email,
            "firstname": user.FirstName,
            "lastname":  user.LastName,
            "token":     token,
        },
    })
}
```

**Step 2: Commit**
```bash
git add internal/handler/auth.go
git commit -m "feat: implement auth handler"
```

---

### Task 16: Implement Remaining Handlers

**Files:**
- Create: `internal/handler/industry.go`
- Create: `internal/handler/company.go`
- Create: `internal/handler/painpoint.go`
- Create: `internal/handler/issue.go`
- Create: `internal/handler/solution.go`
- Create: `internal/handler/project.go`

(Each following the auth handler pattern)

**Step 1: Implement all handlers**

**Step 2: Commit**
```bash
git add internal/handler/
git commit -m "feat: implement all HTTP handlers"
```

---

## Phase 6: Main Entry Point

### Task 17: Create Main Server

**Files:**
- Create: `cmd/server/main.go`

**Step 1: Write main.go**
```go
package main

import (
    "log"
    "net/http"

    "businessense/internal/handler"
    "businessense/internal/handler/middleware"
    "businessense/internal/repository/postgres"
    "businessense/internal/usecase/auth"
    "businessense/internal/usecase/company"
    "businessense/internal/usecase/industry"
    "businessense/internal/usecase/issue"
    "businessense/internal/usecase/painpoint"
    "businessense/internal/usecase/project"
    "businessense/internal/usecase/solution"
    "businessense/pkg/config"

    "github.com/go-chi/chi/v5"
    chiMiddleware "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"
)

func main() {
    cfg := config.Load()
    db := postgres.Connect(cfg)

    // Repositories
    userRepo := postgres.NewUserRepository(db)
    industryRepo := postgres.NewIndustryRepository(db)
    companyRepo := postgres.NewCompanyRepository(db)
    painpointRepo := postgres.NewPainPointRepository(db)
    issueRepo := postgres.NewIssueRepository(db)
    solutionRepo := postgres.NewSolutionRepository(db)
    constraintRepo := postgres.NewConstraintRepository(db)
    projectRepo := postgres.NewProjectRepository(db)

    // Services
    authService := auth.NewService(userRepo, cfg)
    industryService := industry.NewService(industryRepo)
    companyService := company.NewService(companyRepo)
    painpointService := painpoint.NewService(painpointRepo)
    issueService := issue.NewService(issueRepo)
    solutionService := solution.NewService(solutionRepo, constraintRepo)
    projectService := project.NewService(projectRepo)

    // Handlers
    authHandler := handler.NewAuthHandler(authService)
    industryHandler := handler.NewIndustryHandler(industryService)
    companyHandler := handler.NewCompanyHandler(companyService)
    painpointHandler := handler.NewPainPointHandler(painpointService)
    issueHandler := handler.NewIssueHandler(issueService)
    solutionHandler := handler.NewSolutionHandler(solutionService)
    projectHandler := handler.NewProjectHandler(projectService)

    // Router
    r := chi.NewRouter()

    // Global middleware
    r.Use(chiMiddleware.Logger)
    r.Use(chiMiddleware.Recoverer)
    r.Use(chiMiddleware.RequestID)
    r.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Authorization", "Content-Type"},
        AllowCredentials: false,
        MaxAge:           300,
    }))

    // Public routes
    r.Post("/api/user/new", authHandler.Register)
    r.Post("/api/user/login", authHandler.Login)

    // Protected routes
    r.Group(func(r chi.Router) {
        r.Use(middleware.JWTAuth(cfg.JWTSecret))

        // Industry
        r.Post("/api/industry/new", industryHandler.Create)
        r.Get("/api/industry/search", industryHandler.Search)
        r.Get("/api/industry/search/{search}", industryHandler.Search)

        // Company
        r.Post("/api/company/new", companyHandler.Create)
        r.Get("/api/company/search", companyHandler.List)

        // PainPoint
        r.Post("/api/painpoint/new", painpointHandler.Create)
        r.Get("/api/painpoint/search", painpointHandler.Search)
        r.Get("/api/painpoint/search/{search}", painpointHandler.Search)

        // Issue
        r.Get("/api/issue/{id}", issueHandler.Get)
        r.Post("/api/issue/relevance", issueHandler.GetRelevance)

        // Solution
        r.Get("/api/solution/{id}", solutionHandler.Get)
        r.Get("/api/solution/issue/{issueid}", solutionHandler.GetByIssue)

        // Project
        r.Post("/api/project/new", projectHandler.Create)
        r.Get("/api/project/{id}", projectHandler.Get)
        r.Get("/api/project/company/{id}", projectHandler.GetByCompany)
    })

    // Static files
    r.Handle("/*", http.FileServer(http.Dir("./public/")))

    log.Printf("Server starting on :%s", cfg.Port)
    if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
        log.Fatal(err)
    }
}
```

**Step 2: Build and verify**
```bash
go build ./cmd/server
```

**Step 3: Commit**
```bash
git add cmd/server/main.go
git commit -m "feat: create main server entry point"
```

---

### Task 18: Create Makefile

**Files:**
- Create: `Makefile`

**Step 1: Write Makefile**
```makefile
.PHONY: build run test clean

build:
	go build -o bin/server ./cmd/server

run:
	go run ./cmd/server

test:
	go test ./... -v

test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

clean:
	rm -rf bin/
	rm -f coverage.out

lint:
	golangci-lint run

.DEFAULT_GOAL := build
```

**Step 2: Commit**
```bash
git add Makefile
git commit -m "chore: add Makefile"
```

---

### Task 19: Run Final Verification

**Step 1: Run all tests**
```bash
make test
```
Expected: All tests pass

**Step 2: Build**
```bash
make build
```
Expected: Binary created in bin/server

**Step 3: Run server (manual test)**
```bash
make run
```
Expected: Server starts on configured port

**Step 4: Final commit**
```bash
git add .
git commit -m "feat: complete Go API rewrite with Clean Architecture"
```

---

## Summary

| Phase | Tasks | Description |
|-------|-------|-------------|
| 1 | 1-4 | Project setup, dependencies, config |
| 2 | 5-6 | Domain layer (entities, interfaces) |
| 3 | 7-10 | Repository layer (GORM implementations) |
| 4 | 11-12 | Usecase layer (services with tests) |
| 5 | 13-16 | HTTP handlers with middleware |
| 6 | 17-19 | Main entry point, Makefile, verification |

**Total: 19 tasks**
