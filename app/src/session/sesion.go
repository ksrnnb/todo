package session

import (
	"fmt"
	"helpers"
	"net/http"
	"redis"
)

// Start starts session if it doesn't start
func Start(w http.ResponseWriter, r *http.Request) (token string) {
	sessionCookie, err := r.Cookie("sessionID")

	// cookieがない場合
	if err != nil {
		sessionID, token := createToken()
		storeToken(sessionID, token)
		setCookieSessionID(w, sessionID)

		return token
	}

	token = getToken(sessionCookie.Value)
	return token
}

// setCookieSessionID sets cookie / value is sessionID
func setCookieSessionID(w http.ResponseWriter, sessionID string) {
	cookie := createCookie(sessionID)
	http.SetCookie(w, cookie)
}

// createCookie creates cookie / value is sessionID
func createCookie(sessionID string) *http.Cookie {
	cookie := &http.Cookie{
		Name:  "sessionID",
		Value: sessionID,
	}

	return cookie
}

// createToken creates sessionID and token
func createToken() (sessionID string, token string) {
	sessionID = helpers.CreateRandomString32()
	token = helpers.CreateRandomString32()

	return
}

// storeToken stores csrf token
func storeToken(sessionID string, token string) {
	redis.StoreToken(sessionID, token)
}

// getToken gets csrf token
func getToken(sessionID string) (token string) {
	token, err := redis.GetToken(sessionID)

	// TODO: これでいいか要確認
	if err != nil {
		fmt.Println("csrfトークンが取得できませんでした")
		return
	}

	return token
}
