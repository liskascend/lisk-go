package crypto

import (
	"encoding/base64"

	"golang.org/x/crypto/ed25519"
)

// SignMessageWithPrivateKey takes a message and a privateKey and returns a signature as hex string
func SignMessageWithPrivateKey(message string, privKey []byte) string {
	rawMessage := []byte(message)

	signedMessage := ed25519.Sign(ed25519.PrivateKey(privKey), rawMessage)

	return base64.StdEncoding.EncodeToString(signedMessage)
}

// VerifyMessageWithPublicKey takes a message, base64 signature and publicKey and verifies it
func VerifyMessageWithPublicKey(message, signature string, publicKey []byte) (bool, error) {
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}

	isValid := ed25519.Verify(ed25519.PublicKey(publicKey), []byte(message), signatureBytes)
	return isValid, nil
}
