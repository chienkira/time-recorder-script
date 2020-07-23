package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/tidwall/gjson"
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

func (c *Client) postForm(path string, body io.Reader) (gjson.Result, error) {
	req, err := c.newRequest("POST", path, body)
	if err != nil {
		c.logger.Panicln(err)
		return gjson.Result{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := c.HTTPClient.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		c.logger.Panicln(res)
		return gjson.Result{}, err
	}
	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.logger.Panicln(err)
		return gjson.Result{}, err
	}
	bodyString := string(bodyBytes)
	json_res := gjson.Parse(bodyString)
	return json_res, nil
}
