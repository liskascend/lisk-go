package api

import (
	"context"
	"strconv"
)

type (
	// PeerRequest is the request body for a peer request
	PeerRequest struct {
		// IP of the peer
		IP string
		// HTTPPort of the peer
		HTTPPort *int
		// WSPort of the peer
		WSPort *int
		// OS of the peer
		OS string
		// Version of the peer
		Version string
		// State of the peer
		State *int
		// Height of the peer
		Height *int64
		// Broadhash of the peer
		Broadhash string

		ListOptions
	}

	// PeerResponse is the API response for peer requests
	PeerResponse struct {
		// Peers are the results
		Peers []*Peer `json:"data"`
		*GenericResponse
	}

	// Peer is a peer the node is connected to
	Peer struct {
		// IP of the peer
		IP string `json:"ip"`
		// HTTPPort of the peer
		HTTPPort int `json:"httpPort"`
		// WSPort of the peer
		WSPort int `json:"wsPort"`
		// OS of the peer
		OS string `json:"os"`
		// Version of the peer
		Version string `json:"version"`
		// State of the peer
		State int `json:"state"`
		// Height of the peer
		Height int64 `json:"height"`
		// Broadhash of the peer
		Broadhash string `json:"broadhash"`
		// Nonce of the peer
		Nonce string `json:"nonce"`
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
		if options.HTTPPort != nil {
			req.SetQueryParam("httpPort", strconv.Itoa(*options.HTTPPort))
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
