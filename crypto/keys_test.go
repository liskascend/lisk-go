package crypto

import (
	"bytes"
	"encoding/hex"
	"testing"
)

var (
	defaultPassphrase     = "secret"
	defaultPassphraseHash = "2bb80d537b1da3e38bd30361aa855686bde0eacd7162fef6a25fe97bf527a25b"
	defaultPrivateKey, _  = hex.DecodeString("2bb80d537b1da3e38bd30361aa855686bde0eacd7162fef6a25fe97bf527a25b5d036a858ce89f844491762eb89e2bfbd50a4a0a0da658e4b2628b25b117ae09")
	defaultPublicKey, _   = hex.DecodeString("5d036a858ce89f844491762eb89e2bfbd50a4a0a0da658e4b2628b25b117ae09")
	defaultAddress        = "18160565574430594874L"
)

func TestGetAddressFromPublicKey(t *testing.T) {
	if val := GetAddressFromPublicKey(defaultPublicKey); val != defaultAddress {
		t.Errorf("GetAddressFromPublicKey(%v)=%v; want %v", defaultPublicKey, val, defaultAddress)
	}
}

func TestGetPublicKeyFromSecret(t *testing.T) {
	if val := GetPublicKeyFromSecret(defaultPassphrase); !bytes.Equal(val, defaultPublicKey) {
		t.Errorf("GetPublicKeyFromSecret(%v)=%v; want %v", defaultPassphrase, val, defaultPublicKey)
	}
}

func TestGetPrivateKeyFromSecret(t *testing.T) {
	if val := GetPrivateKeyFromSecret(defaultPassphrase); !bytes.Equal(val, defaultPrivateKey) {
		t.Errorf("GetPrivateKeyFromSecret(%v)=%v; want %v", defaultPassphrase, val, defaultPrivateKey)
	}
}
