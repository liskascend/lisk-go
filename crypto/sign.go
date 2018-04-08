package crypto

import (
	"golang.org/x/crypto/ed25519"
)

// SignMessageWithPrivateKey takes a message and a privateKey and returns a signature as hex string
func SignMessageWithPrivateKey(message string, privKey []byte) []byte {
	rawMessage := []byte(message)

	signedMessage := ed25519.Sign(ed25519.PrivateKey(privKey), rawMessage)

	return signedMessage
}

// SignDataWithPrivateKey takes data and a privateKey and returns a signature
func SignDataWithPrivateKey(data []byte, privKey []byte) []byte {
	signedMessage := ed25519.Sign(ed25519.PrivateKey(privKey), data)

	return signedMessage
}

// VerifyMessageWithPublicKey takes a message, signature and publicKey and verifies it
func VerifyMessageWithPublicKey(message string, signature []byte, publicKey []byte) (bool, error) {
	isValid := ed25519.Verify(ed25519.PublicKey(publicKey), []byte(message), signature)
	return isValid, nil
}

// VerifyDataWithPublicKey takes data, a signature and a publicKey and verifies it
func VerifyDataWithPublicKey(data []byte, signature []byte, publicKey []byte) (bool, error) {
	isValid := ed25519.Verify(ed25519.PublicKey(publicKey), data, signature)
	return isValid, nil
}
