package handler

import (
	"encoding/json"
	"net/http"

	"businessense/internal/usecase/painpoint"

	"github.com/go-chi/chi/v5"
)

// PainPointHandler handles painpoint endpoints
type PainPointHandler struct {
	svc *painpoint.Service
}

// NewPainPointHandler creates a new painpoint handler
func NewPainPointHandler(svc *painpoint.Service) *PainPointHandler {
	return &PainPointHandler{svc: svc}
}

// CreatePainPointRequest represents the create painpoint request body
type CreatePainPointRequest struct {
	Name string `json:"name"`
}

// Create handles painpoint creation
func (h *PainPointHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreatePainPointRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondFailure(w, http.StatusBadRequest, "Invalid request")
		return
	}

	if req.Name == "" {
		respondFailure(w, http.StatusBadRequest, "Name is required")
		return
	}

	pp, err := h.svc.Create(req.Name)
	if err != nil {
		respondFailure(w, http.StatusInternalServerError, "Failed to create painpoint")
		return
	}

	respondCreated(w, "Painpoint created", map[string]any{
		"painpoint": pp,
	})
}

// Search handles listing all painpoints or searching by query
func (h *PainPointHandler) Search(w http.ResponseWriter, r *http.Request) {
	search := chi.URLParam(r, "search")

	var painpoints []interface{}
	var err error

	if search == "" {
		result, e := h.svc.List()
		err = e
		for _, pp := range result {
			painpoints = append(painpoints, pp)
		}
	} else {
		result, e := h.svc.Search(search)
		err = e
		for _, pp := range result {
			painpoints = append(painpoints, pp)
		}
	}

	if err != nil {
		respondFailure(w, http.StatusInternalServerError, "Failed to fetch painpoints")
		return
	}

	if painpoints == nil {
		painpoints = []interface{}{}
	}

	respondSuccess(w, "Painpoints retrieved", map[string]any{
		"painpoints": painpoints,
	})
}
