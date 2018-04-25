package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"

	"github.com/agl/ed25519/extra25519"
	"github.com/kevinburke/nacl"
	"github.com/kevinburke/nacl/box"
	"golang.org/x/crypto/pbkdf2"
)

// EncryptPassphraseWithPassword encrypts the given passphrase with the password using AES-256-CBC
func EncryptPassphraseWithPassword(passphrase, password string) (result, tag, iv, salt []byte, err error) {
	iv = make([]byte, 16)
	rand.Read(iv)
	salt = make([]byte, 16)
	rand.Read(salt)

	key := getKeyFromPassword(password, salt)
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	gcm, err := cipher.NewGCMWithNonceSize(aesBlock, 16)
	result = gcm.Seal(result, iv, []byte(passphrase), nil)
	tag = result[len(result)-16:]
	result = result[:len(result)-16]

	return result, tag, iv, salt, err
}

// DecryptPassphraseWithPassword decrypts the given encrypted passphrase with the password using AES-256-CBC
func DecryptPassphraseWithPassword(encryptedPassphrase, iv, salt, tag []byte, password string) (string, error) {
	key := getKeyFromPassword(password, salt)
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	var result []byte
	gcm, err := cipher.NewGCMWithNonceSize(aesBlock, 16)
	if err != nil {
		return "", err
	}

	encryptedData := append(encryptedPassphrase, tag...)
	result, err = gcm.Open(result, iv, encryptedData, nil)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func getKeyFromPassword(password string, salt []byte) []byte {
	return pbkdf2.Key([]byte(password), salt, 1e6, 32, sha256.New)
}

// EncryptMessageWithPrivateKey encrypts and authenticates a message for the given recipient's public key
func EncryptMessageWithPrivateKey(message string, privKey, recipientPubKey []byte) (packet []byte, nonce []byte) {
	naclPrivKey := new([64]byte)
	naclPubKey := new([32]byte)
	convertedPrivKey := new([nacl.KeySize]byte)
	convertedPubKey := new([nacl.KeySize]byte)

	copy(naclPrivKey[:], privKey)
	copy(naclPubKey[:], recipientPubKey)

	extra25519.PrivateKeyToCurve25519(convertedPrivKey, naclPrivKey)
	extra25519.PublicKeyToCurve25519(convertedPubKey, naclPubKey)

	rawNonce := nacl.NewNonce()
	packet = box.Seal(packet, []byte(message), rawNonce, convertedPubKey, convertedPrivKey)

	return packet, (*rawNonce)[:]
}

// DecryptMessageWithPrivateKey decrypts and verifies an encrypted message using the recipients private key
func DecryptMessageWithPrivateKey(packet, nonce, privKey, senderPubKey []byte) (string, bool) {
	naclPrivKey := new([64]byte)
	naclPubKey := new([nacl.KeySize]byte)
	convertedPrivKey := new([nacl.KeySize]byte)
	convertedPubKey := new([nacl.KeySize]byte)
	naclNonce := new([nacl.NonceSize]byte)

	copy(naclPrivKey[:], privKey)
	copy(naclPubKey[:], senderPubKey)
	copy(naclNonce[:], nonce)

	extra25519.PrivateKeyToCurve25519(convertedPrivKey, naclPrivKey)
	extra25519.PublicKeyToCurve25519(convertedPubKey, naclPubKey)

	var message []byte
	message, successful := box.Open(message, packet, naclNonce, convertedPubKey, convertedPrivKey)

	return string(message), successful
}
