package transactions

import (
	"github.com/liskascend/lisk-go/crypto"
)

// NewTransaction creates a new value transfer transaction and signs it using the given secrets.
// The second secret is optional and only required for lisk wallets with a second signature.
func NewTransaction(recipientID string, amount uint64, secret string, secondSecret string, timeOffset int64) *Transaction {
	timestamp := getCurrentTimeWithOffset(timeOffset)

	transaction := &Transaction{
		Type:        TransactionTypeNormal,
		Amount:      amount,
		RecipientID: recipientID,
		Timestamp:   timestamp,
	}

	pubKey := crypto.GetPublicKeyFromSecret(secret)
	transaction.SenderPublicKey = pubKey

	privKey := crypto.GetPrivateKeyFromSecret(secret)
	transaction.Sign(privKey)

	// Add a second signature if a second secret is given
	if secondSecret != "" {
		privKey2 := crypto.GetPrivateKeyFromSecret(secondSecret)
		transaction.SecondSign(privKey2)
	}

	return transaction
}
