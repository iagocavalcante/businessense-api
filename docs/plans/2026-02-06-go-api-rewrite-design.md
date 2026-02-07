# Go API Rewrite Design

> **Date:** 2026-02-06
> **Status:** Approved

## Overview

Complete rewrite of the businessense Go API using modern patterns and Clean Architecture.

## Current Issues

| Aspect | Current | Problem |
|--------|---------|---------|
| Go Version | Pre-modules | No go.mod, outdated dependency management |
| ORM | `jinzhu/gorm` | Deprecated, should use `gorm.io/gorm` v2 |
| JWT | `dgrijalva/jwt-go` | Deprecated, security vulnerabilities |
| Structure | models/controllers/utils | Mixed concerns, no services layer |
| Error Handling | `map[string]interface{}` | No typed errors, inconsistent |
| Testing | None | No tests exist |
| Config | `godotenv` in init() | Global state, hard to test |

## Technology Stack

- **Go:** 1.22+ with modules
- **Router:** Chi (lightweight, stdlib compatible)
- **ORM:** GORM v2 (`gorm.io/gorm`)
- **JWT:** `github.com/golang-jwt/jwt/v5`
- **Testing:** testify + mockery
- **Architecture:** Clean Architecture

## Project Structure

```
businessense-api/
├── cmd/
│   └── server/
│       └── main.go              # Entry point, DI setup
├── internal/
│   ├── domain/                  # Business entities (no deps)
│   │   ├── user.go
│   │   ├── company.go
│   │   ├── industry.go
│   │   ├── painpoint.go
│   │   ├── issue.go
│   │   ├── solution.go
│   │   ├── constraint.go
│   │   └── project.go
│   ├── usecase/                 # Business logic (services)
│   │   ├── auth/
│   │   │   ├── service.go
│   │   │   └── service_test.go
│   │   ├── company/
│   │   ├── industry/
│   │   ├── painpoint/
│   │   ├── issue/
│   │   ├── solution/
│   │   └── project/
│   ├── repository/              # Data access
│   │   ├── interfaces.go        # Repository interfaces
│   │   └── postgres/
│   │       ├── db.go
│   │       ├── user.go
│   │       ├── company.go
│   │       └── ...
│   └── handler/                 # HTTP layer
│       ├── auth.go
│       ├── company.go
│       ├── industry.go
│       ├── painpoint.go
│       ├── issue.go
│       ├── solution.go
│       ├── project.go
│       ├── response.go
│       └── middleware/
│           ├── auth.go
│           └── logging.go
├── pkg/
│   ├── config/
│   │   └── config.go
│   └── validator/
│       └── validator.go
├── migrations/
│   └── ...
├── go.mod
├── go.sum
├── Makefile
└── .env.example
```

## Domain Layer

Pure Go structs with no external dependencies. Repository interfaces defined alongside entities.

```go
// internal/domain/user.go
package domain

import "time"

type User struct {
    ID        uint      `json:"id"`
    Email     string    `json:"email"`
    FirstName string    `json:"firstName"`
    LastName  string    `json:"lastName"`
    Password  string    `json:"-"` // Never serialized
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

type UserRepository interface {
    Create(user *User) error
    FindByID(id uint) (*User, error)
    FindByEmail(email string) (*User, error)
}
```

## Usecase Layer (Services)

Business logic with dependency injection. Uses typed errors.

```go
// internal/usecase/auth/service.go
package auth

import (
    "errors"
    "businessense/internal/domain"
    "businessense/pkg/config"
)

var (
    ErrInvalidCredentials = errors.New("invalid credentials")
    ErrUserExists         = errors.New("user already exists")
    ErrUserNotFound       = errors.New("user not found")
)

type Service struct {
    userRepo  domain.UserRepository
    jwtSecret string
}

func NewService(userRepo domain.UserRepository, cfg *config.Config) *Service {
    return &Service{
        userRepo:  userRepo,
        jwtSecret: cfg.JWTSecret,
    }
}

func (s *Service) Register(email, password, firstName, lastName string) (*domain.User, string, error) {
    // Implementation
}

func (s *Service) Login(email, password string) (*domain.User, string, error) {
    // Implementation
}
```

## HTTP Handler Layer

Thin handlers that delegate to services. Request DTOs with validation.

```go
// internal/handler/auth.go
package handler

type AuthHandler struct {
    authService *auth.Service
}

type RegisterRequest struct {
    Email     string `json:"email" validate:"required,email"`
    Password  string `json:"password" validate:"required,min=6"`
    FirstName string `json:"firstName" validate:"required"`
    LastName  string `json:"lastName" validate:"required"`
}

type AuthResponse struct {
    User  *domain.User `json:"user"`
    Token string       `json:"token"`
}
```

## Response Format

Consistent JSON responses:

```go
// Success
{
    "user": {...},
    "token": "..."
}

// Error
{
    "error": "Bad Request",
    "message": "invalid email format"
}
```

## Router Setup

Chi router with middleware chain and route groups:

```go
r := chi.NewRouter()

// Global middleware
r.Use(chiMiddleware.Logger)
r.Use(chiMiddleware.Recoverer)
r.Use(cors.Handler(...))

// Public routes
r.Post("/api/user/new", authHandler.Register)
r.Post("/api/user/login", authHandler.Login)

// Protected routes
r.Group(func(r chi.Router) {
    r.Use(middleware.JWTAuth(cfg.JWTSecret))
    // Protected endpoints...
})
```

## API Endpoints (Same as Current)

| Method | Path | Handler | Auth |
|--------|------|---------|------|
| POST | /api/user/new | Register | No |
| POST | /api/user/login | Login | No |
| POST | /api/industry/new | CreateIndustry | Yes |
| GET | /api/industry/search | ListIndustries | Yes |
| POST | /api/company/new | CreateCompany | Yes |
| GET | /api/company/search | ListCompanies | Yes |
| POST | /api/painpoint/new | CreatePainPoint | Yes |
| GET | /api/painpoint/search | ListPainPoints | Yes |
| GET | /api/issue/{id} | GetIssue | Yes |
| POST | /api/issue/relevance | GetIssueRelevance | Yes |
| GET | /api/solution/{id} | GetSolution | Yes |
| GET | /api/solution/issue/{id} | GetSolutionsForIssue | Yes |
| POST | /api/project/new | CreateProject | Yes |
| GET | /api/project/{id} | GetProject | Yes |
| GET | /api/project/company/{id} | GetProjectsByCompany | Yes |

## Testing Strategy

- **Unit tests:** Services with mocked repositories
- **Integration tests:** Handlers with test database
- **Table-driven tests:** Standard Go pattern
- **Mocks:** Generated with mockery

## Dependencies

```
github.com/go-chi/chi/v5
github.com/go-chi/cors
github.com/golang-jwt/jwt/v5
gorm.io/gorm
gorm.io/driver/postgres
golang.org/x/crypto
github.com/stretchr/testify
github.com/vektra/mockery/v2
```
