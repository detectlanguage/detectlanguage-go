package detectlanguage_test

import (
	"testing"

	"github.com/detectlanguage/detectlanguage-go"
	"github.com/stretchr/testify/assert"
)

func TestAPIErrorWithoutMessage(t *testing.T) {
	err := detectlanguage.APIError{StatusCode: 404, Status: "Not found"}
	assert.Equal(t, "404 Not found", err.Error())
}

func TestAPIErrorWithMessage(t *testing.T) {
	err := detectlanguage.APIError{
		StatusCode: 401,
		Status:     "Unauthorized",
		Code:       1,
		Message:    "Invalid API key",
	}

	assert.Equal(t, "Invalid API key", err.Error())
}

func TestDetectionError(t *testing.T) {
	err := detectlanguage.DetectionError{"Test error"}
	assert.Equal(t, "Test error", err.Error())
}
