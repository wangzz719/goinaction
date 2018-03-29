package handlers

import (
	"net/http"
	"encoding/json"
)

func SendJson(w http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "wzz",
		Email: "wzhizhao@gmail.com",
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func Routes() {
	http.HandleFunc("/sendjson", SendJson)
}
