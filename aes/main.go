package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {

}
func Encrypt(text string) (string, error) {
	plaintext := []byte(text)
	key := []byte("*eaciit-standard-chartered-apps*")
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	result := fmt.Sprintf("%x", ciphertext)
	return result, nil
}

func Decrypt(text string) (string, error) {
	ciphertext, err := hex.DecodeString(text)
	if err != nil {
		return "", err
	}
	key := []byte("*eaciit-standard-chartered-apps*")
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", err
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	res, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%s", res)
	return result, nil
}
