package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func sayHello(w http.ResponseWriter, _ *http.Request) {
	page, err := ioutil.ReadFile("pages/home.html")
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = w.Write(page)
}
