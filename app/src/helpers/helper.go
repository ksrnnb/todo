package helpers

import (
	"crypto/rand"
	"math/big"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

// IsUUID judges argumets is uuid or not
func IsUUID(str string) bool {
	_, err := uuid.Parse(str)

	return err == nil
}

// CreateRandomString32 creates random 32 digit string
func CreateRandomString32() string {
	seed := "abcdefghijklmnopqlstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seedLength := big.NewInt(int64(len(seed)))
	// len(seed) returns int
	// big.NewInt(int64) returns big.Int

	var token string
	tokenLength := 32

	for i := 0; i < tokenLength; i++ {
		randomBigInt, _ := rand.Int(rand.Reader, seedLength)
		// rand.Int(io.Reader, big.Int) returns big.Int

		token = token + string(seed[randomBigInt.Int64()])
		// big.Int.Int64() returns Int64
	}

	return token
}

// GetPath remove first slash /
func GetPath(r *http.Request) string {
	return strings.TrimLeft(r.URL.Path, "/")
}
