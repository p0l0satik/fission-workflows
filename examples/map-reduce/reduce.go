package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
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
	reduced := ""
	for _, word := range words {
		wordNum := strings.Split(word, ":")
		nums := strings.Split(wordNum[1], ";")
		reduced += wordNum[0] + ":" + strconv.Itoa(len(nums)) + ", "
	}
	reduced = strings.TrimRight(reduced, ", ")
	w.Write([]byte(reduced))
}
