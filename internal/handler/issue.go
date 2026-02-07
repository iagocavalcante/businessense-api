package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"businessense/internal/usecase/issue"

	"github.com/go-chi/chi/v5"
)

// IssueHandler handles issue endpoints
type IssueHandler struct {
	svc *issue.Service
}

// NewIssueHandler creates a new issue handler
func NewIssueHandler(svc *issue.Service) *IssueHandler {
	return &IssueHandler{svc: svc}
}

// Get handles getting a single issue by ID
func (h *IssueHandler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondFailure(w, http.StatusBadRequest, "Invalid issue ID")
		return
	}

	iss, err := h.svc.GetByID(uint(id))
	if err != nil {
		respondFailure(w, http.StatusNotFound, "Issue not found")
		return
	}

	respondSuccess(w, "Issue retrieved", map[string]any{
		"issue": iss,
	})
}

// GetRelevance handles getting issue relevance based on painpoint IDs
func (h *IssueHandler) GetRelevance(w http.ResponseWriter, r *http.Request) {
	var painPointIDs []uint
	if err := json.NewDecoder(r.Body).Decode(&painPointIDs); err != nil {
		respondFailure(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if len(painPointIDs) == 0 {
		respondFailure(w, http.StatusBadRequest, "Painpoint IDs are required")
		return
	}

	relevance, err := h.svc.GetRelevance(painPointIDs)
	if err != nil {
		respondFailure(w, http.StatusInternalServerError, "Failed to calculate relevance")
		return
	}

	respondSuccess(w, "Relevance calculated", map[string]any{
		"issues": relevance,
	})
}
