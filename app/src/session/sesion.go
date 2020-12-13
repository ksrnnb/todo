package session

import (
	"helpers"
	"net/http"
	"redis"
)

// Start starts session if it doesn't start
// TODO: セッション固定化対策は大丈夫か？　要確認
func Start(w http.ResponseWriter, r *http.Request) (token string) {
	sessionCookie, err := r.Cookie("sessionID")

	// cookieがない場合
	if err != nil {
		return getNewToken(w)
	}

	token, err = getTokenFromRedis(sessionCookie.Value)

	// cookieは設定されているけれど、tokenが見つからない場合
	if err != nil {
		return getNewToken(w)
	}

	return token
}

// getNewToken creates new token and set sessionID
func getNewToken(w http.ResponseWriter) (token string) {
	sessionID, token := createToken()
	storeToken(sessionID, token)
	setCookieSessionID(w, sessionID)

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

// getTokenFromRedis gets csrf token
func getTokenFromRedis(sessionID string) (token string, err error) {
	token, err = redis.GetToken(sessionID)

	return
}
