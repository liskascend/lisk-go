package api

import (
	"context"
	"strconv"
)

type (
	BlockRequest struct {
		BlockID            string
		Height             *int64
		GeneratorPublicKey string

		ListOptions
	}

	BlockResponse struct {
		Blocks []*Block `json:"data"`
		*GenericResponse
	}

	Block struct {
		ID                   string `json:"id"`
		Version              int    `json:"version"`
		Height               int    `json:"height"`
		Timestamp            int    `json:"timestamp"`
		GeneratorAddress     string `json:"generatorAddress"`
		GeneratorPublicKey   string `json:"generatorPublicKey"`
		optionsLength        int    `json:"optionsLength"`
		optionsHash          string `json:"optionsHash"`
		BlockSignature       string `json:"blockSignature"`
		Confirmations        int    `json:"confirmations"`
		PreviousBlockID      string `json:"previousBlockId"`
		NumberOfTransactions int    `json:"numberOfTransactions"`
		TotalAmount          string `json:"totalAmount"`
		TotalFee             string `json:"totalFee"`
		Reward               string `json:"reward"`
		TotalForged          string `json:"totalForged"`
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
