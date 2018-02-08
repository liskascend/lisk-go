package transactions

import (
	"encoding/base64"
	"testing"
)

var (
	bytesWithDuplicates    = [][]byte{{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 2, 3}}
	bytesWithoutDuplicates = [][]byte{{1, 2, 3}, {4, 2, 3}, {1, 4, 3}}
)

func TestByteSliceContainsDuplicates(t *testing.T) {
	if val := bytesSliceContainsDuplicates(bytesWithDuplicates); !val {
		t.Errorf("bytesSliceContainsDuplicates(%v)=%v; want %v", bytesWithDuplicates, val, true)
	}

	if val := bytesSliceContainsDuplicates(bytesWithoutDuplicates); val {
		t.Errorf("bytesSliceContainsDuplicates(%v)=%v; want %v", bytesWithoutDuplicates, val, false)
	}
}

func TestSerializeTransactionHash(t *testing.T) {
	transaction := &Transaction{
		Type:            0,
		Amount:          uint64(defaultAmount),
		RecipientID:     defaultRecipient,
		Timestamp:       uint32(defaultTimestamp),
		SenderPublicKey: defaultSenderPublicKey,
		signature:       defaultSignature,
	}

	if val, err := transaction.Hash(); base64.StdEncoding.EncodeToString(val) != "9gom2kcLHcIz/VJu1zBsHYSDb54uzugsnsRzGeCRBHQ=" || err != nil {
		t.Errorf("Transaction.Serialize() returns wrong data: %v; error: %v", val, err)
	}
}

func TestSerializeTransactionID(t *testing.T) {
	transaction := &Transaction{
		Type:            0,
		Amount:          uint64(defaultAmount),
		RecipientID:     defaultRecipient,
		Timestamp:       uint32(defaultTimestamp),
		SenderPublicKey: defaultSenderPublicKey,
		signature:       defaultSignature,
	}

	if val, err := transaction.ID(); val != defaultTransactionId || err != nil {
		t.Errorf("Transaction.ID() returns wrong data: %v; error: %v", val, err)
	}
}

func TestTransaction_Fee(t *testing.T) {
	if val, err := (&Transaction{
		Type: 0,
	}).Fee(); val != 0.1*fixedPoint || err != nil {
		t.Errorf("Transaction.Fee() returns wrong data: %v,%v; want %v", val, err, 0.1*fixedPoint)
	}

	if val, err := (&Transaction{
		Type:  0,
		Asset: DataAsset("test"),
	}).Fee(); val != 0.2*fixedPoint || err != nil {
		t.Errorf("Transaction.Fee() returns wrong data: %v,%v; want %v", val, err, 0.2*fixedPoint)
	}

	if val, err := (&Transaction{
		Type: 1,
	}).Fee(); val != 5*fixedPoint || err != nil {
		t.Errorf("Transaction.Fee() returns wrong data: %v,%v; want %v", val, err, 5*fixedPoint)
	}

	if val, err := (&Transaction{
		Type: 2,
	}).Fee(); val != 25*fixedPoint || err != nil {
		t.Errorf("Transaction.Fee() returns wrong data: %v,%v; want %v", val, err, 25*fixedPoint)
	}

	if val, err := (&Transaction{
		Type: 3,
	}).Fee(); val != 1*fixedPoint || err != nil {
		t.Errorf("Transaction.Fee() returns wrong data: %v,%v; want %v", val, err, 1*fixedPoint)
	}

	if val, err := (&Transaction{
		Type: 4,
		Asset: &RegisterMultisignatureAccountAsset{
			AddKeys: [][]byte{{}},
		},
	}).Fee(); val != 10*fixedPoint || err != nil {
		t.Errorf("Transaction.Fee() returns wrong data: %v,%v; want %v", val, err, 10*fixedPoint)
	}

	if val, err := (&Transaction{
		Type: 4,
	}).Fee(); err == nil {
		t.Errorf("Transaction.Fee() returns wrong data: %v,%v; expected error", val, err)
	}

	if val, err := (&Transaction{
		Type: 5,
	}).Fee(); val != 25*fixedPoint || err != nil {
		t.Errorf("Transaction.Fee() returns wrong data: %v,%v; want %v", val, err, 25*fixedPoint)
	}

	if val, err := (&Transaction{
		Type: 6,
	}).Fee(); val != 0.1*fixedPoint || err != nil {
		t.Errorf("Transaction.Fee() returns wrong data: %v,%v; want %v", val, err, 0.1*fixedPoint)
	}

	if val, err := (&Transaction{
		Type: 7,
	}).Fee(); val != 0.1*fixedPoint || err != nil {
		t.Errorf("Transaction.Fee() returns wrong data: %v,%v; want %v", val, err, 0.1*fixedPoint)
	}

	if val, err := (&Transaction{
		Type: 8,
	}).Fee(); err == nil {
		t.Errorf("Transaction.Fee() returns wrong data: %v,%v; expected error", val, err)
	}
}
