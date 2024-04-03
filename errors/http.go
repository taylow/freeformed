package errors

import (
	"encoding/json"
	"net/http"
)

// HandleWithError wraps a HandleWithErrorFunc and returns a http.HandlerFunc that can be used to handle http requests with errors
func HandleWithError(handler HandleWithErrorFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			w.Header().Set("Content-Type", "application/json")
			if httpErr, ok := err.(*Error); ok {
				w.WriteHeader(httpErr.Code)
				switch r.Header.Get("Accept") {
				case "application/json":
					json.NewEncoder(w).Encode(httpErr.Response())
				case "text/plain":
					w.Write([]byte(httpErr.Response().Message))
				default:
					json.NewEncoder(w).Encode(httpErr.Response())
				}
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Message: "internal server error", Code: http.StatusInternalServerError})
		}
	}
}

// HandleWithErrorFunc is a function that handles a http request and returns an error
type HandleWithErrorFunc func(w http.ResponseWriter, r *http.Request) error

// ErrorResponse represents an error response
type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// Handle404 is a http.HandlerFunc that returns a 404 response
func Handle404() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "not found", Code: http.StatusNotFound})
	}
}
