package synochat

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

type ClientOpts struct {
	BaseUrl    string
	HTTPClient *http.Client
}

type Client struct {
	BaseURL    *url.URL
	httpClient *http.Client
}

func NewClient(baseURL string) (*Client, error) {
	if len(strings.TrimSpace(baseURL)) <= 0 {
		return nil, errors.New("url cannot be empty or just whitespaces")
	}

	u, err := url.ParseRequestURI(baseURL)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return nil, err
	}

	return &Client{
		httpClient: &http.Client{},
		BaseURL:    u,
	}, nil
}

func NewCustomClient(baseURL string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	if len(strings.TrimSpace(baseURL)) <= 0 {
		return nil, errors.New("url cannot be empty or just whitespaces")
	}

	u, err := url.ParseRequestURI(baseURL)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return nil, err
	}

	return &Client{
		httpClient: httpClient,
		BaseURL:    u,
	}, nil
}
