package client

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	Client   *http.Client
	BasePath string
}

func NewHTTPClient(basePath string) Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{
		Timeout:   90 * time.Second,
		Transport: tr,
	}

	client := &Client{
		Client:   c,
		BasePath: basePath,
	}
	return *client
}

func (c *Client) Do(request *http.Request) (*http.Response, error) {
	return c.Client.Do(request)
}

// returns response body if status code is 200
func (c *Client) DoRequest(req *http.Request) ([]byte, error) {
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status code %d", resp.StatusCode)
	}

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyText, nil
}
