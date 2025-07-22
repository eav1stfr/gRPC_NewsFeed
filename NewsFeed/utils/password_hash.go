package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

func Hash(password string) (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", ErrorGeneratingSaltForHashing
	}

	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	saltBase64 := base64.StdEncoding.EncodeToString(salt)
	hashBase64 := base64.StdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("%s.%s", saltBase64, hashBase64)
	return encodedHash, nil
}

func VerifyPassword(password, encodedHash string) error {
	parts := strings.Split(encodedHash, ".")
	if len(parts) != 2 {
		return InvalidEncodedHashPasswordError
	}

	saltBase64 := parts[0]
	hashedPasswordBase64 := parts[1]

	salt, err := base64.StdEncoding.DecodeString(saltBase64)
	if err != nil {
		return DecodingPasswordError
	}

	hashedPassword, err := base64.StdEncoding.DecodeString(hashedPasswordBase64)
	if err != nil {
		return DecodingPasswordError
	}

	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	if len(hash) != len(hashedPassword) {
		return IncorrectPassword
	}

	if subtle.ConstantTimeCompare(hash, hashedPassword) == 1 {
		return nil
	}
	return IncorrectPassword
}
