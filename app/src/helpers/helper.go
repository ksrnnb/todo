package helpers

import (
	"github.com/google/uuid"
)

// uuidかどうかを判定
func IsUUID(str string) bool {

	_, err := uuid.Parse(str)

	return err == nil
}