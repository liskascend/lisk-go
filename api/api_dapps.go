package api

import (
	"context"
	"strconv"
)

type (
	DappRequest struct {
		TransactionID string
		Name          string

		ListOptions
	}

	DappResponse struct {
		Dapps []*Dapp `json:"data"`
		*GenericResponse
	}

	Dapp struct {
		TransactionID string `json:"transactionId"`
		Icon          string `json:"icon"`
		Category      int    `json:"category"`
		Type          int    `json:"type"`
		Link          string `json:"link"`
		Tags          string `json:"tags"`
		Description   string `json:"description"`
		Name          string `json:"name"`
	}
)

// GetDapps searches for Dapps on the blockchain.
// Search parameters can be specified in options.
// Limit is set to 100 by default
func (c *Client) GetDapps(ctx context.Context, options *DappRequest) (*DappResponse, error) {
	req := c.restClient.R().SetContext(ctx)

	if options != nil {
		if options.TransactionID != "" {
			req.SetQueryParam("transactionId", options.TransactionID)
		}
		if options.Name != "" {
			req.SetQueryParam("name", options.Name)
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

	req.SetResult(&DappResponse{})
	req.SetError(Error{})

	res, err := req.Get("api/dapps")
	if err != nil {
		return nil, err
	}

	return res.Result().(*DappResponse), res.Error().(error)
}
