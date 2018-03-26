package transactions

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"testing"

	"golang.org/x/crypto/ed25519"
)

type (
	DataAssetTest struct {
		in  DataAsset
		out bool
	}
	CastVoteAssetTest struct {
		in  *CastVoteAsset
		out bool
	}
	RegisterDelegateTest struct {
		in  *RegisterDelegateAsset
		out bool
	}
	RegisterSecondSignatureTest struct {
		in  *RegisterSecondSignatureAsset
		out bool
	}
	RegisterMultisignatureAccountTest struct {
		in  *RegisterMultisignatureAccountAsset
		out bool
	}
)

var (
	publicKey, _   = hex.DecodeString("5d036a858ce89f844491762eb89e2bfbd50a4a0a0da658e4b2628b25b117ae09")
	publicKey2, _  = hex.DecodeString("6d046a858ce89f844491762eb89e2bfbd50a4a0a0da658e4b2628b25b118ae09")
	DataAssetTests = []DataAssetTest{
		{DataAsset(""), true},
		{DataAsset("olf"), true},
		{DataAsset("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"), false},
	}
	CastVoteAssetTests = []CastVoteAssetTest{
		{&CastVoteAsset{}, true},
		{&CastVoteAsset{Votes: randomValidPublicKeys(20), Unvotes: randomValidPublicKeys(13)}, true},
		{&CastVoteAsset{Votes: randomValidPublicKeys(20), Unvotes: randomValidPublicKeys(15)}, false},
		{&CastVoteAsset{Votes: randomValidPublicKeys(33)}, true},
		{&CastVoteAsset{Votes: [][]byte{publicKey}}, true},
		{&CastVoteAsset{Votes: [][]byte{publicKey, publicKey2}}, true},
		{&CastVoteAsset{Votes: [][]byte{publicKey, publicKey, publicKey2}}, false},
		{&CastVoteAsset{Votes: [][]byte{publicKey, []byte("abc")}}, false},
		{&CastVoteAsset{Unvotes: [][]byte{publicKey}}, true},
		{&CastVoteAsset{Votes: [][]byte{publicKey}, Unvotes: [][]byte{publicKey}}, false},
		{&CastVoteAsset{Unvotes: [][]byte{publicKey, []byte("abc")}}, false},
		{&CastVoteAsset{Unvotes: [][]byte{publicKey, publicKey, []byte("abc")}}, false},
	}
	RegisterDelegateAssetTests = []RegisterDelegateTest{
		{in: &RegisterDelegateAsset{}, out: false},
		{in: &RegisterDelegateAsset{"", publicKey}, out: false},
		{in: &RegisterDelegateAsset{"olf", nil}, out: false},
		{in: &RegisterDelegateAsset{"aaaaaaaaaaaaaaaaaaaaa", publicKey}, out: false},
		{in: &RegisterDelegateAsset{"abc", publicKey}, out: true},
	}
	RegisterSecondSignatureTests = []RegisterSecondSignatureTest{
		{&RegisterSecondSignatureAsset{}, false},
		{&RegisterSecondSignatureAsset{PublicKey: []byte("abc")}, false},
		{&RegisterSecondSignatureAsset{PublicKey: publicKey}, true},
	}
	RegisterMultisignatureAccountTests = []RegisterMultisignatureAccountTest{
		{&RegisterMultisignatureAccountAsset{}, false},
		{&RegisterMultisignatureAccountAsset{AddKeys: [][]byte{publicKey, publicKey2}}, true},
		{&RegisterMultisignatureAccountAsset{RemoveKeys: [][]byte{publicKey, publicKey2}}, true},
		{&RegisterMultisignatureAccountAsset{AddKeys: [][]byte{publicKey}, RemoveKeys: [][]byte{publicKey2}}, true},
		{&RegisterMultisignatureAccountAsset{AddKeys: [][]byte{publicKey}, RemoveKeys: [][]byte{publicKey}}, false},
		{&RegisterMultisignatureAccountAsset{AddKeys: [][]byte{publicKey, publicKey2, publicKey}}, false},
		{&RegisterMultisignatureAccountAsset{RemoveKeys: [][]byte{publicKey, publicKey2, publicKey}}, false},
		{&RegisterMultisignatureAccountAsset{AddKeys: [][]byte{publicKey, []byte("abc")}}, false},
		{&RegisterMultisignatureAccountAsset{RemoveKeys: [][]byte{publicKey, []byte("abc")}}, false},
	}
)

func TestDataAsset_IsValid(t *testing.T) {
	for i, test := range DataAssetTests {
		val, err := test.in.IsValid()
		if val != test.out {
			t.Errorf("#%d: DataAsset.IsValid(%v)=%v,%v; want %v", i, test.in, val, err, test.out)
		}
	}
}

func TestCastVoteAsset_IsValid(t *testing.T) {
	for i, test := range CastVoteAssetTests {
		val, err := test.in.IsValid()
		if val != test.out {
			t.Errorf("#%d: CastVoteAsset.IsValid(%v)=%v,%v; want %v", i, test.in, val, err, test.out)
		}
	}
}

func TestRegisterDelegateAsset_IsValid(t *testing.T) {
	for i, test := range RegisterDelegateAssetTests {
		val, err := test.in.IsValid()
		if val != test.out {
			t.Errorf("#%d: RegisterDelegateAsset.IsValid(%v)=%v,%v; want %v", i, test.in, val, err, test.out)
		}
	}
}

func TestRegisterSecondSignatureAsset_IsValid(t *testing.T) {
	for i, test := range RegisterSecondSignatureTests {
		val, err := test.in.IsValid()
		if val != test.out {
			t.Errorf("#%d: RegisterSecondSignatureAsset.IsValid(%v)=%v,%v; want %v", i, test.in, val, err,
				test.out)
		}
	}
}

func TestRegisterMultisignatureAccountAsset_IsValid(t *testing.T) {
	for i, test := range RegisterMultisignatureAccountTests {
		val, err := test.in.IsValid()
		if val != test.out {
			t.Errorf("#%d: RegisterMultisignatureAccountAsset.IsValid(%v)=%v,%v; want %v", i, test.in, val, err,
				test.out)
		}
	}
}

func randomValidPublicKeys(num int) [][]byte {
	var keys [][]byte
	for i := 0; i < num; i++ {
		key := make([]byte, ed25519.PublicKeySize)
		rand.Read(key)
		keys = append(keys, key)
	}

	return keys
}

func TestDataAssetSerialize(t *testing.T) {
	asset := DataAsset("abc")
	if data, err := asset.serialize(); !bytes.Equal(data, []byte("abc")) || err != nil {
		t.Errorf("DataAsset.serialize()=%v,%v; want %v", data, err, []byte("abc"))
	}
}

func TestDataAsset_MarshalJSON(t *testing.T) {
	asset := DataAsset("abc")
	if data, err := asset.MarshalJSON(); string(data) != `{"data":"abc"}` || err != nil {
		t.Errorf("DataAsset.MarshalJSON()=%s,%v; want %v", data, err, "")
	}
}

func TestRegisterSecondSignatureAssetSerialize(t *testing.T) {
	asset := &RegisterSecondSignatureAsset{
		PublicKey: defaultSenderPublicKey,
	}
	if data, err := asset.serialize(); !bytes.Equal(data, defaultSenderPublicKey) || err != nil {
		t.Errorf("RegisterSecondSignatureAsset.serialize()=%v,%v; want %v", data, err, defaultSenderPublicKey)
	}
}

func TestRegisterDelegateAssetSerialize(t *testing.T) {
	asset := &RegisterDelegateAsset{
		PublicKey: defaultSenderPublicKey,
		Username:  "abc",
	}
	if data, err := asset.serialize(); !bytes.Equal(data, []byte("abc")) || err != nil {
		t.Errorf("RegisterDelegateAsset.serialize()=%v,%v; want %v", data, err, []byte("abc"))
	}
}

func TestCastVoteAssetSerialize(t *testing.T) {
	asset := &CastVoteAsset{
		Votes:   [][]byte{defaultSenderPublicKey},
		Unvotes: [][]byte{defaultSenderSecondPublicKey},
	}
	if data, err := asset.serialize(); !bytes.Equal(data, []byte("+5d036a858ce89f844491762eb89e2bfbd50a4a0a0da658e4b2628b25b117ae09-0401c8ac9f29ded9e1e4d5b6b43051cb25b22f27c7b7b35092161e851946f82f")) || err != nil {
		t.Errorf("CastVoteAsset.serialize()=%v,%v; want %v", data, err, []byte("abc"))
	}
}

func TestRegisterMultisignatureAccountAssetSerialize(t *testing.T) {
	asset := &RegisterMultisignatureAccountAsset{
		Min:        48,
		Lifetime:   49,
		AddKeys:    [][]byte{defaultSenderPublicKey},
		RemoveKeys: [][]byte{defaultSenderSecondPublicKey},
	}
	if data, err := asset.serialize(); !bytes.Equal(data, []byte("01+5d036a858ce89f844491762eb89e2bfbd50a4a0a0da658e4b2628b25b117ae09-0401c8ac9f29ded9e1e4d5b6b43051cb25b22f27c7b7b35092161e851946f82f")) || err != nil {
		t.Errorf("RegisterMultisignatureAccountAsset.serialize()=%v,%v; want %v", data, err, []byte("abc"))
	}
}

func TestRegisterSecondSignatureAsset_MarshalJSON(t *testing.T) {
	asset := &RegisterSecondSignatureAsset{
		PublicKey: defaultSenderPublicKey,
	}
	if data, err := asset.MarshalJSON(); string(data) != `{"signature":{"publicKey":"5d036a858ce89f844491762eb89e2bfbd50a4a0a0da658e4b2628b25b117ae09"}}` || err != nil {
		t.Errorf("RegisterSecondSignatureAsset.MarshalJSON()=%s,%v; want %v", data, err, "")
	}
}

func TestRegisterDelegateAsset_MarshalJSON(t *testing.T) {
	asset := &RegisterDelegateAsset{
		PublicKey: defaultSenderPublicKey,
		Username:  "abc",
	}
	if data, err := asset.MarshalJSON(); string(data) != `{"delegate":{"username":"abc","publicKey":"5d036a858ce89f844491762eb89e2bfbd50a4a0a0da658e4b2628b25b117ae09"}}` || err != nil {
		t.Errorf("RegisterDelegateAsset.MarshalJSON()=%s,%v; want %v", data, err, "")
	}
}

func TestCastVoteAsset_MarshalJSON(t *testing.T) {
	asset := &CastVoteAsset{
		Votes:   [][]byte{defaultSenderPublicKey},
		Unvotes: [][]byte{defaultSenderSecondPublicKey},
	}
	if data, err := asset.MarshalJSON(); string(data) != `{"votes":["+5d036a858ce89f844491762eb89e2bfbd50a4a0a0da658e4b2628b25b117ae09","-0401c8ac9f29ded9e1e4d5b6b43051cb25b22f27c7b7b35092161e851946f82f"]}` || err != nil {
		t.Errorf("RegisterDelegateAsset.MarshalJSON()=%s,%v; want %v", data, err, "")
	}
}

func TestRegisterMultisignatureAccountAsset_MarshalJSON(t *testing.T) {
	asset := &RegisterMultisignatureAccountAsset{
		Min:        48,
		Lifetime:   49,
		AddKeys:    [][]byte{defaultSenderPublicKey},
		RemoveKeys: [][]byte{defaultSenderSecondPublicKey},
	}
	if data, err := asset.MarshalJSON(); string(data) != `{"multisignature":{"keysgroup":["+5d036a858ce89f844491762eb89e2bfbd50a4a0a0da658e4b2628b25b117ae09","-0401c8ac9f29ded9e1e4d5b6b43051cb25b22f27c7b7b35092161e851946f82f"],"min":48,"lifetime":49}}` || err != nil {
		t.Errorf("RegisterDelegateAsset.MarshalJSON()=%s,%v; want %v", data, err, "")
	}
}
