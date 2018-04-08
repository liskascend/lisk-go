package api

import (
	"context"

	"github.com/liskascend/lisk-go/transactions"
)

type (
	// SignatureSendResponse is the API response for signature requests
	SignatureSendResponse struct {
		Result struct {
			Message string `json:"message"`
		} `json:"data"`
		*GenericResponse
	}
)

// SendSignature submits the signature for a multisignature transaction to the network.
func (c *Client) SendSignature(ctx context.Context, transaction *transactions.Transaction) (*TransactionSendResponse, error) {
	req := c.restClient.R().SetContext(ctx)

	req.SetBody(transaction)

	req.SetResult(&TransactionSendResponse{})
	req.SetError(Error{})

	res, err := req.Post("api/signatures")
	if err != nil {
		return nil, err
	}

	return res.Result().(*TransactionSendResponse), res.Error().(error)
}
