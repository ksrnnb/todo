package validations

import (
	"net/http"
)

func ValidateDeleteItem(r *http.Request) bool {
	return isUUID(r) && isInteger(r, "id")
}