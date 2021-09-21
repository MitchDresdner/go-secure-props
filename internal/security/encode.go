package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	b64 "encoding/base64"
	"fmt"
	"io"
)

// encode a string, hardcoded to use AES
//    see: https://pkg.go.dev/crypto/cipher#NewGCMWithNonceSize
func Encode(secret, msg string, ) (string, error) {

	text := []byte(msg)
	key := []byte(secret)

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("encode::NewCipher", err)
		return "", err
	}

	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCMWithNonceSize(c, NONCE_SIZE)
	if err != nil {
		fmt.Println("encode::NewGCM", err)
		return "", err
	}

	// byte array the sized by the nonce, passed later to Seal
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println("encode::ReadFull", err)
		return "", err
	}

	// Encrypt and authenticates plaintext, authenticate the additional data
	// and append the result to dst returning the updated slice.
	// Nonce must be of length NonceSize() bytes and unique for a given key.
	//
	//enc := gcm.Seal(nonce, nonce, text, nil)
	enc := gcm.Seal(nonce, nonce, text, nil)

	// Base64 encode the encrypted data
	b64Enc := b64.StdEncoding.EncodeToString(enc)

	// Decode and verify result
	//fmt.Println("ENC: ", enc)
	//fmt.Println("Encode B64: ", b64Enc)
	//
	//b64Dec, _ := b64.StdEncoding.DecodeString(b64Enc)
	//plaintext, err := gcm.Open(nil, nonce, b64Dec, nil)
	//if err != nil {
	//	fmt.Println("Decode::gcm.Open", err)
	//	return "", err
	//}
	//
	//fmt.Println("Decoded: ", string(plaintext))

	return b64Enc, nil
}

