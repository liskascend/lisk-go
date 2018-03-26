package api

import (
	"context"
	"strconv"
)

type (
	DelegateVoterRequest VoterRequest

	VoterRequest struct {
		Address         string
		PublicKey       string
		SecondPublicKey string
		Username        string

		ListOptions
	}

	DelegateVoterResponse struct {
		DelegateWithVoters DelegateWithVoters `json:"data"`
		*GenericResponse
	}

	DelegateWithVoters struct {
		Username  string   `json:"username"`
		PublicKey string   `json:"publicKey,omitempty"`
		Votes     int32    `json:"votes"`
		Address   string   `json:"address"`
		Balance   string   `json:"balance"`
		Voters    []*Voter `json:"voters"`
	}

	Voter struct {
		Address   string `json:"address"`
		PublicKey string `json:"publicKey"`
		Balance   int64  `json:"balance"`
	}

	VotesResponse struct {
		VoteData VotesData
		*GenericResponse
	}

	VotesData struct {
		Address        string `json:"address"`
		Balance        int64  `json:"balance"`
		Username       string `json:"username"`
		PublicKey      string `json:"publicKey"`
		VotesUsed      int    `json:"votesUsed"`
		VotesAvailable int    `json:"votesAvailable"`
		Votes []struct {
			Address   string `json:"address"`
			PublicKey string `json:"publicKey"`
			Balance   int64  `json:"balance"`
			Username  string `json:"username"`
		} `json:"votes"`
	}
)

// GetDelegateVoters returns the voters for a specific delegate.
// Search parameters can be specified in options.
// Limit is set to 100 by default
func (c *Client) GetDelegateVoters(ctx context.Context, options *DelegateVoterRequest) (*DelegateVoterResponse, error) {
	req := c.restClient.R().SetContext(ctx)

	if options != nil {
		if options.Username != "" {
			req.SetQueryParam("username", options.Username)
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

	req.SetResult(&DelegateVoterResponse{})
	req.SetError(Error{})

	res, err := req.Get("api/voters")
	if err != nil {
		return nil, err
	}

	return res.Result().(*DelegateVoterResponse), res.Error().(error)
}

// GetVotes returns the votes that a specific address has casted.
// Search parameters can be specified in options.
// Limit is set to 100 by default
func (c *Client) GetVotes(ctx context.Context, options *VoterRequest) (*VotesResponse, error) {
	req := c.restClient.R().SetContext(ctx)

	if options != nil {
		if options.Username != "" {
			req.SetQueryParam("username", options.Username)
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

	req.SetResult(&VotesResponse{})
	req.SetError(Error{})

	res, err := req.Get("api/votes")
	if err != nil {
		return nil, err
	}

	return res.Result().(*VotesResponse), res.Error().(error)
}
