package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func Encrypt(plainText string, key []byte) string {
	cipherText, err := encode(plainText, key)
	if err != nil {
		fmt.Println("Encryption error:", err)
		return ""
	}
	return cipherText
}

func Decrypt(cipherText string, key []byte) string {
	decrypted, err := decode(cipherText, key)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return ""
	}
	return decrypted
}

func encode(plainText string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Pad the plaintext if needed
	plaintext := []byte(plainText)
	if len(plaintext)%aes.BlockSize != 0 {
		padding := aes.BlockSize - (len(plaintext) % aes.BlockSize)
		paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
		plaintext = append(plaintext, paddingText...)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decode(cipherText string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext is too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// Unpad the plaintext
	padding := int(ciphertext[len(ciphertext)-1])
	ciphertext = ciphertext[:len(ciphertext)-padding]

	return string(ciphertext), nil
}
