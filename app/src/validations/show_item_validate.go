package validations

import (
	"net/http"
)

// validation for show item
func ValidateShowItem(r *http.Request) bool {
	return isUUID(r)
}