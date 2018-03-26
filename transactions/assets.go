package transactions

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"errors"

	"golang.org/x/crypto/ed25519"
)

type (
	// emptyAsset is an asset placeholder required for json marshaling
	emptyAsset struct{}

	// DataAsset is an asset that is used to attach data to a normal transaction
	DataAsset string

	// RegisterSecondSignatureAsset is the asset needed for a TransactionTypeSecondSecretRegistration
	RegisterSecondSignatureAsset struct {
		// PublicKey is the public key for the second secret
		PublicKey []byte
	}

	// RegisterDelegateAsset is the asset needed for a TransactionTypeDelegateRegistration
	RegisterDelegateAsset struct {
		// Username is the username for the delegate
		Username string
		// PublicKey is the public key of the delegate
		PublicKey []byte
	}

	// CastVoteAsset is the asset needed for a TransactionTypeVote
	// The maximum number of votes is 33
	CastVoteAsset struct {
		// Votes is a slice of the public keys of the delegates to vote for
		Votes [][]byte
		// Unvotes is a slice of the public keys of the delegates to remove votes from
		Unvotes [][]byte
	}

	// RegisterMultisignatureAccountAsset is the asset needed for a TransactionTypeMultisignatureRegistration
	RegisterMultisignatureAccountAsset struct {
		// Min is the minimum number of signatures that is required for a transaction
		Min byte `json:"min"`
		// Lifetime is the lifetime of the transaction in which it can be signed
		Lifetime byte `json:"lifetime"`
		// AddKeys is a slice of keys to add to the multisignature wallet
		AddKeys [][]byte `json:"_"`
		// AddKeys is a slice of keys to remove from the multisignature wallet
		RemoveKeys [][]byte `json:"_"`
	}

	// CreateDappAsset is the asset needed for a TransactionTypeDappRegistration
	CreateDappAsset struct {
		// Dapp is the dapp that should be registered
		Dapp *Dapp
	}

	// TransferInDappAsset is the asset needed for a TransactionTypeTransferInSidechain
	TransferInDappAsset struct {
		// DappID is the ID of the DApp to transfer lisk in
		DappID string `json:"dappId"`
	}

	// TransferOutDappAsset is the asset needed for a TransactionTypeTransferOutSidechain
	TransferOutDappAsset struct {
		// DappID is the ID of the DApp to transfer lisk out of
		DappID string `json:"dappId"`
		// TransactionID is the ID of the withdrawal transaction on the sidechain
		TransactionID string `json:"transactionId"`
	}
)

func (r *RegisterMultisignatureAccountAsset) serialize() ([]byte, error) {
	if valid, err := r.IsValid(); !valid {
		return nil, err
	}

	var dst []byte
	dst = append(dst, r.Min)
	dst = append(dst, r.Lifetime)

	for _, signature := range r.AddKeys {
		sig := "+" + hex.EncodeToString(signature)
		dst = append(dst, sig...)
	}
	for _, signature := range r.RemoveKeys {
		sig := "-" + hex.EncodeToString(signature)
		dst = append(dst, sig...)
	}

	return dst, nil
}

// IsValid returns whether the asset is valid
func (r *RegisterMultisignatureAccountAsset) IsValid() (bool, error) {
	if len(r.AddKeys)+len(r.RemoveKeys) == 0 {
		return false, errors.New("no keys specified to remove or add")
	}

	if bytesSliceContainsDuplicates(append(r.AddKeys, r.RemoveKeys...)) {
		return false, errors.New("addKeys and/or removeKeys contain duplicates")
	}

	for _, key := range append(r.AddKeys, r.RemoveKeys...) {
		if len(key) != ed25519.PublicKeySize {
			return false, fmt.Errorf(
				"key %s has invalid length %d",
				hex.EncodeToString(key),
				len(key))
		}
	}

	return true, nil
}

// MarshalJSON marshals the asset to the lisk JSON format
func (r *RegisterMultisignatureAccountAsset) MarshalJSON() ([]byte, error) {
	var keysgroup []string
	for _, signature := range r.AddKeys {
		key := "+" + hex.EncodeToString(signature)
		keysgroup = append(keysgroup, key)
	}
	for _, signature := range r.RemoveKeys {
		key := "-" + hex.EncodeToString(signature)
		keysgroup = append(keysgroup, key)
	}

	type AssetWrapper RegisterMultisignatureAccountAsset
	return json.Marshal(&struct {
		Payload interface{} `json:"multisignature"`
	}{
		Payload: struct {
			Keysgroup []string `json:"keysgroup"`
			*AssetWrapper
		}{
			Keysgroup:    keysgroup,
			AssetWrapper: (*AssetWrapper)(r),
		},
	})
}

func (c *CastVoteAsset) serialize() ([]byte, error) {
	if valid, err := c.IsValid(); !valid {
		return nil, err
	}

	var dst []byte
	for _, vote := range c.Votes {
		v := "+" + hex.EncodeToString(vote)
		dst = append(dst, v...)
	}
	for _, unvote := range c.Unvotes {
		u := "-" + hex.EncodeToString(unvote)
		dst = append(dst, u...)
	}
	return dst, nil
}

// IsValid returns whether the asset is valid
func (c *CastVoteAsset) IsValid() (bool, error) {
	if len(c.Votes)+len(c.Unvotes) > 33 {
		return false, errors.New("too many votes/unvotes")
	}

	for _, vote := range append(c.Votes, c.Unvotes...) {
		if len(vote) != ed25519.PublicKeySize {
			return false, fmt.Errorf(
				"vote %s has invalid length %d",
				hex.EncodeToString(vote),
				len(vote))
		}
	}

	if bytesSliceContainsDuplicates(append(c.Votes, c.Unvotes...)) {
		return false, errors.New("votes and/or votes contain duplicates")
	}

	return true, nil
}

// MarshalJSON marshals the asset to the lisk JSON format
func (c *CastVoteAsset) MarshalJSON() ([]byte, error) {
	var votes []string

	for _, vote := range c.Votes {
		key := "+" + hex.EncodeToString(vote)
		votes = append(votes, key)
	}
	for _, unvote := range c.Unvotes {
		key := "-" + hex.EncodeToString(unvote)
		votes = append(votes, key)
	}

	return json.Marshal(&struct {
		Votes []string `json:"votes"`
	}{
		Votes: votes,
	})
}

func (r *RegisterDelegateAsset) serialize() ([]byte, error) {
	if valid, err := r.IsValid(); !valid {
		return nil, err
	}
	return []byte(r.Username), nil
}

// IsValid returns whether the asset is valid
func (r *RegisterDelegateAsset) IsValid() (bool, error) {
	if len(r.Username) == 0 {
		return false, errors.New("username must not be empty")
	}

	if len(r.Username) > 20 {
		return false, errors.New("username exceeds the maximum length of 20 characters")
	}

	if len(r.PublicKey) != ed25519.PublicKeySize {
		return false, errors.New("invalid public key length")
	}

	return true, nil
}

// MarshalJSON marshals the asset to the lisk JSON format
func (r *RegisterDelegateAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Payload interface{} `json:"delegate"`
	}{
		Payload: struct {
			Username  string `json:"username"`
			PublicKey string `json:"publicKey"`
		}{
			Username:  r.Username,
			PublicKey: hex.EncodeToString(r.PublicKey),
		},
	})
}

func (r *RegisterSecondSignatureAsset) serialize() ([]byte, error) {
	if valid, err := r.IsValid(); !valid {
		return nil, err
	}
	return r.PublicKey, nil
}

// IsValid returns whether the asset is valid
func (r *RegisterSecondSignatureAsset) IsValid() (bool, error) {
	if len(r.PublicKey) != ed25519.PublicKeySize {
		return false,
			fmt.Errorf(
				"public key %s has invalid length %d",
				hex.EncodeToString(r.PublicKey),
				len(r.PublicKey))
	}
	return true, nil
}

// MarshalJSON marshals the asset to the lisk JSON format
func (r *RegisterSecondSignatureAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Payload interface{} `json:"signature"`
	}{
		Payload: struct {
			PublicKey string `json:"publicKey"`
		}{
			PublicKey: hex.EncodeToString(r.PublicKey),
		},
	})
}

func (a DataAsset) serialize() ([]byte, error) {
	if valid, err := a.IsValid(); !valid {
		return nil, err
	}
	return []byte(a), nil
}

// IsValid returns whether the asset is valid
func (a DataAsset) IsValid() (bool, error) {
	if len([]byte(a)) > byteSizeData {
		return false, errors.New("data length exceeds maximum payload size")
	}
	return true, nil
}

// MarshalJSON marshals the asset to the lisk JSON format
func (a DataAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Data string `json:"data"`
	}{
		Data: string(a),
	})
}
