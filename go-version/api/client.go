package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
)

type Client struct {
	baseUrl    string
	HTTPClient *http.Client
	logger     *log.Logger
}

func NewClient(baseUrl string, logger *log.Logger) *Client {
	if logger == nil {
		logger = log.New(os.Stdout, "[API]", log.LstdFlags)
	}

	return &Client{
		baseUrl:    baseUrl,
		HTTPClient: &http.Client{},
		logger:     logger,
	}
}

func decodeBody(res *http.Response, out interface{}) error {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	return decoder.Decode(out)
}

func (c *Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	url := c.baseUrl + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		c.logger.Panicln(err)
		return nil, err
	}
	userAgent := fmt.Sprintf("API Client Go(%s)", runtime.Version())
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", userAgent)

	return req, nil
}

func (c *Client) post(path string, body io.Reader) (map[string]interface{}, error) {
	req, err := c.newRequest("POST", path, body)
	if err != nil {
		c.logger.Panicln(err)
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		c.logger.Panicln(res)
		return nil, err
	}
	var json_res map[string]interface{}
	decodeBody(res, &json_res)
	c.logger.Println(json_res)
	return json_res, nil
}
