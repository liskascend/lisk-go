package api

type (
	Config struct {
		Host       Host
		Hosts      []Host
		RandomHost bool
	}
	Host struct {
		Hostname string
		Port     int
		Secure   bool
	}
)

var (
	DefaultConfig = &Config{
		Hosts:      []Host{{"lisk.fr0m.space", 4000, false}},
		RandomHost: true,
	}
)
