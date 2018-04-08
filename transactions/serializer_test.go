package transactions

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
)

var (
	defaultRecipient                = "58191285901858109L"
	defaultSenderPublicKey, _       = hex.DecodeString("5d036a858ce89f844491762eb89e2bfbd50a4a0a0da658e4b2628b25b117ae09")
	defaultSenderId                 = "18160565574430594874L"
	defaultSenderSecondPublicKey, _ = hex.DecodeString("0401c8ac9f29ded9e1e4d5b6b43051cb25b22f27c7b7b35092161e851946f82f")
	defaultAmount                   = 1000
	defaultNoAmount                 = 0
	defaultTimestamp                = 141738
	defaultTransactionId            = "13987348420913138422"
	defaultSignature, _             = hex.DecodeString("618a54975212ead93df8c881655c625544bce8ed7ccdfe6f08a42eecfb1adebd051307be5014bb051617baf7815d50f62129e70918190361e5d4dd4796541b0a")
	defaultSecondSignature, _       = hex.DecodeString("b00c4ad1988bca245d74435660a278bfe6bf2f5efa8bda96d927fabf8b4f6fcfdcb2953f6abacaa119d6880987a55dea0e6354bc8366052b45fa23145522020f")
	defaultAppId                    = "1234213"
	defaultDelegateUsername         = "MyDelegateUsername"
	defaultRequesterPublicKey, _    = hex.DecodeString("5d036a858ce89f844491762eb89e2bfbd50a4a0a0da658e4b2628b25b117ae09")
)

func TestSerializeTransactionType0(t *testing.T) {
	transaction := &Transaction{
		Type:            0,
		Amount:          uint64(defaultAmount),
		RecipientID:     defaultRecipient,
		Timestamp:       uint32(defaultTimestamp),
		SenderPublicKey: defaultSenderPublicKey,
		signature:       defaultSignature,
	}

	if val, err := transaction.Serialize(); base64.StdEncoding.EncodeToString(val) != "AKopAgBdA2qFjOifhESRdi64niv71QpKCg2mWOSyYoslsReuCQDOvKqNNBU96AMAAAAAAABhilSXUhLq2T34yIFlXGJVRLzo7XzN/m8IpC7s+xrevQUTB75QFLsFFhe694FdUPYhKecJGBkDYeXU3UeWVBsK" || err != nil {
		t.Errorf("Transaction.Serialize() returns wrong data: %v; error: %v", val, err)
	}
}

func TestSerializeTransactionType0WithData(t *testing.T) {
	transaction := &Transaction{
		Type:            0,
		Amount:          uint64(defaultAmount),
		RecipientID:     defaultRecipient,
		Timestamp:       uint32(defaultTimestamp),
		SenderPublicKey: defaultSenderPublicKey,
		signature:       defaultSignature,
		Asset:           DataAsset("Hello Lisk! Some data in here!..."),
	}

	if val, err := transaction.Serialize(); base64.StdEncoding.EncodeToString(val) != "AKopAgBdA2qFjOifhESRdi64niv71QpKCg2mWOSyYoslsReuCQDOvKqNNBU96AMAAAAAAABIZWxsbyBMaXNrISBTb21lIGRhdGEgaW4gaGVyZSEuLi5hilSXUhLq2T34yIFlXGJVRLzo7XzN/m8IpC7s+xrevQUTB75QFLsFFhe694FdUPYhKecJGBkDYeXU3UeWVBsK" || err != nil {
		t.Errorf("Transaction.Serialize() returns wrong data: %v; error: %v", val, err)
	}
}

func TestSerializeTransactionType0WithInvalidData(t *testing.T) {
	transaction := &Transaction{
		Type:            0,
		Amount:          uint64(defaultAmount),
		RecipientID:     defaultRecipient,
		Timestamp:       uint32(defaultTimestamp),
		SenderPublicKey: defaultSenderPublicKey,
		signature:       defaultSignature,
		Asset:           DataAsset("Hello Lisk! Some data in here!...aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	}

	if val, err := transaction.Serialize(); err == nil {
		t.Errorf("Transaction.Serialize() returns wrong data: %v, nil; expected error", val)
	}
}

func TestSerializeTransactionType0WithInvalidSignature(t *testing.T) {
	transaction := &Transaction{
		Type:            0,
		Amount:          uint64(defaultAmount),
		RecipientID:     defaultRecipient,
		Timestamp:       uint32(defaultTimestamp),
		SenderPublicKey: defaultSenderPublicKey,
		signature:       defaultSignature[:32],
	}

	if val, err := transaction.Serialize(); err == nil {
		t.Errorf("Transaction.Serialize() returns wrong data: %v, nil; expected error", val)
	}
}

func TestSerializeTransactionType0WithSecondSignature(t *testing.T) {
	transaction := &Transaction{
		Type:            0,
		Amount:          uint64(defaultAmount),
		RecipientID:     defaultRecipient,
		Timestamp:       uint32(defaultTimestamp),
		SenderPublicKey: defaultSenderPublicKey,
		signature:       defaultSignature,
		secondSignature: defaultSecondSignature,
	}

	if val, err := transaction.Serialize(); base64.StdEncoding.EncodeToString(val) != "AKopAgBdA2qFjOifhESRdi64niv71QpKCg2mWOSyYoslsReuCQDOvKqNNBU96AMAAAAAAABhilSXUhLq2T34yIFlXGJVRLzo7XzN/m8IpC7s+xrevQUTB75QFLsFFhe694FdUPYhKecJGBkDYeXU3UeWVBsKsAxK0ZiLyiRddENWYKJ4v+a/L176i9qW2Sf6v4tPb8/cspU/arrKoRnWiAmHpV3qDmNUvINmBStF+iMUVSICDw==" || err != nil {
		t.Errorf("Transaction.Serialize() returns wrong data: %v; error: %v", val, err)
	}
}

func TestSerializeTransactionType0WithInvalidSecondSignature(t *testing.T) {
	transaction := &Transaction{
		Type:            0,
		Amount:          uint64(defaultAmount),
		RecipientID:     defaultRecipient,
		Timestamp:       uint32(defaultTimestamp),
		SenderPublicKey: defaultSenderPublicKey,
		secondSignature: defaultSecondSignature[:32],
	}

	if val, err := transaction.Serialize(); err == nil {
		t.Errorf("Transaction.Serialize() returns wrong data: %v, nil; expected error", val)
	}
}

func TestSerializeTransactionType0WithMultiSig(t *testing.T) {
	transaction := &Transaction{
		Type:                          0,
		Amount:                        uint64(1000),
		RecipientID:                   defaultRecipient,
		Timestamp:                     uint32(defaultTimestamp),
		SenderPublicKey:               defaultSenderPublicKey,
		signature:                     defaultSignature,
		TransactionRequesterPublicKey: defaultRequesterPublicKey,
	}

	if val, err := transaction.Serialize(); base64.StdEncoding.EncodeToString(val) != "AKopAgBdA2qFjOifhESRdi64niv71QpKCg2mWOSyYoslsReuCV0DaoWM6J+ERJF2LrieK/vVCkoKDaZY5LJiiyWxF64JAM68qo00FT3oAwAAAAAAAGGKVJdSEurZPfjIgWVcYlVEvOjtfM3+bwikLuz7Gt69BRMHvlAUuwUWF7r3gV1Q9iEp5wkYGQNh5dTdR5ZUGwo=" || err != nil {
		t.Errorf("Transaction.Serialize() returns wrong data: %v; error: %v", val, err)
	}
}

func TestSerializeTransactionType0WithInvalidMultiSig(t *testing.T) {
	transaction := &Transaction{
		Type:                          0,
		Amount:                        uint64(1000),
		RecipientID:                   defaultRecipient,
		Timestamp:                     uint32(defaultTimestamp),
		SenderPublicKey:               defaultSenderPublicKey,
		signature:                     defaultSignature,
		TransactionRequesterPublicKey: defaultRequesterPublicKey[:31],
	}

	if val, err := transaction.Serialize(); err == nil {
		t.Errorf("Transaction.Serialize() returns wrong data: %v, nil; expected error", val)
	}
}

func TestSerializeTransactionType1(t *testing.T) {
	transaction := &Transaction{
		Type:            1,
		Amount:          uint64(defaultNoAmount),
		Timestamp:       uint32(defaultTimestamp),
		SenderPublicKey: defaultSenderPublicKey,
		signature:       defaultSignature,
		Asset:           &RegisterSecondSignatureAsset{PublicKey: defaultSenderSecondPublicKey},
	}

	if val, err := transaction.Serialize(); base64.StdEncoding.EncodeToString(val) != "AaopAgBdA2qFjOifhESRdi64niv71QpKCg2mWOSyYoslsReuCQAAAAAAAAAAAAAAAAAAAAAEAcisnyne2eHk1ba0MFHLJbIvJ8e3s1CSFh6FGUb4L2GKVJdSEurZPfjIgWVcYlVEvOjtfM3+bwikLuz7Gt69BRMHvlAUuwUWF7r3gV1Q9iEp5wkYGQNh5dTdR5ZUGwo=" || err != nil {
		t.Errorf("Transaction.Serialize() returns wrong data: %v; error: %v", val, err)
	}
}

func TestSerializeTransactionType2(t *testing.T) {
	transaction := &Transaction{
		Type:            2,
		Amount:          uint64(defaultNoAmount),
		Timestamp:       uint32(defaultTimestamp),
		SenderPublicKey: defaultSenderPublicKey,
		signature:       defaultSignature,
		Asset:           &RegisterDelegateAsset{Username: defaultDelegateUsername, PublicKey: defaultSenderPublicKey},
	}

	if val, err := transaction.Serialize(); base64.StdEncoding.EncodeToString(val) != "AqopAgBdA2qFjOifhESRdi64niv71QpKCg2mWOSyYoslsReuCQAAAAAAAAAAAAAAAAAAAABNeURlbGVnYXRlVXNlcm5hbWVhilSXUhLq2T34yIFlXGJVRLzo7XzN/m8IpC7s+xrevQUTB75QFLsFFhe694FdUPYhKecJGBkDYeXU3UeWVBsK" || err != nil {
		t.Errorf("Transaction.Serialize() returns wrong data: %v; error: %v", val, err)
	}
}

func TestSerializeTransactionType3(t *testing.T) {
	transaction := &Transaction{
		Type:            3,
		RecipientID:     defaultRecipient,
		Amount:          uint64(defaultNoAmount),
		Timestamp:       uint32(defaultTimestamp),
		SenderPublicKey: defaultSenderPublicKey,
		signature:       defaultSignature,
		Asset: &CastVoteAsset{
			Votes: [][]byte{defaultSenderPublicKey, defaultSenderSecondPublicKey},
		},
	}

	if val, err := transaction.Serialize(); base64.StdEncoding.EncodeToString(val) != "A6opAgBdA2qFjOifhESRdi64niv71QpKCg2mWOSyYoslsReuCQDOvKqNNBU9AAAAAAAAAAArNWQwMzZhODU4Y2U4OWY4NDQ0OTE3NjJlYjg5ZTJiZmJkNTBhNGEwYTBkYTY1OGU0YjI2MjhiMjViMTE3YWUwOSswNDAxYzhhYzlmMjlkZWQ5ZTFlNGQ1YjZiNDMwNTFjYjI1YjIyZjI3YzdiN2IzNTA5MjE2MWU4NTE5NDZmODJmYYpUl1IS6tk9+MiBZVxiVUS86O18zf5vCKQu7Psa3r0FEwe+UBS7BRYXuveBXVD2ISnnCRgZA2Hl1N1HllQbCg==" || err != nil {
		t.Errorf("Transaction.Serialize() returns wrong data: %v; error: %v", val, err)
	}
}

func TestSerializeTransactionType4(t *testing.T) {
	transaction := &Transaction{
		Type:            4,
		Amount:          uint64(defaultNoAmount),
		Timestamp:       uint32(defaultTimestamp),
		SenderPublicKey: defaultSenderPublicKey,
		signature:       defaultSignature,
		Asset: &RegisterMultisignatureAccountAsset{
			Min:      2,
			Lifetime: 5,
			AddKeys:  [][]byte{defaultSenderPublicKey, defaultSenderSecondPublicKey},
		},
	}

	if val, err := transaction.Serialize(); base64.StdEncoding.EncodeToString(val) != "BKopAgBdA2qFjOifhESRdi64niv71QpKCg2mWOSyYoslsReuCQAAAAAAAAAAAAAAAAAAAAACBSs1ZDAzNmE4NThjZTg5Zjg0NDQ5MTc2MmViODllMmJmYmQ1MGE0YTBhMGRhNjU4ZTRiMjYyOGIyNWIxMTdhZTA5KzA0MDFjOGFjOWYyOWRlZDllMWU0ZDViNmI0MzA1MWNiMjViMjJmMjdjN2I3YjM1MDkyMTYxZTg1MTk0NmY4MmZhilSXUhLq2T34yIFlXGJVRLzo7XzN/m8IpC7s+xrevQUTB75QFLsFFhe694FdUPYhKecJGBkDYeXU3UeWVBsK" || err != nil {
		t.Errorf("Transaction.Serialize() returns wrong data: %v; error: %v", val, err)
	}
}

func TestMarshalTransaction(t *testing.T) {
	transaction := &Transaction{
		Type:            0,
		Amount:          uint64(defaultAmount),
		RecipientID:     defaultRecipient,
		Timestamp:       uint32(defaultTimestamp),
		SenderPublicKey: defaultSenderPublicKey,
		signature:       defaultSignature,
	}

	if val, err := transaction.MarshalJSON(); string(val) != "{\"type\":0,\"id\":\"13987348420913138422\",\"senderId\":\"18160565574430594874L\",\"amount\":\"1000\",\"fee\":\"10000000\",\"recipientId\":\"58191285901858109L\",\"timestamp\":141738,\"asset\":{},\"senderPublicKey\":\"5d036a858ce89f844491762eb89e2bfbd50a4a0a0da658e4b2628b25b117ae09\",\"signature\":\"618a54975212ead93df8c881655c625544bce8ed7ccdfe6f08a42eecfb1adebd051307be5014bb051617baf7815d50f62129e70918190361e5d4dd4796541b0a\"}" || err != nil {
		t.Errorf("Transaction.MarshalJSON() returns wrong data: %s; error: %v", val, err)
	}
}

func TestMarshalTransactionInvalid(t *testing.T) {
	transaction := &Transaction{
		Type:        0,
		Amount:      uint64(defaultAmount),
		RecipientID: defaultRecipient,
		Timestamp:   uint32(defaultTimestamp),
		signature:   defaultSignature,
	}

	if val, err := transaction.MarshalJSON(); err == nil {
		t.Errorf("Transaction.MarshalJSON() returns wrong data: %s, nil; expected error", val)
	}
}

func TestTransactionInvalidType(t *testing.T) {
	transaction := &Transaction{
		Type:            10,
		Amount:          uint64(defaultAmount),
		RecipientID:     defaultRecipient,
		SenderPublicKey: defaultSenderPublicKey,
		Timestamp:       uint32(defaultTimestamp),
		signature:       defaultSignature,
	}

	if val, err := transaction.IsValid(); err == nil {
		t.Errorf("Transaction.IsValid() returns wrong data: %v, nil; expected error", val)
	}
}

func TestTransactionMissingAsset(t *testing.T) {
	transaction := &Transaction{
		Type:            TransactionTypeSecondSecretRegistration,
		Amount:          uint64(defaultAmount),
		RecipientID:     defaultRecipient,
		SenderPublicKey: defaultSenderPublicKey,
		Timestamp:       uint32(defaultTimestamp),
		signature:       defaultSignature,
	}

	if val, err := transaction.IsValid(); err == nil {
		t.Errorf("Transaction.IsValid() returns wrong data: %v, nil; expected error", val)
	}
}

func TestTransactionInvalidAssetType(t *testing.T) {
	transaction := &Transaction{
		Type:            TransactionTypeNormal,
		Amount:          uint64(defaultAmount),
		RecipientID:     defaultRecipient,
		SenderPublicKey: defaultSenderPublicKey,
		Timestamp:       uint32(defaultTimestamp),
		signature:       defaultSignature,
		Asset:           &CastVoteAsset{},
	}

	if val, err := transaction.IsValid(); err == nil {
		t.Errorf("Transaction.IsValid() returns wrong data: %v, nil; expected error", val)
	}

	transaction.Type = TransactionTypeSecondSecretRegistration
	transaction.Asset = DataAsset("abc")
	if val, err := transaction.IsValid(); err == nil {
		t.Errorf("Transaction.IsValid() returns wrong data: %v, nil; expected error", val)
	}

	transaction.Type = TransactionTypeDelegateRegistration
	if val, err := transaction.IsValid(); err == nil {
		t.Errorf("Transaction.IsValid() returns wrong data: %v, nil; expected error", val)
	}

	transaction.Type = TransactionTypeVote
	if val, err := transaction.IsValid(); err == nil {
		t.Errorf("Transaction.IsValid() returns wrong data: %v, nil; expected error", val)
	}

	transaction.Type = TransactionTypeMultisignatureRegistration
	if val, err := transaction.IsValid(); err == nil {
		t.Errorf("Transaction.IsValid() returns wrong data: %v, nil; expected error", val)
	}
}
