package rest

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
)

func MCEncrypt(data interface{}, key string) (string, error) {
	// Convert the data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Decode the key from hex to bytes
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	if len(keyBytes) != 32 {
		return "", fmt.Errorf("Needs a 256-bit key!")
	}

	// AES-256-CBC requires an IV
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Create the AES cipher
	mode := cipher.NewCBCEncrypter(block, iv)

	// Pad the data to be a multiple of the block size
	paddedData := pad(jsonData, aes.BlockSize)

	// Encrypt the data
	encrypted := make([]byte, len(paddedData))
	mode.CryptBlocks(encrypted, paddedData)

	// Create HMAC-SHA256 signature
	h := hmac.New(sha256.New, keyBytes)
	h.Write(encrypted)
	signature := h.Sum(nil)[:10]

	// Concatenate signature, IV, and encrypted data
	combined := append(signature, append(iv, encrypted...)...)

	// Base64 encode the result
	encoded := base64.StdEncoding.EncodeToString(combined)

	return encoded, nil
}

func pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}
