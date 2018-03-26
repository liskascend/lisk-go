package transactions

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"errors"

	"github.com/liskascend/lisk-go/crypto"
	"golang.org/x/crypto/ed25519"
)

// Serialize serializes the transaction in the binary format used
func (t *Transaction) Serialize() ([]byte, error) {
	if valid, err := t.IsValid(); !valid {
		return nil, err
	}

	// Create new buffer
	dst := new(bytes.Buffer)

	// First byte is the transaction type
	binary.Write(dst, binary.LittleEndian, t.Type)

	// Append the timestamp
	binary.Write(dst, binary.LittleEndian, t.Timestamp)

	// Append the sender's public key
	transactionSenderPubKey := make([]byte, ed25519.PublicKeySize)
	copy(transactionSenderPubKey, t.SenderPublicKey)
	binary.Write(dst, binary.LittleEndian, transactionSenderPubKey)

	// Append the requester's public key if given
	if len(t.TransactionRequesterPublicKey) > 0 {
		binary.Write(dst, binary.LittleEndian, t.TransactionRequesterPublicKey)
	}

	// Append the recipientId
	transactionRecipientID := make([]byte, byteSizeRecipientID)
	if len(t.RecipientID) != 0 {
		numericAddress := new(big.Int)
		numericAddress.SetString(t.RecipientID[:len(t.RecipientID)-1], 10)
		numericAddressBytes := numericAddress.Bytes()
		copy(transactionRecipientID[cap(transactionRecipientID)-len(numericAddressBytes):], numericAddressBytes)
	}
	binary.Write(dst, binary.LittleEndian, transactionRecipientID)

	// Append the amount to be transferred
	binary.Write(dst, binary.LittleEndian, t.Amount)

	// Append asset data if given
	if t.Asset != nil {
		transactionAssetData, err := t.Asset.serialize()
		if err != nil {
			return nil, err
		}
		binary.Write(dst, binary.LittleEndian, transactionAssetData)
	}

	// Append signatures (both optional)
	if len(t.signature) > 0 {
		binary.Write(dst, binary.LittleEndian, t.signature)
	}
	if len(t.secondSignature) > 0 {
		binary.Write(dst, binary.LittleEndian, t.secondSignature)
	}

	return dst.Bytes(), nil
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

	// Pack the transaction in the serializable format
	preparedTransaction := &serializableTransaction{
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
	}

	// Add an empty asset because it's required
	if preparedTransaction.Asset == nil {
		preparedTransaction.Asset = struct{}{}
	}

	return json.Marshal(preparedTransaction)
}
