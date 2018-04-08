package api

import (
	"context"
	"errors"
	"strconv"
)

type (
	// DelegateRequest is the request body for a single Delegate request
	DelegateRequest struct {
		// Address of the delegate
		Address string
		// PublicKey of the delegate
		PublicKey string
		// SecondPublicKey of the delegate
		SecondPublicKey string
		// Username of the delegate
		Username string
		// Rank of the delegate
		Rank string
	}

	// ForgingStatsRequest is the request body for a forgingStats request
	ForgingStatsRequest struct {
		// Address of the delegate
		Address string

		// FromTimestamp is the starting point of the stats
		FromTimestamp int64
		// ToTimestamp is the ending point of the stats
		ToTimestamp int64
	}

	// DelegatesRequest is the request body to request Delegates
	DelegatesRequest struct {
		// Address of the delegate
		Address string
		// PublicKey of the delegate
		PublicKey string
		// SecondPublicKey of the delegate
		SecondPublicKey string
		// Username of the delegate
		Username string
		// Rank of the delegate
		Rank string

		ListOptions
	}

	// DelegateResponse is the API response for single Delegate requests
	DelegateResponse struct {
		// Delegate is the resulting delegate
		Delegate *Delegate `json:"data"`
		*GenericResponse
	}

	// DelegatesResponse is the API response for Delegate requests
	DelegatesResponse struct {
		// Delegates are the results
		Delegates []*Delegate `json:"data"`
		*GenericResponse
	}

	// NextForgersResponse is the API response for the next forger request
	NextForgersResponse struct {
		// NextForgers are the results
		NextForgers []*DelegateWithSlot `json:"data"`
		*GenericResponse
	}

	// DelegateWithSlot is a delegate with its next forging slot
	DelegateWithSlot struct {
		// Username of the delegate
		Username string `json:"username"`
		// PublicKey of the delegate
		PublicKey string `json:"publicKey"`
		// Address of the delegate
		Address string `json:"address"`
		// NextSlot of the delegate
		NextSlot int `json:"nextSlot"`
	}

	// Delegate is a delegate on the Lisk blockchain
	Delegate struct {
		// Username of the delegate
		Username string `json:"username"`
		// Vote amount of the delegate
		Vote int64 `json:"vote,string"`
		// Rewards of the delegate
		Rewards int `json:"rewards,string"`
		// ProducedBlocks is the number of blocks produced by the delegate
		ProducedBlocks int `json:"producedBlocks"`
		// MissedBlocks is the number of blocks missed by the delegate
		MissedBlocks int `json:"missedBlocks"`
		// Rate of the delegate
		Rate int `json:"rate"`
		// Approval percentage of the delegate
		Approval float64 `json:"approval"`
		// Productivity of the delegate
		Productivity float64 `json:"productivity"`
		// Rank of the delegate
		Rank int `json:"rank"`
		// Account of the delegate
		Account *Account `json:"account"`
	}

	// ForgingStatsResponse is the API response for forgingStats requests
	ForgingStatsResponse struct {
		// Stats is the result
		Stats ForgingStats `json:"data"`
		*GenericResponse
	}

	// ForgingStats are the forgingStats of a delegate
	ForgingStats struct {
		// Fees forged by the delegate
		Fees string `json:"fees"`
		// Rewards forged by the delegate
		Rewards string `json:"rewards"`
		// Forged is the total amount forged by the delegate
		Forged string `json:"forged"`
		// Count of blocks forged
		Count string `json:"count"`
	}
)

// GetDelegate gets a Delegate from the blockchain.
// Query parameters can be specified in options.
// The Account field of the delegate response does not contain a username. Use the delegate username instead
func (c *Client) GetDelegate(ctx context.Context, options *DelegateRequest) (*DelegateResponse, error) {
	req := c.restClient.R().SetContext(ctx)

	if options == nil {
		return nil, errors.New("options is <nil>, please specify your query parameters")
	}

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

	req.SetResult(&DelegatesResponse{})
	req.SetError(Error{})

	res, err := req.Get("api/delegates")
	if err != nil {
		return nil, err
	}

	var delegate *Delegate
	if len(res.Result().(*DelegatesResponse).Delegates) > 0 {
		delegate = res.Result().(*DelegatesResponse).Delegates[0]
	}

	result := &DelegateResponse{
		Delegate:        delegate,
		GenericResponse: res.Result().(*DelegatesResponse).GenericResponse,
	}

	return result, res.Error().(error)

}

// SearchDelegates fuzzy searches for Delegates by username.
// Limit is set to 100 by default
// The Account field of the delegate response does not contain a username. Use the delegate username instead
func (c *Client) SearchDelegates(ctx context.Context, username string, listOptions *ListOptions) (*DelegatesResponse, error) {
	req := c.restClient.R().SetContext(ctx)

	if username != "" {
		req.SetQueryParam("search", username)
	}

	if listOptions != nil {
		if listOptions.Limit <= 0 {
			listOptions.Limit = 100
		}

		if listOptions.Offset < 0 {
			listOptions.Offset = 0
		}

		req.SetQueryParam("limit", strconv.Itoa(listOptions.Limit))
		req.SetQueryParam("offset", strconv.Itoa(listOptions.Offset))

		if listOptions.Sort != "" {
			req.SetQueryParam("sort", string(listOptions.Sort))
		}
	}

	req.SetResult(&DelegatesResponse{})
	req.SetError(Error{})

	res, err := req.Get("api/dapps")
	if err != nil {
		return nil, err
	}

	return res.Result().(*DelegatesResponse), res.Error().(error)
}

// GetNextForgers returns the next forging delegates.
// Limit is set to 100 by default
// The Account field of the delegate response does not contain a username. Use the delegate username instead
func (c *Client) GetNextForgers(ctx context.Context, listOptions *ListOptions) (*NextForgersResponse, error) {
	req := c.restClient.R().SetContext(ctx)

	if listOptions != nil {
		if listOptions.Limit <= 0 {
			listOptions.Limit = 100
		}

		if listOptions.Offset < 0 {
			listOptions.Offset = 0
		}

		req.SetQueryParam("limit", strconv.Itoa(listOptions.Limit))
		req.SetQueryParam("offset", strconv.Itoa(listOptions.Offset))

		if listOptions.Sort != "" {
			req.SetQueryParam("sort", string(listOptions.Sort))
		}
	}
	req.SetResult(&NextForgersResponse{})
	req.SetError(Error{})

	res, err := req.Get("api/delegates/forgers")
	if err != nil {
		return nil, err
	}

	return res.Result().(*NextForgersResponse), res.Error().(error)
}

// GetForgingStats returns the forgingStats for a delegate.
func (c *Client) GetForgingStats(ctx context.Context, options *ForgingStatsRequest) (*ForgingStatsResponse, error) {
	req := c.restClient.R().SetContext(ctx)

	req.SetPathParams(
		map[string]string{
			"address": options.Address,
		})

	if options.FromTimestamp <= 0 {
		options.FromTimestamp = 0
	}

	if options.ToTimestamp <= 0 {
		options.ToTimestamp = 0
	}

	req.SetQueryParam("fromTimestamp", strconv.FormatInt(options.FromTimestamp, 10))
	if options.ToTimestamp != 0 {
		req.SetQueryParam("toTimestamp", strconv.FormatInt(options.ToTimestamp, 10))
	}

	req.SetResult(&ForgingStatsResponse{})
	req.SetError(Error{})

	res, err := req.Get("api/delegates/{address}/forging_statistics")
	if err != nil {
		return nil, err
	}

	return res.Result().(*ForgingStatsResponse), res.Error().(error)
}
