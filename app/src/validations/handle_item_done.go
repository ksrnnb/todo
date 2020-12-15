package validations

import (
	"net/http"
)

// validation for create item
func ValidateHandleDone(r *http.Request) bool {
	return isUUID(r) && isInteger(r, "id")
}