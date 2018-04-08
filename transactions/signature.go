package transactions

import (
	"github.com/liskascend/lisk-go/crypto"
)

// Sign signs the transaction with the given privateKey.
// This has to be redone when any fields of the transaction are changed.
func (t *Transaction) Sign(privateKey []byte) error {
	hash, err := t.Hash()
	if err != nil {
		return err
	}

	t.signature = crypto.SignMessageWithPrivateKey(string(hash), privateKey)

	return nil
}

// SecondSign adds a second signature to the transaction using the given privateKey.
// This has to be redone when any fields of the transaction are changed.
func (t *Transaction) SecondSign(privateKey []byte) error {
	hash, err := t.Hash()
	if err != nil {
		return err
	}

	t.secondSignature = crypto.SignMessageWithPrivateKey(string(hash), privateKey)

	return nil
}
