package crypto

import (
	"encoding/hex"
	"testing"
)

var (
	defaultPassword            = "password"
	defaultMessage2            = "Some default text."
	encryptedMessage, _        = hex.DecodeString("299390b9cbb92fe6a43daece2ceaecbacd01c7c03cfdba51d693b5c0e2b65c634115")
	encryptedMessageNonce, _   = hex.DecodeString("df4c8b09e270d2cb3f7b3d53dfa8a6f3441ad3b14a13fb66")
	defaultSecondPublicKey, _  = hex.DecodeString("7ef45cd525e95b7a86244bbd4eb4550914ad06301013958f4dd64d32ef7bc588")
	defaultSecondPrivateKey, _ = hex.DecodeString("314852d7afb0d4c283692fef8a2cb40e30c7a5df2ed79994178c10ac168d6d977ef45cd525e95b7a86244bbd4eb4550914ad06301013958f4dd64d32ef7bc588")
)

func TestEncryptDecryptPassphraseWithPassword(t *testing.T) {
	data, tag, iv, salt, err := EncryptPassphraseWithPassword(defaultMessage, defaultPassword)
	if err != nil {
		t.Errorf("EncryptPassphraseWithPassword(%v,%v) throws error: %v", defaultMessage, defaultPassword, err)
	}

	result, err := DecryptPassphraseWithPassword(data, iv, salt, tag, defaultPassword)
	if err != nil {
		t.Errorf("DecryptPassphraseWithPassword(%v,%v) throws error: %v", defaultMessage, defaultPassword, err)
	}

	if result != defaultMessage {
		t.Errorf("DecryptPassphraseWithPassword(...,%v)=%v; want %v", defaultPassword, result, defaultMessage)
	}

	result, err = DecryptPassphraseWithPassword(data, iv, salt, []byte{}, defaultPassword)
	if err == nil {
		t.Errorf("DecryptPassphraseWithPassword(invalid auth tag)=%v; should throw error", result)
	}
}

func TestEncryptDecryptMessageWithPrivateKey(t *testing.T) {
	// Integration test
	enc, nonce := EncryptMessageWithPrivateKey(defaultMessage, defaultPrivateKey, defaultSecondPublicKey)
	result, successful := DecryptMessageWithPrivateKey(enc, nonce, defaultSecondPrivateKey, defaultPublicKey)

	if result != defaultMessage || !successful {
		t.Errorf("DecryptPassphraseWithPassword(EncryptMessageWithPrivateKey)=%v,%v; should be %v,%v", result, successful, defaultMessage, true)
	}

	// Test with constant encrypted data
	result, successful = DecryptMessageWithPrivateKey(encryptedMessage, encryptedMessageNonce, defaultSecondPrivateKey, defaultSecondPublicKey)
	if result != defaultMessage || !successful {
		t.Errorf("DecryptPassphraseWithPassword(EncryptMessageWithPrivateKey)=%v,%v; should be %v,%v", result, successful, defaultMessage2, true)
	}
}
