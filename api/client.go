package api

import (
	"github.com/go-resty/resty"
)

type (
	// Client represents a client for the Lisk API
	Client struct {
		restClient *resty.Client
		config     *Config
	}
)

// NewClient returns a new client for the Lisk API
func NewClient() *Client {
	return NewClientWithCustomConfig(DefaultConfig)
}

// NewClientWithCustomConfig returns a new client for the Lisk API and uses a custom config
func NewClientWithCustomConfig(config *Config) *Client {
	restClient := resty.New()
	restClient.SetDebug(config.Debug)

	var hostURL string

	if config.RandomHost {
		hostURL = config.GetRandomHost().GetHostURL()
	} else {
		hostURL = config.Host.GetHostURL()
	}
	restClient.SetHostURL(hostURL)

	return &Client{
		restClient: restClient,
		config:     config,
	}
}

// SetHost sets the Lisk node for the client requests
func (c *Client) SetHost(host Host) {
	c.restClient.SetHostURL(host.GetHostURL())
}

// ChangeRandomHost selects a new host from the pool for the client requests
func (c *Client) ChangeRandomHost() {
	c.restClient.SetHostURL(c.config.GetRandomHost().GetHostURL())
}
