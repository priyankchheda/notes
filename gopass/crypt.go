package gopass

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

func encrypt(text []byte, key []byte, fileName string) {
	// generate a new aes cipher using our 32 bytes long key
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	// gcm or Golois/Counter Mode, is a mode of operation
	// for symmetric key Cryptographic block ciphers
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())

	// populates our nonce with a cryptographically secure random sequence
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	// here we encrypt our text using the Seal function.
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	//time, for a given key.
	ciphertext := gcm.Seal(nonce, nonce, text, nil)

	err = ioutil.WriteFile(fileName, ciphertext, 0600)
	if err != nil {
		fmt.Println(err)
	}
}

// Decrypt is the decryption function
func Decrypt(fileName string, key []byte) []byte {
	ciphertext, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}
	return plaintext
}
