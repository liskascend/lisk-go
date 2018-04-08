package api

import (
	"context"
	"strconv"
	"time"
)

type (
	// QueueRequest is the request body for a transaction queue request
	QueueRequest struct {
		// ID of the transaction in the queue
		ID string
		// RecipientID of the transaction in the queue
		RecipientID string
		// RecipientPublicKey of the transaction in the queue
		RecipientPublicKey string
		// SenderID of the transaction in the queue
		SenderID string
		// SenderPublicKey of the transaction in the queue
		SenderPublicKey string
		// Type of the transaction in the queue
		Type *int

		ListOptions
	}

	// QueueResponse is the API response for transaction queue requests
	QueueResponse struct {
		// Transactions are the results
		Transactions []*Transaction `json:"data"`
		*GenericResponse
	}

	// Transaction is a transaction in the pool or on the blockchain
	Transaction struct {
		// ID of the transaction
		ID string `json:"id"`
		// Amount of the transaction
		Amount string `json:"amount"`
		// Fee of the transaction
		Fee string `json:"fee"`
		// Type of the transaction
		Type int `json:"type"`
		// Height of the transaction
		Height int `json:"height"`
		// BlockID of the transaction
		BlockID string `json:"blockId"`
		// Timestamp of the transaction
		Timestamp int `json:"timestamp"`
		// SenderID of the transaction
		SenderID string `json:"senderId"`
		// SenderPublicKey of the transaction
		SenderPublicKey string `json:"senderPublicKey"`
		// SenderSecondPublicKey of the transaction
		SenderSecondPublicKey string `json:"senderSecondPublicKey"`
		// RecipientID of the transaction
		RecipientID string `json:"recipientId"`
		// RecipientPublicKey of the transaction
		RecipientPublicKey string `json:"recipientPublicKey"`
		// Signature of the transaction
		Signature string `json:"signature"`
		// SignSignature of the transaction
		SignSignature string `json:"signSignature"`
		// Multisignatures of the transaction
		Multisignatures []string `json:"signatures"`
		// Confirmations of the transaction
		Confirmations int `json:"confirmations"`
		// Asset of the transaction
		Asset struct{} `json:"asset"`
		// ReceivedAt timestamp of the transaction
		ReceivedAt time.Time `json:"receivedAt"`
		// Relays is the number of times the transaction was relayed
		Relays int `json:"relays"`
		// Ready status of the transaction
		Ready bool `json:"ready"`
	}

	// TransactionState is the state of a transaction in the queue
	TransactionState string
)

const (
	// TransactionStateUnprocessed is the state of a transaction the was not processed yet
	TransactionStateUnprocessed TransactionState = "unprocessed"
	// TransactionStateUnconfirmed is the state of a transaction the was not confirmed yet
	TransactionStateUnconfirmed TransactionState = "unconfirmed"
	// TransactionStateUnsigned is the state of a transaction that is missing signatures
	TransactionStateUnsigned TransactionState = "unsigned"
)

// GetPendingTransactions searches for transactions with the given state.
// Search parameters can be specified in options.
// Limit is set to 100 by default
func (c *Client) GetPendingTransactions(ctx context.Context, state TransactionState, options *QueueRequest) (*QueueResponse, error) {
	req := c.restClient.R().SetContext(ctx)

	req.SetPathParams(map[string]string{
		"state": string(state),
	})
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
		if options.Type != nil {
			req.SetQueryParam("type", strconv.Itoa(*options.Type))
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

	req.SetResult(&QueueResponse{})
	req.SetError(Error{})

	res, err := req.Get("api/node/transactions/{state}")
	if err != nil {
		return nil, err
	}

	return res.Result().(*QueueResponse), res.Error().(error)
}
