package detectlanguage_test

import (
	"os"
	"testing"

	"github.com/detectlanguage/detectlanguage-go"
	"github.com/stretchr/testify/assert"
)

var client = detectlanguage.New(os.Getenv("DETECTLANGUAGE_API_KEY"))

func TestDetect(t *testing.T) {
	detections, err := client.Detect("labas rytas")

	if assert.NoError(t, err) {
		assert.Equal(t, detections[0].Language, "lt")
		assert.True(t, detections[0].Reliable)
		assert.Greater(t, detections[0].Confidence, float32(0))
	}
}

func TestDetectCode(t *testing.T) {
	code, err := client.DetectCode("labas rytas")

	if assert.NoError(t, err) {
		assert.Equal(t, code, "lt")
	}
}

func TestDetectCodeFailure(t *testing.T) {
	code, err := client.DetectCode("")

	assert.EqualError(t, err, "Language not detected")
	assert.Equal(t, code, "")
}

func TestDetectBatch(t *testing.T) {
	query := []string{"labas rytas", "good morning"}
	detections, err := client.DetectBatch(query)

	if assert.NoError(t, err) {
		assert.Equal(t, detections[0][0].Language, "lt")
		assert.Equal(t, detections[1][0].Language, "en")
	}
}
