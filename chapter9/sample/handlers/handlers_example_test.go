package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"log"
	"fmt"
)

func ExampleSendJson() {
	r, _ := http.NewRequest("GET", "/sendjson", nil)
	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, r)

	u := struct {
		Name string
		Email string
	}{}

	if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
		log.Println("ERROR: ", err)
	}
	fmt.Println(u)
	// Output:
	// {wzz wzhizhao@gmail.com}
}
