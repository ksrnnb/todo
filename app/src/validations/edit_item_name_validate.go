package validations

import (
	"net/http"
)

func ValidateEditItemName(r *http.Request) bool {
	return isUUID(r) && isInteger(r, "id") && isNotControlChar(r, "name")
}