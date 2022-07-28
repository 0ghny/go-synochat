package synochat_test

import (
	"testing"

	"github.com/0ghny/go-synochat"
	"github.com/stretchr/testify/require"
)

const (
	sutBaseURL = "https://url"
	sutToken   = "atoken"
)

func IgnoreTestSendMessageToChannel(t *testing.T) {
	c, err := synochat.NewClient(sutBaseURL)

	require.Nil(t, err)
	require.NotNil(t, c)

	err = c.SendMessage(&synochat.ChatMessage{Text: "Hello from automated test"}, sutToken)

	require.Nil(t, err)
}
