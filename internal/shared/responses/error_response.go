package responses

import (
	"errors"
	"igaming-platform/internal/shared/fallacies"
	"net/http"
)

type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
}

func Error(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	detail := err.Error()

	switch {
	case errors.Is(err, fallacies.TournamentNotFound):
		status = http.StatusNotFound

	case errors.Is(err, fallacies.PrizesAlreadyDistributed):
		status = http.StatusConflict

	}

	WriteJSON(w, status, ErrorResponse{
		Status: status,
		Error:  detail,
	})
}
