package validations

import (
	"fmt"
	"helpers"
	"net/http"
	"regexp"
	"strconv"
)

// 制御文字以外
func isNotControlChar(r *http.Request, key string) bool {
	str := r.PostFormValue(key)
	reg := regexp.MustCompilePOSIX(`^[^[:cntrl:]]+$`)

	ok := reg.MatchString(str)

	if !ok {
		setError("入力された文字に制御文字が含まれています")
	}

	return ok
}

func isInteger(r *http.Request, key string) bool {
	str := r.PostFormValue(key)
	_, err := strconv.Atoi(str)

	if err != nil {
		setError(key + "が整数ではありません")
	}

	return err == nil
}

//
func isUUID(r *http.Request) bool {
	path := helpers.GetPath(r)
	ok := helpers.IsUUID(path)

	if !ok {
		setError("パスがUUIDではありません")
	}

	return ok
}

func getID(r *http.Request) (id string) {
	id = r.PostFormValue("id")
	return
}


func getPath(r *http.Request) (path string) {
	path = helpers.GetPath(r)
	return
}

func setError(message string) {
	// cookieに入れる？
	fmt.Println(message)
}