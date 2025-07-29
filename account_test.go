package detectlanguage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountStatus(t *testing.T) {
	response, err := client.AccountStatus()

	if assert.NoError(t, err) {
		assert.NotEmpty(t, response.Date)
		assert.Equal(t, "ACTIVE", response.Status)
	}
}
