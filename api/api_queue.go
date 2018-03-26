package api

import (
	"context"
	"strconv"
	"time"
)

type (
	QueueRequest struct {
		ID                 string
		RecipientID        string
		RecipientPublicKey string
		SenderID           string
		SenderPublicKey    string
		Type               *int

		ListOptions
	}

	QueueResponse struct {
		Transactions []*Transaction `json:"data"`
		*GenericResponse
	}

	Transaction struct {
		ID                    string    `json:"id"`
		Amount                string    `json:"amount"`
		Fee                   string    `json:"fee"`
		Type                  int       `json:"type"`
		Height                int       `json:"height"`
		BlockID               string    `json:"blockId"`
		Timestamp             int       `json:"timestamp"`
		SenderID              string    `json:"senderId"`
		SenderPublicKey       string    `json:"senderPublicKey"`
		SenderSecondPublicKey string    `json:"senderSecondPublicKey"`
		RecipientID           string    `json:"recipientId"`
		RecipientPublicKey    string    `json:"recipientPublicKey"`
		Signature             string    `json:"signature"`
		SignSignature         string    `json:"signSignature"`
		Multisignatures       []string  `json:"signatures"`
		Confirmations         int       `json:"confirmations"`
		Asset                 struct{}  `json:"asset"`
		ReceivedAt            time.Time `json:"receivedAt"`
		Relays                int       `json:"relays"`
		Ready                 bool      `json:"ready"`
	}

	TransactionState string
)

const (
	TransactionStateUnprocessed TransactionState = "unprocessed"
	TransactionStateUnconfirmed TransactionState = "unconfirmed"
	TransactionStateUnsigned    TransactionState = "unsigned"
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
