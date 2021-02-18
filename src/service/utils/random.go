package utils

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func RandomString(size int) string {
	var randomBytes = make([]byte, size)

	_, err := rand.Read(randomBytes)

	if err != nil {
		log.Fatal("Failed on generate bytes")
	}

	return  hex.EncodeToString(randomBytes)
}
