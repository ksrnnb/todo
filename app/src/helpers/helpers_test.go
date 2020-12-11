package helpers

import (
	"testing"
	"github.com/google/uuid"
)

func TestIsUUIDFunction(t *testing.T) {
	uuid := uuid.New().String()
	
	if ! IsUUID(uuid) {
		t.Errorf("isUUID function is wrong because arg is uuid but it returns false")
	}

	if IsUUID("string_test") {
		t.Errorf("isUUID function is wrong because arg isn't uuid but it returns true")
	}
}