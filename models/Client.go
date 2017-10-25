package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	*http.Client
}

func (c *Client) LoadResponse(path string, i interface{}) error {
	full_path := path

	fmt.Println("querying..." + full_path)
	rsp, e := c.Get(full_path)
	if e != nil {
		return e
	}

	defer rsp.Body.Close()

	b, e := ioutil.ReadAll(rsp.Body)
	if e != nil {
		return e
	}
	if rsp.Status[0] != '2' {
		return fmt.Errorf("expected status 2xx, got %s: %s", rsp.Status, string(b))
	}

	return json.Unmarshal(b, &i)
}

func New() (*Client, error) {
	return &Client{Client: &http.Client{}}, nil
}
