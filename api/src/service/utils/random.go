package utils

import (
	"crypto/rand"
	"crypto/sha256"
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


func TextHasher (text string) string{
	var hash = sha256.New()

	for i:= 0; i<3; i++{
		hash.Write([]byte(text))
		text = hex.EncodeToString(hash.Sum(nil))
	}

	return text
}

