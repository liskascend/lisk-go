package transactions

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
)

var (
	defaultPrivateKey, _ = hex.DecodeString("314852d7afb0d4c283692fef8a2cb40e30c7a5df2ed79994178c10ac168d6d977ef45cd525e95b7a86244bbd4eb4550914ad06301013958f4dd64d32ef7bc588")
)

func TestTransaction_Sign(t *testing.T) {
	transaction := Transaction{SenderPublicKey: defaultSenderPublicKey}
	err := transaction.Sign(defaultPrivateKey)

	if base64.StdEncoding.EncodeToString(transaction.signature) != "dPFiRMXaoW8zAooOemp6sGv9BR6obHnjHOdWmg28n5QzJzs85+sNIvJpLzOxOq3NnCFqbnEChsdzCgMyandbCQ==" || err != nil {
		t.Errorf("Transaction.Sign() generates wrong signature: %v; error: %v", base64.StdEncoding.EncodeToString(transaction.signature), err)
	}
}

func TestTransaction_SecondSign(t *testing.T) {
	transaction := Transaction{SenderPublicKey: defaultSenderPublicKey}
	err := transaction.SecondSign(defaultPrivateKey)

	if base64.StdEncoding.EncodeToString(transaction.secondSignature) != "dPFiRMXaoW8zAooOemp6sGv9BR6obHnjHOdWmg28n5QzJzs85+sNIvJpLzOxOq3NnCFqbnEChsdzCgMyandbCQ==" || err != nil {
		t.Errorf("Transaction.SecondSign() generates wrong signature: %v; error: %v", base64.StdEncoding.EncodeToString(transaction.signature), err)
	}
}
