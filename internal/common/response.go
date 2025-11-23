package common

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func WriteJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func Success(w http.ResponseWriter, message string, data interface{}) {
	WriteJSON(w, http.StatusOK, SuccessResponse{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	})
}

func Error(w http.ResponseWriter, status int, message string, err string) {
	WriteJSON(w, status, ErrorResponse{
		Status:  status,
		Message: message,
		Error:   err,
	})
}
