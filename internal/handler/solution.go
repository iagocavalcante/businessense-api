package handler

import (
	"net/http"
	"strconv"

	"businessense/internal/usecase/solution"

	"github.com/go-chi/chi/v5"
)

// SolutionHandler handles solution endpoints
type SolutionHandler struct {
	svc *solution.Service
}

// NewSolutionHandler creates a new solution handler
func NewSolutionHandler(svc *solution.Service) *SolutionHandler {
	return &SolutionHandler{svc: svc}
}

// Get handles getting a single solution by ID
func (h *SolutionHandler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondFailure(w, http.StatusBadRequest, "Invalid solution ID")
		return
	}

	sol, err := h.svc.GetByID(uint(id))
	if err != nil {
		respondFailure(w, http.StatusNotFound, "Solution not found")
		return
	}

	respondSuccess(w, "Solution retrieved", map[string]any{
		"solution": sol,
	})
}

// GetByIssue handles getting solutions for a specific issue with constraints
func (h *SolutionHandler) GetByIssue(w http.ResponseWriter, r *http.Request) {
	issueIDStr := chi.URLParam(r, "issueid")
	issueID, err := strconv.Atoi(issueIDStr)
	if err != nil {
		respondFailure(w, http.StatusBadRequest, "Invalid issue ID")
		return
	}

	solutions, err := h.svc.GetByIssue(uint(issueID))
	if err != nil {
		respondFailure(w, http.StatusInternalServerError, "Failed to fetch solutions")
		return
	}

	// Build response with constraints for each solution
	var solutionData []map[string]any
	for _, sol := range solutions {
		constraints, err := h.svc.GetConstraints(uint(issueID), sol.ID)
		if err != nil {
			constraints = nil
		}

		solutionData = append(solutionData, map[string]any{
			"id":          sol.ID,
			"name":        sol.Name,
			"description": sol.Description,
			"createdAt":   sol.CreatedAt,
			"updatedAt":   sol.UpdatedAt,
			"constraints": constraints,
		})
	}

	if solutionData == nil {
		solutionData = []map[string]any{}
	}

	respondSuccess(w, "Solutions retrieved", map[string]any{
		"solutions": solutionData,
	})
}
