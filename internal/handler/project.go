package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"businessense/internal/usecase/project"

	"github.com/go-chi/chi/v5"
)

// ProjectHandler handles project endpoints
type ProjectHandler struct {
	svc *project.Service
}

// NewProjectHandler creates a new project handler
func NewProjectHandler(svc *project.Service) *ProjectHandler {
	return &ProjectHandler{svc: svc}
}

// CreateProjectRequest represents the create project request body
type CreateProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CompanyID   uint   `json:"companyId"`
}

// Create handles project creation
func (h *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondFailure(w, http.StatusBadRequest, "Invalid request")
		return
	}

	if req.Name == "" {
		respondFailure(w, http.StatusBadRequest, "Name is required")
		return
	}

	if req.CompanyID == 0 {
		respondFailure(w, http.StatusBadRequest, "Company ID is required")
		return
	}

	proj, err := h.svc.Create(req.Name, req.Description, req.CompanyID)
	if err != nil {
		respondFailure(w, http.StatusInternalServerError, "Failed to create project")
		return
	}

	respondCreated(w, "Project created", map[string]any{
		"project": proj,
	})
}

// Get handles getting a single project by ID
func (h *ProjectHandler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondFailure(w, http.StatusBadRequest, "Invalid project ID")
		return
	}

	proj, err := h.svc.GetByID(uint(id))
	if err != nil {
		respondFailure(w, http.StatusNotFound, "Project not found")
		return
	}

	respondSuccess(w, "Project retrieved", map[string]any{
		"project": proj,
	})
}

// GetByCompany handles getting projects for a specific company
func (h *ProjectHandler) GetByCompany(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondFailure(w, http.StatusBadRequest, "Invalid company ID")
		return
	}

	projects, err := h.svc.ListByCompany(uint(id))
	if err != nil {
		respondFailure(w, http.StatusInternalServerError, "Failed to fetch projects")
		return
	}

	respondSuccess(w, "Projects retrieved", map[string]any{
		"projects": projects,
	})
}
