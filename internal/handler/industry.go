package handler

import (
	"encoding/json"
	"net/http"

	"businessense/internal/usecase/industry"

	"github.com/go-chi/chi/v5"
)

// IndustryHandler handles industry endpoints
type IndustryHandler struct {
	svc *industry.Service
}

// NewIndustryHandler creates a new industry handler
func NewIndustryHandler(svc *industry.Service) *IndustryHandler {
	return &IndustryHandler{svc: svc}
}

// CreateIndustryRequest represents the create industry request body
type CreateIndustryRequest struct {
	Name string `json:"name"`
}

// Create handles industry creation
func (h *IndustryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateIndustryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondFailure(w, http.StatusBadRequest, "Invalid request")
		return
	}

	if req.Name == "" {
		respondFailure(w, http.StatusBadRequest, "Name is required")
		return
	}

	ind, err := h.svc.Create(req.Name)
	if err != nil {
		respondFailure(w, http.StatusInternalServerError, "Failed to create industry")
		return
	}

	respondCreated(w, "Industry created", map[string]any{
		"industry": ind,
	})
}

// Search handles listing all industries or searching by query
func (h *IndustryHandler) Search(w http.ResponseWriter, r *http.Request) {
	search := chi.URLParam(r, "search")

	var industries []interface{}
	var err error

	if search == "" {
		result, e := h.svc.List()
		err = e
		for _, ind := range result {
			industries = append(industries, ind)
		}
	} else {
		result, e := h.svc.Search(search)
		err = e
		for _, ind := range result {
			industries = append(industries, ind)
		}
	}

	if err != nil {
		respondFailure(w, http.StatusInternalServerError, "Failed to fetch industries")
		return
	}

	if industries == nil {
		industries = []interface{}{}
	}

	respondSuccess(w, "Industries retrieved", map[string]any{
		"industries": industries,
	})
}
