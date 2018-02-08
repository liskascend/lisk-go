package crypto

import (
	"encoding/hex"
	"testing"
)

var (
	defaultMessage    = "Some default text."
	signPublicKey, _  = hex.DecodeString("7ef45cd525e95b7a86244bbd4eb4550914ad06301013958f4dd64d32ef7bc588")
	signPrivateKey, _ = hex.DecodeString("314852d7afb0d4c283692fef8a2cb40e30c7a5df2ed79994178c10ac168d6d977ef45cd525e95b7a86244bbd4eb4550914ad06301013958f4dd64d32ef7bc588")
	defaultSignature  = "l07qwsfn2dpCqic8jKro5ut2b6KaMbN3MvIubS5h6EAhBoSeYeNVH/cNfTWRcKYZhmnhBhtrSqYZl+Jrh+OnBA=="
	wrongSignature    = "l08qwsfn2dpCqic8jKro5ut2b6KaMbN3MvIubS5h6EAhBoSeYeNVH/cNfTWRcKYZhmnhBhtrSqYZl+Jrh+OnBA=="
)

func TestSignMessageWithPrivateKey(t *testing.T) {
	if val := SignMessageWithPrivateKey(defaultMessage, signPrivateKey); val != defaultSignature {
		t.Errorf("SignMessageWithPrivateKey(%v,%v)=%v; want %v", defaultMessage, signPrivateKey, val, defaultSignature)
	}
}

func TestVerifyMessageWithPublicKey(t *testing.T) {
	if val, err := VerifyMessageWithPublicKey(defaultMessage, defaultSignature, signPublicKey); !val || err != nil {
		t.Errorf("SignMessageWithPrivateKey(%v,%v,%v)=%v,%v; want %v,%v", defaultMessage, defaultSignature, signPrivateKey, val, err, true, nil)
	}

	if val, err := VerifyMessageWithPublicKey(defaultMessage, wrongSignature, signPublicKey); val || err != nil {
		t.Errorf("SignMessageWithPrivateKey(%v,%v,%v)=%v,%v; want %v,%v", defaultMessage, wrongSignature, signPrivateKey, val, err, false, nil)
	}

	if val, err := VerifyMessageWithPublicKey(defaultMessage, "abc", signPublicKey); val || err == nil {
		t.Errorf("VerifyMessageWithPublicKey(%v,%v,%v)=%v,%v; want %v,%v", defaultMessage, "abc", signPrivateKey, val, err, defaultSignature, "err")
	}
}
