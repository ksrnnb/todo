package middleware

import (
	"log"
	"net/http"
	"redis"
)

// Csrf middleware checks whether csrf token exists
func Csrf(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	sessionCookie, err := r.Cookie("sessionID")
	sessionID := sessionCookie.Value

	if err != nil {
		log.Println("session doesn't exist")
	}

	csrfTokenRedis, err := redis.GetToken(sessionID)

	if err != nil {
		log.Println("token couldn't find")
	}

	csrfTokenForm := r.PostFormValue("_token")

	if csrfTokenRedis != csrfTokenForm {
		log.Println("token didn't match")
	}

	return w, r
}
