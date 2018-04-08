package api

import (
	"context"
	"strconv"
)

type (
	// BlockRequest is the request body to request blocks
	BlockRequest struct {
		// BlockID of the block
		BlockID string
		// Height if the block
		Height *int64
		// GeneratorPublicKey is the public key of the delegate that forged the block
		GeneratorPublicKey string

		ListOptions
	}

	// BlockResponse is the API response for block requests
	BlockResponse struct {
		// Blocks are the results
		Blocks []*Block `json:"data"`
		*GenericResponse
	}

	// Block is a Lisk block
	Block struct {
		// ID of the block
		ID string `json:"id"`
		// Version of the block
		Version int `json:"version"`
		// Height of the block
		Height int `json:"height"`
		// Timestamp of the block
		Timestamp int `json:"timestamp"`
		// GeneratorAddress of the block
		GeneratorAddress string `json:"generatorAddress"`
		// GeneratorPublicKey of the block
		GeneratorPublicKey string `json:"generatorPublicKey"`
		// OptionsLength of the block
		OptionsLength int `json:"optionsLength"`
		// OptionsHash of the block
		OptionsHash string `json:"optionsHash"`
		// BlockSignature of the block
		BlockSignature string `json:"blockSignature"`
		// Confirmations of the block
		Confirmations int `json:"confirmations"`
		// PreviousBlockID of the block
		PreviousBlockID string `json:"previousBlockId"`
		// NumberOfTransactions of the block
		NumberOfTransactions int `json:"numberOfTransactions"`
		// TotalAmount of the block
		TotalAmount string `json:"totalAmount"`
		// TotalFee of the block
		TotalFee string `json:"totalFee"`
		// Reward of the block
		Reward string `json:"reward"`
		// TotalForged of the block
		TotalForged string `json:"totalForged"`
	}
)

// GetBlocks searches for blocks on the blockchain.
// Search parameters can be specified in options.
// Limit is set to 100 by default
func (c *Client) GetBlocks(ctx context.Context, options *BlockRequest) (*BlockResponse, error) {
	req := c.restClient.R().SetContext(ctx)

	if options != nil {
		if options.BlockID != "" {
			req.SetQueryParam("blockId", options.BlockID)
		}
		if options.Height != nil {
			req.SetQueryParam("height", strconv.FormatInt(*options.Height, 10))
		}
		if options.GeneratorPublicKey != "" {
			req.SetQueryParam("generatorPublicKey", options.GeneratorPublicKey)
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

	req.SetResult(&BlockResponse{})
	req.SetError(Error{})

	res, err := req.Get("api/blocks")
	if err != nil {
		return nil, err
	}

	return res.Result().(*BlockResponse), res.Error().(error)
}
