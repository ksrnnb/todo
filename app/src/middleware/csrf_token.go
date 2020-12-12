package middleware

import (
	"log"
	"net/http"
	"redis"
)

// Csrf middleware checks whether csrf token exists
func Csrf(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	sessionCookie, err := r.Cookie("sessionID")

	if err != nil {
		log.Fatalln("session doesn't exist")
	}

	csrfToken, err := redis.GetToken(sessionCookie.Value)

	if err != nil {
		log.Fatalln("token couldn't find")
	}

	if csrfToken != sessionCookie.Value {
		log.Fatalln("token did't match")
	}

	return w, r
}
