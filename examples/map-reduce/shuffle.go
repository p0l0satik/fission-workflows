package main

import (
	"io/ioutil"
	"net/http"
	"sort"
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
	sort.Strings(words)
	combined := ""
	last := ""
	for _, word := range words {
		wordNum := strings.Split(word, ":")
		if wordNum[0] == last {
			combined += ";" + wordNum[1]
		} else {
			combined = strings.TrimRight(combined, ";")
			if len(combined) > 0 {
				combined += ","
			}
			combined += wordNum[0] + ":" + wordNum[1]
			last = wordNum[0]
		}
	}
	w.Write([]byte(combined))
}
