package validations

import (
	"net/http"
)

// validation for create item
func ValidateCreateItem(r *http.Request) bool {
	return isUUID(r) && isNotControlChar(r, "name")
}