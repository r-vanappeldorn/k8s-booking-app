package errors

import (
	"encoding/json"
	"net/http"
)

type JSONErrorResponse struct {
	StatusCode int              `json:"status_code"`
	Detail     *JSONErrorDetail `json:"detail"`
}

type JSONErrorDetail struct {
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Fields  []*FieldError `json:"fields"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewJSONErrorResponse(statusCode int, message string, code string) *JSONErrorResponse {
	return &JSONErrorResponse{
		StatusCode: statusCode,
		Detail: &JSONErrorDetail{
			Code:    code,
			Message: message,
		},
	}
}

func (e *JSONErrorResponse) AddFieldErrors(fieldErrors []*FieldError) {
	e.Detail.Fields = append(e.Detail.Fields, fieldErrors...)
}

func WriteErrorResponse(w http.ResponseWriter, e *JSONErrorResponse) {
	w.WriteHeader(e.StatusCode)
	json.NewEncoder(w).Encode(e)
}
