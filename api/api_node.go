package api

import (
	"context"
	"time"
)

type (
	ConstantsResponse struct {
		Constants *Constants `json:"data"`
		*GenericResponse
	}

	Constants struct {
		Epoch     time.Time  `json:"epoch"`
		Milestone int64      `json:"milestone,string"`
		Build     string     `json:"build"`
		Commit    string     `json:"commit"`
		Version   string     `json:"version"`
		Nethash   string     `json:"nethash"`
		Supply    string     `json:"supply"`
		Reward    LiskAmount `json:"reward,string"`
		Nonce     string     `json:"nonce"`
		Fees      *Fees      `json:"fees"`
	}

	Fees struct {
		Send             LiskAmount `json:"send,string"`
		Vote             LiskAmount `json:"vote,string"`
		SecondSignature  LiskAmount `json:"secondSignature,string"`
		Delegate         LiskAmount `json:"delegate,string"`
		Multisignature   LiskAmount `json:"multisignature,string"`
		DappRegistration LiskAmount `json:"dappRegistration,string"`
		DappWithdrawal   LiskAmount `json:"dappWithdrawal,string"`
		DappDeposit      LiskAmount `json:"dappDeposit,string"`
		Data             LiskAmount `json:"data,string"`
	}

	NodeStatusReponse struct {
		NodeStatus *NodeStatus `json:"data"`
		*GenericResponse
	}

	NodeStatus struct {
		Broadhash     string `json:"broadhash"`
		Consensus     int    `json:"consensus"`
		Height        int    `json:"height"`
		Loaded        bool   `json:"loaded"`
		NetworkHeight int    `json:"networkHeight"`
		Syncing       bool   `json:"syncing"`
		Transactions struct {
			Unconfirmed int `json:"unconfirmed"`
			Unsigned    int `json:"unsigned"`
			Unprocessed int `json:"unprocessed"`
			Confirmed   int `json:"confirmed"`
			Total       int `json:"total"`
		} `json:"transactions"`
	}

	ForgingStatusRequest struct {
		PublicKey string
	}

	ForgingStatusResponse struct {
		ForgingStatus []*ForgingStatus `json:"data"`
		*GenericResponse
	}

	ForgingStatus struct {
		Forging   bool   `json:"forging"`
		PublicKey string `json:"publicKey"`
	}

	ForgingToggleRequest struct {
		DecryptionKey string `json:"decryptionKey"`
		PublicKey     string `json:"publicKey"`
	}

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

// GetConstants returns the chain constants.
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

// GetForgingStatus returns the forging status of the node.
// You can optionally specify a publicKey to query for to only return the forging status for that delegate.
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
