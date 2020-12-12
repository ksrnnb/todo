package helpers

import (
	"github.com/google/uuid"
	"crypto/rand"
	"math/big"
)

// uuidかどうかを判定
func IsUUID(str string) bool {
	_, err := uuid.Parse(str)

	return err == nil
}

func CreateRandomString32() string {
	seed := "abcdefghijklmnopqlstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed_len := big.NewInt(int64(len(seed)))
	// len(seed) returns int
	// big.NewInt(int64) returns big.Int
	
	var token string
	token_len := 32

	for i := 0; i < token_len; i++ {
		randomBigInt, _ := rand.Int(rand.Reader, seed_len)
		// rand.Int(io.Reader, big.Int) returns big.Int

		token = token + string(seed[randomBigInt.Int64()])
		// big.Int.Int64() returns Int64
	}

	return token
}