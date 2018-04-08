package api

import (
	"context"
	"time"
)

type (
	// ConstantsResponse is the API response for constants requests
	ConstantsResponse struct {
		// Constants is the result
		Constants *Constants `json:"data"`
		*GenericResponse
	}

	// Constants are the constants of the node
	Constants struct {
		// Epoch of the node
		Epoch time.Time `json:"epoch"`
		// Milestone of the node
		Milestone int64 `json:"milestone,string"`
		// Build of the node
		Build string `json:"build"`
		// Commit of the node
		Commit string `json:"commit"`
		// Version of the node
		Version string `json:"version"`
		// Nethash of the node
		Nethash string `json:"nethash"`
		// Supply of the node
		Supply string `json:"supply"`
		// Reward is the blockReward of the node
		Reward LiskAmount `json:"reward,string"`
		// Nonce of the node
		Nonce string `json:"nonce"`
		// Fees of the node
		Fees *Fees `json:"fees"`
	}

	// Fees are the transaction fees of the blockchain
	Fees struct {
		// Send is the fee for a send transaction
		Send LiskAmount `json:"send,string"`
		// Vote is the fee for a vote transaction
		Vote LiskAmount `json:"vote,string"`
		// SecondSignature is the fee for a second signature registration transaction
		SecondSignature LiskAmount `json:"secondSignature,string"`
		// Delegate is the fee for a delegate registration transaction
		Delegate LiskAmount `json:"delegate,string"`
		// Multisignature is the fee for a multisignature creation/update transaction
		Multisignature LiskAmount `json:"multisignature,string"`
		// DappRegistration is the fee for a dapp creation transaction
		DappRegistration LiskAmount `json:"dappRegistration,string"`
		// DappWithdrawal is the fee for a dapp withdrawal transaction
		DappWithdrawal LiskAmount `json:"dappWithdrawal,string"`
		// DappDeposit is the fee for a dapp deposit transaction
		DappDeposit LiskAmount `json:"dappDeposit,string"`
		// Data is the additional fee for a data transaction
		Data LiskAmount `json:"data,string"`
	}

	// NodeStatusReponse is the API response for node status requests
	NodeStatusReponse struct {
		// NodeStatus is the result
		NodeStatus *NodeStatus `json:"data"`
		*GenericResponse
	}

	// NodeStatus is the status of a node
	NodeStatus struct {
		// Broadhash of the node
		Broadhash string `json:"broadhash"`
		// Consensus status of the node
		Consensus int `json:"consensus"`
		// Height is the current height of the node
		Height int `json:"height"`
		// Loaded indicates whether the blockchain loading is finished
		Loaded bool `json:"loaded"`
		// NetworkHeight is the network height
		NetworkHeight int `json:"networkHeight"`
		// Syncing indicates whether the node is currently syncing the blockchain
		Syncing bool `json:"syncing"`
		// Transactions is information on the (pending) transactions
		Transactions struct {
			// Unconfirmed is the number of unconfirmed transactions in the pool
			Unconfirmed int `json:"unconfirmed"`
			// Unsigned is the number of unsigned transactions in the pool
			Unsigned int `json:"unsigned"`
			// Unprocessed is the number of unprocessed transactions in the pool
			Unprocessed int `json:"unprocessed"`
			// Confirmed is the number of confirmed transactions
			Confirmed int `json:"confirmed"`
			// Total is the total number of transactions
			Total int `json:"total"`
		} `json:"transactions"`
	}

	// ForgingStatusRequest is the request body to request the forging status of a node
	ForgingStatusRequest struct {
		// PublicKey is the public key for which the status should be requested
		PublicKey string
	}

	// ForgingStatusResponse is the API response for forging status requests
	ForgingStatusResponse struct {
		// ForgingStatus are the results
		ForgingStatus []*ForgingStatus `json:"data"`
		*GenericResponse
	}

	// ForgingStatus is the forging status of a node
	ForgingStatus struct {
		// Forging indicates whether a node is currently forging
		Forging bool `json:"forging"`
		// PublicKey is the key the node is forging with
		PublicKey string `json:"publicKey"`
	}

	// ForgingToggleRequest is the request body to toggle forging on a node
	ForgingToggleRequest struct {
		DecryptionKey string `json:"decryptionKey"`
		PublicKey     string `json:"publicKey"`
	}

	// ForgingToggleResponse is the API response for forging toggle requests
	ForgingToggleResponse struct {
		ForgingStatus *ForgingStatus `json:"data"`
		*GenericResponse
	}
)

// GetConstants returns the chain constants.
func (c *Client) GetConstants(ctx context.Context) (*ConstantsResponse, error) {
	req := c.restClient.R().SetContext(ctx)

	req.SetResult(&ConstantsResponse{})
	req.SetError(Error{})

	res, err := req.Get("api/node/constants")
	if err != nil {
		return nil, err
	}

	return res.Result().(*ConstantsResponse), res.Error().(error)
}

// GetNodeStatus returns the status of the node.
func (c *Client) GetNodeStatus(ctx context.Context) (*NodeStatusReponse, error) {
	req := c.restClient.R().SetContext(ctx)

	req.SetResult(&NodeStatusReponse{})
	req.SetError(Error{})

	res, err := req.Get("api/node/status")
	if err != nil {
		return nil, err
	}

	return res.Result().(*NodeStatusReponse), res.Error().(error)
}

// GetForgingStatus returns the forging status of the node.
// You can optionally specify a publicKey to query for to only return the forging status for that delegate.
func (c *Client) GetForgingStatus(ctx context.Context, options *ForgingStatusRequest) (*ForgingStatsResponse, error) {
	req := c.restClient.R().SetContext(ctx)

	if options != nil && options.PublicKey != "" {
		req.SetQueryParam("publicKey", options.PublicKey)
	}

	req.SetResult(&ForgingStatsResponse{})
	req.SetError(Error{})

	res, err := req.Get("api/node/status/forging")
	if err != nil {
		return nil, err
	}

	return res.Result().(*ForgingStatsResponse), res.Error().(error)
}

// ToggleForging toggles forging on a specific key.
func (c *Client) ToggleForging(ctx context.Context, options *ForgingToggleRequest) (*ForgingToggleResponse, error) {
	req := c.restClient.R().SetContext(ctx)

	req.SetBody(options)

	req.SetResult(&ForgingStatsResponse{})
	req.SetError(Error{})

	res, err := req.Put("api/node/status/forging")
	if err != nil {
		return nil, err
	}

	var forgingStatus *ForgingStatus
	if len(res.Result().(*ForgingStatusResponse).ForgingStatus) != 0 {
		forgingStatus = res.Result().(*ForgingStatusResponse).ForgingStatus[0]
	}

	result := &ForgingToggleResponse{
		ForgingStatus:   forgingStatus,
		GenericResponse: res.Result().(*DelegatesResponse).GenericResponse,
	}

	return result, res.Error().(error)
}
