package handler

import (
	"encoding/json"
	"net/http"

	"businessense/internal/handler/middleware"
	"businessense/internal/usecase/company"
)

// CompanyHandler handles company endpoints
type CompanyHandler struct {
	svc *company.Service
}

// NewCompanyHandler creates a new company handler
func NewCompanyHandler(svc *company.Service) *CompanyHandler {
	return &CompanyHandler{svc: svc}
}

// CreateCompanyRequest represents the create company request body
type CreateCompanyRequest struct {
	Name       string `json:"name"`
	IndustryID uint   `json:"industryId"`
}

// Create handles company creation
func (h *CompanyHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	if userID == 0 {
		respondFailure(w, http.StatusUnauthorized, "User not authenticated")
		return
	}

	var req CreateCompanyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondFailure(w, http.StatusBadRequest, "Invalid request")
		return
	}

	if req.Name == "" {
		respondFailure(w, http.StatusBadRequest, "Name is required")
		return
	}

	comp, err := h.svc.Create(req.Name, req.IndustryID, userID)
	if err != nil {
		respondFailure(w, http.StatusInternalServerError, "Failed to create company")
		return
	}

	respondCreated(w, "Company created", map[string]any{
		"company": comp,
	})
}

// List handles listing companies for the current user
func (h *CompanyHandler) List(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	if userID == 0 {
		respondFailure(w, http.StatusUnauthorized, "User not authenticated")
		return
	}

	companies, err := h.svc.ListByUser(userID)
	if err != nil {
		respondFailure(w, http.StatusInternalServerError, "Failed to fetch companies")
		return
	}

	respondSuccess(w, "Companies retrieved", map[string]any{
		"companies": companies,
	})
}
