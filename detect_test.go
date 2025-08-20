package detectlanguage_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetect(t *testing.T) {
	detections, err := client.Detect(context.TODO(), "labas rytas")

	if assert.NoError(t, err) {
		assert.Equal(t, "lt", detections[0].Language)
		assert.Greater(t, detections[0].Score, float64(0))
	}
}

func TestDetectCode(t *testing.T) {
	code, err := client.DetectCode(context.TODO(), "labas rytas")

	if assert.NoError(t, err) {
		assert.Equal(t, "lt", code)
	}
}

func TestDetectCodeFailure(t *testing.T) {
	code, err := client.DetectCode(context.TODO(), " ")

	assert.EqualError(t, err, "Language not detected")
	assert.Equal(t, code, "")
}

func TestDetectBatch(t *testing.T) {
	query := []string{"labas rytas", "good morning"}
	detections, err := client.DetectBatch(context.TODO(), query)

	if assert.NoError(t, err) {
		assert.Equal(t, "lt", detections[0][0].Language)
		assert.Equal(t, "en", detections[1][0].Language)
	}
}
