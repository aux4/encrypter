package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
    charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func generateSecret(length int) (string, error) {
  secret := make([]byte, length)
  charsetLength := big.NewInt(int64(len(charset)))

  for i := range secret {
    position, err := rand.Int(rand.Reader, charsetLength)
    if err != nil {
      return "", fmt.Errorf("Error generating secret: %s", err.Error())
    }
    secret[i] = charset[position.Int64()]
  }

  return string(secret), nil
}
