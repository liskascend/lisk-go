package api

import (
	"context"
	"strconv"
)

type (
	// DelegateVoterRequest is the request body for a delegate voter request
	DelegateVoterRequest VoterRequest

	// VoterRequest is the request body for a voter request
	VoterRequest struct {
		// Address of the voter
		Address string
		// PublicKey of the voter
		PublicKey string
		// SecondPublicKey of the voter
		SecondPublicKey string
		// Username of the voter
		Username string

		ListOptions
	}

	// DelegateVoterResponse is the API response for delegate voter requests
	DelegateVoterResponse struct {
		// DelegateWithVoters is the result
		DelegateWithVoters DelegateWithVoters `json:"data"`
		*GenericResponse
	}

	// DelegateWithVoters is a delegate with information on its voters
	DelegateWithVoters struct {
		// Username of the delegate
		Username string `json:"username"`
		// PublicKey of the delegate
		PublicKey string `json:"publicKey,omitempty"`
		// Username of the delegate
		Votes int32 `json:"votes"`
		// Address of the delegate
		Address string `json:"address"`
		// Balance of the delegate
		Balance string `json:"balance"`
		// Voters of the delegate
		Voters []*Voter `json:"voters"`
	}

	// Voter is the detail information for a voter
	Voter struct {
		// Address of the voter
		Address string `json:"address"`
		// PublicKey of the voter
		PublicKey string `json:"publicKey"`
		// Balance of the voter
		Balance int64 `json:"balance"`
	}

	// VotesResponse is the API response for voter requests
	VotesResponse struct {
		// VoteData is the result
		VoteData VotesData
		*GenericResponse
	}

	// VotesData is detailed information on a users votes
	VotesData struct {
		// Address of the voter
		Address string `json:"address"`
		// Balance of the voter
		Balance int64 `json:"balance"`
		// Username of the voter
		Username string `json:"username"`
		// PublicKey of the voter
		PublicKey string `json:"publicKey"`
		// VotesUsed is the number of votes used
		VotesUsed int `json:"votesUsed"`
		// VotesAvailable is the number of votes available
		VotesAvailable int `json:"votesAvailable"`
		// Votes are the votes of the voter
		Votes []struct {
			// Address of the delegate
			Address string `json:"address"`
			// PublicKey of the delegate
			PublicKey string `json:"publicKey"`
			// Balance of the delegate
			Balance int64 `json:"balance"`
			// Username of the delegate
			Username string `json:"username"`
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
