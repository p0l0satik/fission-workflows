package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	body := string(reqBody)
	words := strings.Split(body, ",")
	mapped := ""
	for _, word := range words {
		mapped += word + ":1,"
	}
	mapped = strings.TrimRight(mapped, ",")
	w.Write([]byte(mapped))
}
