package synochat

import (
	"encoding/json"
	"io"
	"net/http"
)

// "{\"error\":{\"code\":120,\"errors\":{\"name\":\"payload\",\"reason\":\"required\"}},\"success\":false}"
type ApiResponse struct {
	Success bool `json:"success"`
	Error   struct {
		Code   int `json:"code"`
		Errors struct {
			Name   string `json:"name"`
			Reason string `json:"reason"`
		} `json:"errors,omitempty"`
	}
}

// Creates an ApiResponse from an http.Response object
func NewApiResponseFromHttpResponse(resp *http.Response) (*ApiResponse, error) {
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response ApiResponse
	if err = json.Unmarshal(b, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
