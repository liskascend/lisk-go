package api

import (
	"context"
	"strconv"
)

type (
	// DappRequest is the request body to request Dapps
	DappRequest struct {
		// TransactionID is the transactionID of the Dapp creation
		TransactionID string
		// Name is the name of the Dapp
		Name string

		ListOptions
	}

	// DappResponse is the API response for Dapp requests
	DappResponse struct {
		// Dapps are the results
		Dapps []*Dapp `json:"data"`
		*GenericResponse
	}

	// Dapp is a Lisk Dapp
	Dapp struct {
		// TransactionID is the transactionID of the Dapp creation
		TransactionID string `json:"transactionId"`
		// Icon is the icon of the Dapp
		Icon string `json:"icon"`
		// Category is the category of the Dapp
		Category int `json:"category"`
		// Type is the type of the Dapp
		Type int `json:"type"`
		// Link is the link to the Dapp
		Link string `json:"link"`
		// Tags are the tags of the Dapp
		Tags string `json:"tags"`
		// Description is the description of the Dapp
		Description string `json:"description"`
		// Name is the name of the Dapp
		Name string `json:"name"`
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
