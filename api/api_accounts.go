package api

import (
	"context"
	"strconv"
)

type (
	AccountRequest struct {
		Address         string
		PublicKey       string
		SecondPublicKey string
		Username        string

		ListOptions
	}

	AccountResponse struct {
		Accounts []*Account `json:"data"`
		*GenericResponse
	}

	Account struct {
		Address            string    `json:"address"`
		PublicKey          string    `json:"publicKey"`
		Balance            int64     `json:"balance,string"`
		UnconfirmedBalance int64     `json:"unconfirmedBalance,string"`
		SecondPublicKey    string    `json:"secondPublicKey"`
		Delegate           *Delegate `json:"delegate,omitempty"`
	}
)

// GetAccounts searches for accounts on the blockchain.
// Search parameters can be specified in options.
// Limit is set to 100 by default
// The Account field of the Delegate field in the response is empty as it would cause a circular reference.
func (c *Client) GetAccounts(ctx context.Context, options *AccountRequest) (*AccountResponse, error) {
	req := c.restClient.R().SetContext(ctx)

	if options != nil {
		if options.Address != "" {
			req.SetQueryParam("address", options.Address)
		}
		if options.PublicKey != "" {
			req.SetQueryParam("publicKey", options.PublicKey)
		}
		if options.SecondPublicKey != "" {
			req.SetQueryParam("secondPublicKey", options.SecondPublicKey)
		}
		if options.Username != "" {
			req.SetQueryParam("username", options.Username)
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

	req.SetResult(&AccountResponse{})
	req.SetError(Error{})

	res, err := req.Get("api/accounts")
	if err != nil {
		return nil, err
	}

	return res.Result().(*AccountResponse), res.Error().(error)
}