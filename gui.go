package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	go createTray()
	http.HandleFunc("/rest/gui/message", showMsg)
	http.ListenAndServe(":8081", nil)
}

func showMsg(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		msg := string(body)
		showMessageAll(msg)
	}
}
