package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"businessense/internal/usecase/auth"
)

// AuthHandler handles authentication endpoints
type AuthHandler struct {
	authService *auth.Service
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authService *auth.Service) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// RegisterRequest represents the registration request body
type RegisterRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// LoginRequest represents the login request body
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register handles user registration
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondFailure(w, http.StatusBadRequest, "Invalid request")
		return
	}

	user, token, err := h.authService.Register(req.Email, req.Password, req.FirstName, req.LastName)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrUserExists):
			respondFailure(w, http.StatusConflict, "Email address already in use by another user.")
		default:
			respondFailure(w, http.StatusInternalServerError, "Failed to create account, connection error.")
		}
		return
	}

	respondCreated(w, "Account has been created", map[string]any{
		"account": map[string]any{
			"ID":        user.ID,
			"email":     user.Email,
			"firstname": user.FirstName,
			"lastname":  user.LastName,
			"token":     token,
		},
	})
}

// Login handles user authentication
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondFailure(w, http.StatusBadRequest, "Invalid request")
		return
	}

	user, token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		respondFailure(w, http.StatusUnauthorized, "Invalid login credentials. Please try again")
		return
	}

	respondSuccess(w, "Logged In", map[string]any{
		"account": map[string]any{
			"ID":        user.ID,
			"email":     user.Email,
			"firstname": user.FirstName,
			"lastname":  user.LastName,
			"token":     token,
		},
	})
}
