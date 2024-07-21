package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func encrypt(secret string, initVector string, text string) (string, error) {
	var plaintextBlock []byte
	length := len(text)

	if length%len(initVector) != 0 {
		padding := len(initVector) - (length % len(initVector))
    plaintextBlock = make([]byte, length+padding)
    copy(plaintextBlock[length:], bytes.Repeat([]byte{uint8(padding)}, padding))
	}

  copy(plaintextBlock, text)

  block, err := aes.NewCipher([]byte(secret))
  if err != nil {
    return "", fmt.Errorf("Error creating new cipher: %s", err.Error())
  }

  encrypted := make([]byte, len(plaintextBlock))

  mode := cipher.NewCBCEncrypter(block, []byte(initVector))
  mode.CryptBlocks(encrypted, plaintextBlock)

  hex := base64.StdEncoding.EncodeToString(encrypted)
	return hex, nil
}

func decrypt(secret string, initVector string, text string) (string, error) {
  decoded, err := base64.StdEncoding.DecodeString(text)
  if err != nil {
    return "", fmt.Errorf("Error decoding text: %s", err.Error())
  }

  block, err := aes.NewCipher([]byte(secret))
  if err != nil {
    return "", fmt.Errorf("Error creating new cipher: %s", err.Error())
  }

  decrypted := make([]byte, len(decoded))
  mode := cipher.NewCBCDecrypter(block, []byte(initVector))
  mode.CryptBlocks(decrypted, decoded)

  plaintext := unpadding(decrypted)
  return string(plaintext), nil
}

func unpadding(bytes []byte) []byte {
  length := len(bytes)
  padding := int(bytes[length-1])
  return bytes[:length-padding]
}
