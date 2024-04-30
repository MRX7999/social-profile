package quickosint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client представляет клиент для работы с API QuickOSINT.
type Client struct {
	BaseURL    string
	Token      string
	ClientID   string
	HTTPClient *http.Client
}

// NewClient создает новый клиент для работы с API QuickOSINT.
func NewClient(baseURL, clientID string) (*Client, error) {
	token, err := readAPIKey("apikey.txt")
	if err != nil {
		return nil, err
	}

	return &Client{
		BaseURL:    baseURL,
		Token:      token,
		ClientID:   clientID,
		HTTPClient: &http.Client{},
	}, nil
}

// readAPIKey читает API ключ из файла apikey.txt.
func readAPIKey(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes.TrimSpace(data)), nil
}

// SearchDetail выполняет поиск детального отчета для указанного запроса.
func (c *Client) SearchDetail(query string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/v1/search/detail/%s", c.BaseURL, query)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Set("X-ClientId", c.ClientID)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// SearchDetailBatch выполняет пакетный поиск детального отчета для списка запросов.
func (c *Client) SearchDetailBatch(queries []string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/v1/search/detail", c.BaseURL)

	payload := map[string]interface{}{
		"Items": queries,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Set("X-ClientId", c.ClientID)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
