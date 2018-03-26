package api

import (
	"context"
	"strconv"
)

type (
	PeerRequest struct {
		IP        string
		HttpPort  *int
		WSPort    *int
		OS        string
		Version   string
		State     *int
		Height    *int64
		Broadhash string

		ListOptions
	}

	PeerResponse struct {
		Peers []*Peer `json:"data"`
		*GenericResponse
	}

	Peer struct {
		IP        string `json:"ip"`
		HTTPPort  int    `json:"httpPort"`
		WsPort    int    `json:"wsPort"`
		Os        string `json:"os"`
		Version   string `json:"version"`
		State     int    `json:"state"`
		Height    int64  `json:"height"`
		Broadhash string `json:"broadhash"`
		Nonce     string `json:"nonce"`
	}
)

// GetPeers searches for peers.
// Search parameters can be specified in options.
// Limit is set to 100 by default
func (c *Client) GetPeers(ctx context.Context, options *PeerRequest) (*PeerResponse, error) {
	req := c.restClient.R().SetContext(ctx)

	if options != nil {
		if options.IP != "" {
			req.SetQueryParam("ip", options.IP)
		}
		if options.HttpPort != nil {
			req.SetQueryParam("httpPort", strconv.Itoa(*options.HttpPort))
		}
		if options.WSPort != nil {
			req.SetQueryParam("wsPort", strconv.Itoa(*options.WSPort))
		}
		if options.OS != "" {
			req.SetQueryParam("os", options.OS)
		}
		if options.Version != "" {
			req.SetQueryParam("version", options.Version)
		}
		if options.State != nil {
			req.SetQueryParam("state", strconv.Itoa(*options.State))
		}
		if options.Height != nil {
			req.SetQueryParam("height", strconv.FormatInt(*options.Height, 10))
		}
		if options.Broadhash != "" {
			req.SetQueryParam("broadhash", options.Broadhash)
		}

		if options.Limit <= 0 {
			options.Limit = 100
		}

		if options.Offset < 0 {
			options.Offset = 0
		}

		req.SetQueryParam("limit", strconv.Itoa(options.Limit))
		req.SetQueryParam("offset", strconv.Itoa(options.Offset))

		if options.Sort != "" {
			req.SetQueryParam("sort", string(options.Sort))
		}
	}

	req.SetResult(&PeerResponse{})
	req.SetError(Error{})

	res, err := req.Get("api/peers")
	if err != nil {
		return nil, err
	}

	return res.Result().(*PeerResponse), res.Error().(error)
}
