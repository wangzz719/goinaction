package handlers_test

import (
	"github.com/wangzz719/goinaction/chapter9/sample/handlers"
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func init()  {
	handlers.Routes()
}

func TestSendJson(t *testing.T) {
	t.Log("Given the need to test the SendJSON endpoint.")
	req, err := http.NewRequest("GET", "/sendjson", nil)
	if err != nil {
		t.Fatal("\tShould be able to create a request.", ballotX, err)
	}
	t.Log("\tShould be able to create a request.", checkMark)

	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, req)

	if rw.Code != 200 {
		t.Fatal("\t Should receive \"200\"", ballotX, rw.Code)
	}
	t.Log("\t Should receive \"200\"", checkMark)

	u := struct {
		Name string
		Email string
	}{}

	if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
		t.Fatal("\tShould decode the response.", ballotX)
	}
	t.Log("\tShould decode the response.", checkMark)
	if u.Name == "wzz" {
		t.Log("\tShould have a Name.", checkMark)
	} else {
		t.Error("\tShould have a Name.", ballotX, u.Name)
	}

	if u.Email == "wzhizhao@gmail.com" {
		t.Log("\tShould have an Email.", checkMark)
	} else {
		t.Error("\tShould have an Email.", ballotX, u.Name)
	}
}