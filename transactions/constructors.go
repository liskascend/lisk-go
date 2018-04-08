package transactions

import (
	"fmt"

	"github.com/liskascend/lisk-go/crypto"
)

// NewTransaction creates a new value transfer transaction and signs it using the given secrets.
// The second secret is optional and only required for lisk wallets with a second signature.
func NewTransaction(recipientID string, amount uint64, secret string, secondSecret string, timeOffset int64) (
	*Transaction, error) {
	timestamp := GetCurrentTimeWithOffset(timeOffset)

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

	return transaction, nil
}

// NewTransactionWithData creates a new value transfer transaction with data and signs it using the given secrets.
// The second secret is optional and only required for lisk wallets with a second signature.
// Data can be a string or byte slice with a maximum length of 64 bytes.
func NewTransactionWithData(
	recipientID string, amount uint64, secret string, secondSecret string, timeOffset int64, data interface{}) (
	*Transaction, error) {
	timestamp := GetCurrentTimeWithOffset(timeOffset)

	var castedData []byte

	switch v := data.(type) {
	case []byte:
		castedData = v
	case string:
		castedData = []byte(v)
	default:
		return nil, fmt.Errorf("data has invalid type: %T", v)
	}

	transaction := &Transaction{
		Type:        TransactionTypeNormal,
		Amount:      amount,
		RecipientID: recipientID,
		Timestamp:   timestamp,
		Asset:       DataAsset(castedData),
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

	return transaction, nil
}

// NewSecondSignatureTransaction creates a new transaction to create a second signature
// and signs it using the given secrets.
func NewSecondSignatureTransaction(
	recipientID string, secret string, newSecondSecret string, timeOffset int64) (
	*Transaction, error) {
	timestamp := GetCurrentTimeWithOffset(timeOffset)

	secondSecretPublicKey := crypto.GetPublicKeyFromSecret(newSecondSecret)

	transaction := &Transaction{
		Type:        TransactionTypeSecondSecretRegistration,
		Amount:      0,
		RecipientID: recipientID,
		Timestamp:   timestamp,
		Asset: &RegisterSecondSignatureAsset{
			PublicKey: secondSecretPublicKey,
		},
	}

	pubKey := crypto.GetPublicKeyFromSecret(secret)
	transaction.SenderPublicKey = pubKey

	privKey := crypto.GetPrivateKeyFromSecret(secret)
	transaction.Sign(privKey)

	return transaction, nil
}

// NewVoteTransaction creates a new vote transaction and signs it using the given secrets.
// The second secret is optional and only required for lisk wallets with a second signature.
// The votes and unvotes are binary representations of the public keys of the relevant delegates.
func NewVoteTransaction(recipientID string, secret string, secondSecret string, timeOffset int64,
	votes [][]byte, unvotes [][]byte) (
	*Transaction, error) {
	timestamp := GetCurrentTimeWithOffset(timeOffset)

	transaction := &Transaction{
		Type:        TransactionTypeVote,
		Amount:      0,
		RecipientID: recipientID,
		Timestamp:   timestamp,
		Asset: &CastVoteAsset{
			Votes:   votes,
			Unvotes: unvotes,
		},
	}

	if valid, err := transaction.Asset.IsValid(); !valid {
		return nil, err
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

	return transaction, nil
}

// NewMultisignatureRegistrationTransaction creates a new transaction to create/update multisignature accounts
// and signs it using the given secrets.
// The second secret is optional and only required for lisk wallets with a second signature.
// The keys are binary representations of the public keys of the relevant delegates.
// Lifetime is the pending transaction lifetime.
// Min is the minimum number of signatures required.
func NewMultisignatureRegistrationTransaction(recipientID string, secret string, secondSecret string, timeOffset int64,
	addKeys [][]byte, removeKeys [][]byte, Lifetime byte, min byte) (
	*Transaction, error) {
	timestamp := GetCurrentTimeWithOffset(timeOffset)

	transaction := &Transaction{
		Type:        TransactionTypeMultisignatureRegistration,
		Amount:      0,
		RecipientID: recipientID,
		Timestamp:   timestamp,
		Asset: &RegisterMultisignatureAccountAsset{
			AddKeys:    addKeys,
			RemoveKeys: removeKeys,
			Lifetime:   Lifetime,
			Min:        min,
		},
	}

	if valid, err := transaction.Asset.IsValid(); !valid {
		return nil, err
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

	return transaction, nil
}
