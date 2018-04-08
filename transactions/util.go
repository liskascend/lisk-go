package transactions

import (
	"bytes"
	"crypto/sha256"
	"errors"

	"github.com/liskascend/lisk-go/crypto"
)

// Hash returns the SHA256 hash of the transaction bytes
func (t *Transaction) Hash() ([]byte, error) {
	data, err := t.Serialize()
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(data)

	return hash[:], nil
}

// ID returns the ID of the transaction
func (t *Transaction) ID() (string, error) {
	hash, err := t.Hash()
	if err != nil {
		return "", err
	}

	return crypto.GetBigNumberStringFromBytes(crypto.GetFirstEightBytesReversed(hash)), nil
}

// Fee returns the calculated fee of the transaction
func (t *Transaction) Fee() (uint32, error) {

	switch t.Type {
	case TransactionTypeNormal:
		if t.Asset != nil {
			return feeSend + feeData, nil
		}
		return feeSend, nil
	case TransactionTypeSecondSecretRegistration:
		return feeSignature, nil
	case TransactionTypeDelegateRegistration:
		return feeDelegate, nil
	case TransactionTypeVote:
		return feeVote, nil
	case TransactionTypeMultisignatureRegistration:
		if asset, hasAsset := t.Asset.(*RegisterMultisignatureAccountAsset); hasAsset {
			return feeMultisignature * (1 + uint32(len(asset.AddKeys)+len(asset.RemoveKeys))), nil
		}

		return 0, errors.New("invalid asset - cannot calculate fee")
	case TransactionTypeDappRegistration:
		return feeDapp, nil
	case TransactionTypeTransferInSidechain:
		return feeTransferIn, nil
	case TransactionTypeTransferOutSidechain:
		return feeTransferOut, nil
	}

	return 0, errors.New("unknown transaction type")
}

func bytesSliceContainsDuplicates(data [][]byte) bool {
	for i, item := range data {
		for j, item2 := range data {
			if i != j && bytes.Equal(item, item2) {
				return true
			}
		}
	}
	return false
}
