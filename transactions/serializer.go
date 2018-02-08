package transactions

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"errors"

	"github.com/slamper/lisk-go/crypto"
	"golang.org/x/crypto/ed25519"
)

// Serialize serializes the transaction in the binary format used
func (t *Transaction) Serialize() ([]byte, error) {
	if valid, err := t.IsValid(); !valid {
		return nil, err
	}

	var dst []byte

	// First byte is the transaction type
	transactionType := []byte{byte(t.Type)}
	dst = append(dst, transactionType...)

	// Append the timestamp
	transactionTimestamp := make([]byte, byteSizeTimestamp)
	binary.LittleEndian.PutUint32(transactionTimestamp, t.Timestamp)
	dst = append(dst, transactionTimestamp...)

	// Append the sender's public key
	transactionSenderPubKey := make([]byte, ed25519.PublicKeySize)
	copy(transactionSenderPubKey, t.SenderPublicKey)
	dst = append(dst, transactionSenderPubKey...)

	// Append the requester's public key if given
	if t.TransactionRequesterPublicKey != nil && len(t.TransactionRequesterPublicKey) != 0 {
		dst = append(dst, t.TransactionRequesterPublicKey...)
	}

	// Append the recipientId
	transactionRecipientID := make([]byte, byteSizeRecipientID)

	if len(t.RecipientID) != 0 {
		numericAddress := new(big.Int)
		numericAddress.SetString(t.RecipientID[:len(t.RecipientID)-1], 10)
		numericAddressBytes := numericAddress.Bytes()
		copy(transactionRecipientID[cap(transactionRecipientID)-len(numericAddressBytes):], numericAddressBytes)
	}

	dst = append(dst, transactionRecipientID...)

	// Append the amount to be transferred
	transactionAmount := make([]byte, byteSizeAmount)
	binary.LittleEndian.PutUint64(transactionAmount, t.Amount)
	dst = append(dst, transactionAmount...)

	// Append asset data if given
	if t.Asset != nil {
		transactionAssetData, err := t.Asset.serialize()
		if err != nil {
			return nil, err
		}
		dst = append(dst, transactionAssetData...)
	}

	// Append signatures (both optional)
	dst = append(dst, t.signature...)
	dst = append(dst, t.secondSignature...)

	return dst, nil
}

// IsValid returns whether the transaction is valid
func (t *Transaction) IsValid() (bool, error) {
	if len(t.SenderPublicKey) != ed25519.PublicKeySize {
		return false, errors.New("invalid or missing SenderPublicKey")
	}

	if t.Type > 7 {
		return false, errors.New("invalid transaction type")
	}

	// TODO implement further validation from server
	switch t.Type {
	case TransactionTypeNormal:
		if t.Asset != nil {
			if _, valid := t.Asset.(DataAsset); !valid {
				return false, errors.New("invalid asset type or missing asset")
			}
		}
	case TransactionTypeSecondSecretRegistration:
		if _, valid := t.Asset.(*RegisterSecondSignatureAsset); !valid {
			return false, errors.New("invalid asset type or missing asset")
		}
	case TransactionTypeDelegateRegistration:
		if _, valid := t.Asset.(*RegisterDelegateAsset); !valid {
			return false, errors.New("invalid asset type or missing asset")
		}
		if t.Amount > 0 {
			return false, errors.New("invalid amount; must be 0")
		}
	case TransactionTypeVote:
		if _, valid := t.Asset.(*CastVoteAsset); !valid {
			return false, errors.New("invalid asset type or missing asset")
		}
	case TransactionTypeMultisignatureRegistration:
		if _, valid := t.Asset.(*RegisterMultisignatureAccountAsset); !valid {
			return false, errors.New("invalid asset type or missing asset")
		}
	}

	if t.Asset != nil {
		if val, err := t.Asset.IsValid(); !val {
			return false, fmt.Errorf("invalid asset: %v", err)
		}
	}

	if len(t.signature) != 0 && len(t.signature) != byteSizeSignatureTransaction {
		return false, errors.New("signature has invalid size")
	}

	if len(t.secondSignature) != 0 && len(t.secondSignature) != byteSizeSecondSignatureTransaction {
		return false, errors.New("secondSignature has invalid size")
	}

	if len(t.TransactionRequesterPublicKey) != 0 && len(t.TransactionRequesterPublicKey) != ed25519.PublicKeySize {
		return false, errors.New("invalid transactionRequesterPubKey size")
	}

	return true, nil
}

// MarshalJSON converts the transaction to a JSON payload that can be sent to the node
func (t *Transaction) MarshalJSON() ([]byte, error) {
	if val, err := t.IsValid(); !val {
		return nil, fmt.Errorf("cannot marshal: %v", err)
	}

	id, err := t.ID()
	if err != nil {
		return nil, err
	}

	fee, err := t.Fee()
	if err != nil {
		return nil, err
	}

	return json.Marshal(&serializableTransaction{
		Type:                          t.Type,
		ID:                            id,
		SenderID:                      crypto.GetAddressFromPublicKey(t.SenderPublicKey),
		Amount:                        t.Amount,
		Fee:                           int(fee),
		RecipientID:                   t.RecipientID,
		Timestamp:                     t.Timestamp,
		Asset:                         t.Asset,
		SenderPublicKey:               hex.EncodeToString(t.SenderPublicKey),
		TransactionRequesterPublicKey: hex.EncodeToString(t.TransactionRequesterPublicKey),
		Signature:                     hex.EncodeToString(t.signature),
		SecondSignature:               hex.EncodeToString(t.secondSignature),
	})
}
