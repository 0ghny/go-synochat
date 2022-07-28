package synochat

import (
	"bytes"
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type ChatMessage struct {
	Text string `json:"text"`
	// FileUrl string `json:"file_url,omitempty"`
}

func (c *Client) SendMessage(msg *ChatMessage, token string) error {
	rUrl := fmt.Sprintf("%s/webapi/entry.cgi?api=SYNO.Chat.External&method=incoming&version=2&token=%s", c.BaseURL, token)

	json_data, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
		return err
	}

	payload := fmt.Sprintf("payload=%s", string(json_data))
	payloadBytes := []byte(payload)
	rsp, err := c.httpClient.Post(rUrl, "application/x-www-form-urlencoded", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Debugf("response status code: %d", rsp.StatusCode)

	apiResponse, err := NewApiResponseFromHttpResponse(rsp)
	if err != nil {
		log.Fatal(err)
	}

	if !apiResponse.Success {
		eer := fmt.Errorf("synochat api response error, name: %s - reason: %s", apiResponse.Error.Errors.Name, apiResponse.Error.Errors.Reason)
		log.Fatal(eer)
		return eer
	}
	return nil
}
