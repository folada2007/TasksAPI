package utils

import (
	"LongTaskAPI/pkg/dto"
	"encoding/json"
	"net/http"
)

func RespondWithErrors(w http.ResponseWriter, code int, massage string) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	return json.NewEncoder(w).Encode(dto.ErrorResponse{
		Error: dto.Error{
			Code:    code,
			Message: massage,
		},
	})
}
