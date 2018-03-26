package api

import (
	"context"
	"errors"
	"strconv"
)

type (
	DelegateRequest struct {
		Address         string
		PublicKey       string
		SecondPublicKey string
		Username        string
		Rank            string
	}

	ForgingStatsRequest struct {
		Address string

		FromTimestamp int64
		ToTimestamp   int64
	}

	DelegatesRequest struct {
		Address         string
		PublicKey       string
		SecondPublicKey string
		Username        string
		Rank            string

		ListOptions
	}

	DelegateResponse struct {
		Delegate *Delegate `json:"data"`
		*GenericResponse
	}

	DelegatesResponse struct {
		Delegates []*Delegate `json:"data"`
		*GenericResponse
	}

	NextForgersResponse struct {
		NextForgers []*DelegateWithSlot `json:"data"`
		*GenericResponse
	}

	DelegateWithSlot struct {
		Username  string `json:"username"`
		PublicKey string `json:"publicKey"`
		Address   string `json:"address"`
		NextSlot  int    `json:"nextSlot"`
	}

	Delegate struct {
		Username       string   `json:"username"`
		Vote           int64    `json:"vote,string"`
		Rewards        int      `json:"rewards,string"`
		ProducedBlocks int      `json:"producedBlocks"`
		MissedBlocks   int      `json:"missedBlocks"`
		Rate           int      `json:"rate"`
		Approval       float64  `json:"approval"`
		Productivity   float64  `json:"productivity"`
		Rank           int      `json:"rank"`
		Account        *Account `json:"account"`
	}

	ForgingStatsResponse struct {
		Stats ForgingStats `json:"data"`
		*GenericResponse
	}

	ForgingStats struct {
		Fees    string `json:"fees"`
		Rewards string `json:"rewards"`
		Forged  string `json:"forged"`
		Count   string `json:"count"`
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

// GetDelegates fuzzy searches for Delegates by username.
// Limit is set to 100 by default
// The Account field of the delegate response does not contain a username. Use the delegate username instead
func (c *Client) SearchDelegate(ctx context.Context, username string, listOptions *ListOptions) (*DelegatesResponse, error) {
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
