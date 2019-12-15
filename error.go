package detectlanguage

import "fmt"

// APIError is an error returned from the API
type APIError struct {
	// Status is the HTTP text status of the response.
	Status string
	// StatusCode is the HTTP numerical status code of the response.
	StatusCode int
	// Code is an error code returned by the API.
	Code int `json:"code"`
	// Message is a human-readable error code returned by the API.
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprintf("%d %s", e.StatusCode, e.Status)
}

type apiErrorResponse struct {
	Error *APIError `json:"error,omitempty"`
}

// DetectionError is retuned when detection fails
type DetectionError struct {
	Message string
}

func (e *DetectionError) Error() string {
	return e.Message
}
