package api

import (
	"github.com/go-resty/resty"
)

type (
	Client struct {
		restClient *resty.Client
		config     *Config
	}
)

func NewClient() *Client {

	restClient := resty.New()
	restClient.SetDebug(true)
	restClient.SetHostURL("http://lisk.fr0m.space:4000/")

	return &Client{
		restClient: restClient,
	}
}
