package handler

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents a standard error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// respondJSON writes a JSON response with the given status code
func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// respondError writes a JSON error response
func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, ErrorResponse{
		Error:   http.StatusText(status),
		Message: message,
	})
}

// respondSuccess writes a success response matching the old API format
func respondSuccess(w http.ResponseWriter, message string, data map[string]any) {
	response := map[string]any{
		"status":  true,
		"message": message,
	}
	for k, v := range data {
		response[k] = v
	}
	respondJSON(w, http.StatusOK, response)
}

// respondCreated writes a created response matching the old API format
func respondCreated(w http.ResponseWriter, message string, data map[string]any) {
	response := map[string]any{
		"status":  true,
		"message": message,
	}
	for k, v := range data {
		response[k] = v
	}
	respondJSON(w, http.StatusCreated, response)
}

// respondFailure writes a failure response matching the old API format
func respondFailure(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]any{
		"status":  false,
		"message": message,
	})
}
