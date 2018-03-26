package api

import (
	"context"
	"strconv"

	"github.com/liskascend/lisk-go/transactions"
)

type (
	TransactionRequest struct {
		ID                 string
		RecipientID        string
		RecipientPublicKey string
		SenderID           string
		SenderPublicKey    string
		BlockID            string
		Type               *int
		Height             *int64
		MinAmount          *int64
		MaxAmount          *int64

		FromTimestamp int64
		ToTimestamp   int64
		ListOptions
	}

	TransactionsResponse struct {
		Transactions []*Transaction `json:"data"`
		*GenericResponse
	}
	TransactionSendResponse struct {
		Result struct {
			Message string `json:"message"`
		} `json:"data"`
		*GenericResponse
	}
)

// GetTransactions searches for transactions on the blockchain.
// Search parameters can be specified in options.
// Limit is set to 100 by default
func (c *Client) GetTransactions(ctx context.Context, options *TransactionRequest) (*TransactionsResponse, error) {
	req := c.restClient.R().SetContext(ctx)

	if options != nil {
		if options.ID != "" {
			req.SetQueryParam("id", options.ID)
		}
		if options.RecipientID != "" {
			req.SetQueryParam("recipientId", options.RecipientID)
		}
		if options.RecipientPublicKey != "" {
			req.SetQueryParam("recipientPublicKey", options.RecipientPublicKey)
		}
		if options.SenderID != "" {
			req.SetQueryParam("senderId", options.SenderID)
		}
		if options.SenderPublicKey != "" {
			req.SetQueryParam("senderPublicKey", options.SenderPublicKey)
		}
		if options.BlockID != "" {
			req.SetQueryParam("blockId", options.BlockID)
		}

		if options.Type != nil {
			req.SetQueryParam("type", strconv.Itoa(*options.Type))
		}
		if options.Height != nil {
			req.SetQueryParam("height", strconv.FormatInt(*options.Height, 10))
		}
		if options.MinAmount != nil {
			req.SetQueryParam("minAmount", strconv.FormatInt(*options.MinAmount, 10))
		}
		if options.MaxAmount != nil {
			req.SetQueryParam("maxAmount", strconv.FormatInt(*options.MaxAmount, 10))
		}

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

	req.SetResult(&TransactionsResponse{})
	req.SetError(Error{})

	res, err := req.Get("api/transactions")
	if err != nil {
		return nil, err
	}

	return res.Result().(*TransactionsResponse), res.Error().(error)
}

// SendTransaction submits the transaction to the network.
func (c *Client) SendTransaction(ctx context.Context, transaction *transactions.Transaction) (*TransactionSendResponse, error) {
	req := c.restClient.R().SetContext(ctx)

	req.SetBody(transaction)

	req.SetResult(&TransactionSendResponse{})
	req.SetError(Error{})

	res, err := req.Post("api/transactions")
	if err != nil {
		return nil, err
	}

	return res.Result().(*TransactionSendResponse), res.Error().(error)
}
