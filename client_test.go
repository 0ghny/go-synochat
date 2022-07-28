package synochat_test

import (
	"testing"

	"github.com/0ghny/go-synochat"
	"github.com/stretchr/testify/assert"
)

func TestNewClient_WithValidUrl_ShouldCreateNewClient(t *testing.T) {
	c, err := synochat.NewClient("http://syno.local")

	assert.Nil(t, err)
	assert.NotNil(t, c)
}

func TestNewClient_WithInvalidBaseURL_ShouldReturnsAnError(t *testing.T) {
	c, err := synochat.NewClient("syno.local")

	assert.Nil(t, c)
	assert.NotNil(t, err)
}

func TestNewClient_WithEmptyBaseURL_ShouldReturnsAnError(t *testing.T) {
	c, err := synochat.NewClient("")

	assert.Nil(t, c)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "url cannot be empty or just whitespaces")
}
