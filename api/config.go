package api

import (
	"math/rand"
	"net/url"
	"strconv"
	"time"
)

type (
	// Config is the config for the Lisk API client
	Config struct {
		// Host to use for API calls - can be nil when RandomHost is true.
		Host Host
		// RandomHost specifies whether a random Host from RandomHostsPool should be used.
		RandomHost bool
		// RandomHostsPool is a list of hosts from which a random one is selected on Client creation.
		RandomHostsPool []Host
		// Debug specifies whether debug logging for the API client should be activated.
		Debug bool
	}
	// Host is a Lisk Node
	Host struct {
		// Hostname is the hostname/IP of the Lisk Node to connect to.
		Hostname string
		// Port is the port used by the Lisk node.
		Port int
		// Secure specifies whether https should be used.
		Secure bool
	}
)

var (
	// DefaultConfig is the default config for the Lisk API client
	DefaultConfig = &Config{
		RandomHostsPool: []Host{{"betanet.lisk.io", 5000, false}},
		RandomHost:      true,
		Debug:           false,
	}
)

// GetHostURL composes a URL from the host details
func (h Host) GetHostURL() string {
	var hostURL url.URL

	hostURL.Host = h.Hostname + ":" + strconv.Itoa(h.Port)

	if h.Secure {
		hostURL.Scheme = "https"
	} else {
		hostURL.Scheme = "http"
	}

	return hostURL.String()
}

// GetRandomHost returns a random host from RandomHostsPool
func (c *Config) GetRandomHost() Host {
	rand.Seed(time.Now().Unix())
	return c.RandomHostsPool[rand.Intn(len(c.RandomHostsPool))]
}
