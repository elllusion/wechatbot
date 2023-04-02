package gpt35

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"
)

const ModelGpt35Turbo = "gpt-3.5-turbo"

const MaxTokensGpt35Turbo = 4096

const (
	RoleUser      RoleType = "user"
	RoleAssistant RoleType = "assistant"
	RoleSystem    RoleType = "system"
)

const DefaultUrl = "https://api.openai.com/v1/chat/completions"

type Client struct {
	transport *http.Client
	apiKey    string
	url       string
}

func NewClient(apiKey, proxy string) *Client {
	proxyUrl, err := url.Parse(proxy)
	if err != nil {
		log.Fatal(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}

	return &Client{
		transport: &http.Client{
			Transport: transport,
			Timeout:   110 * time.Second,
		},
		apiKey: apiKey,
		url:    DefaultUrl,
	}
}

func NewClientCustomUrl(apiKey string, url string) *Client {
	return &Client{
		transport: http.DefaultClient,
		apiKey:    apiKey,
		url:       url,
	}
}

func (c *Client) GetChat(r *Request) (*Response, error) {
	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	httpResp, err := c.transport.Do(req)

	if err != nil {
		return nil, err
	}
	defer func() {
		_ = httpResp.Body.Close()
	}()

	var resp Response
	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
