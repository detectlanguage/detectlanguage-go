package detectlanguage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLanguages(t *testing.T) {
	languages, err := client.Languages()

	if assert.NoError(t, err) {
		assert.NotEmpty(t, languages[0].Code)
		assert.NotEmpty(t, languages[0].Name)
	}
}
