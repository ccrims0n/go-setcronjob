package setcronjob

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const apiBaseURL = "https://www.setcronjob.com/api/"

type responseBody struct {
	io.Reader
}

func (r *responseBody) Close() error {
	return nil
}

type Auth struct {
	Token string
}

func (auth *Auth) validate() error {
	e := nil

	if auth.Token == "" {
		e = errors.New("Token is required")
	}

	return e
}

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	c           httpClient
	token       string
	cronClient  Crons
	groupClient Groups
}

func NewClient(auth Auth) (*Client, error) {
	if err := auth.validate(); err != nil {
		return nil, err
	}

	return &Client{
		c: &http.Client{},
		token: auth.Token,
	}, nil
}

func (c *Client) newRequest(method, target, action string, v url.Values, body io.Reader}) (*http.Request, error) {
	v.Set("token", c.token)

	url := fmt.Sprintf("%s/%s.%s?%s", apiBaseURL, target, action, v.Encode())

	r, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) doRequest(r *http.Request) (*http.Response, error) {
	resp, err := c.c.Do(r)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New("response error")
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var response response

	err = json.Unmarshal(b, &response)

	if err == nil && response.status == "error" {
		return nil, errors.New(response.message)
	}

	resp.Body = &responseBody{
		Reader: bytes.NewReader(b),
	}

	return resp, nil
}

func (c *Client) Crons() Crons {
	if c.cronClient == nil {
		c.cronClient = newCrons(*c)
	}

	return c.cronClient
}

func (c *Client) Groups() Groups {
	if c.groupClient == nil {
		c.groupClient = newGroups(*c)
	}

	return c.groupClient
}