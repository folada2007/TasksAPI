package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeJSONBody(r *http.Request, obj interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(&obj)
}
