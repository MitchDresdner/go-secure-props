package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	b64 "encoding/base64"
	"fmt"
	"io"
)

func Decode(secret, ciphertext string) (msg string, error error){

	key := []byte(secret)

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Decode::NewCipher", err)
		return "", err
	}

	gcm, err := cipher.NewGCMWithNonceSize(c, NONCE_SIZE)
	if err != nil {
		fmt.Println("Decode::NewGCM", err)
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println("Decode::NonceSize > ciphertext")
		return "", nil
	}

	b64Dec, _ := b64.StdEncoding.DecodeString(ciphertext)

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println("encode::ReadFull", err)
		return "", err
	}

	//nonce, _ = hex.DecodeString("765700930cead4a2")
	nonce = b64Dec[:nonceSize]

	plaintext, err := gcm.Open(nil, nonce, b64Dec[nonceSize:], nil)
	if err != nil {
		fmt.Println("Decode::gcm.Open", err)
		return "", err
	}

	return string(plaintext), nil
}
