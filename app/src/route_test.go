package main

import (
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// メイン
func TestMain(m *testing.M) {
	setUp()
	os.Exit(m.Run())
}

// global
var mux *http.ServeMux

// 初期動作
func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/", handleRequest)
}

// トップページへのアクセスを確認する
func TestAccessTopPage(t *testing.T) {
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(writer, request)

	statusCodeCheck(t, writer, 200)
	containsCheck(t, writer, "make todo list")
}

// uuidをpostして、リダイレクトされることを確認
func TestRedirectToTodoPage(t *testing.T) {
	writer := httptest.NewRecorder()
	uuid := uuid.New().String()
	param := strings.NewReader("uuid=" + uuid)

	request, _ := http.NewRequest("POST", "/", param)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	mux.ServeHTTP(writer, request)

	statusCodeCheck(t, writer, 301)
}

// uuid以外をpostして、リダイレクトされることを確認
func TestValidatePostUUID(t *testing.T) {
	writer := httptest.NewRecorder()
	notUUID := "test_string"
	param := strings.NewReader("uuid=" + notUUID)

	request, _ := http.NewRequest("POST", "/", param)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	mux.ServeHTTP(writer, request)

	statusCodeCheck(t, writer, 422)
}

// 指定した文字が含まれることを確認する
func containsCheck(t *testing.T, writer *httptest.ResponseRecorder, str string) {
	// writer.Body type: *bytes.Buffer
	contains := strings.Contains(writer.Body.String(), str)

	if !contains {
		t.Errorf(`Cannot see %s`, str)
	}
}

// ステータスコードを確認する
func statusCodeCheck(t *testing.T, writer *httptest.ResponseRecorder, code int) {
	if writer.Code != code {
		t.Errorf("Response code is %v, expected %d", writer.Code, code)
	}
}
