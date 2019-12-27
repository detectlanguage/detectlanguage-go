package detectlanguage

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientBaseURL(t *testing.T) {
	url, _ := url.Parse("http://localhost")
	client := Client{BaseURL: url}
	assert.Equal(t, url, client.baseURL())
}

func TestClientUserAgent(t *testing.T) {
	userAgent := "Test agent"
	client := Client{UserAgent: userAgent}
	assert.Equal(t, userAgent, client.userAgent())
}
