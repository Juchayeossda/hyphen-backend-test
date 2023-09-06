package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "/ reqeusted")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "/hello reqeusted")
	})

	msg :=
		`
	[server] started
	[server] open port: 10000
	[server] open endpoint: /
	[server] open endpoint: /hello

	제발 잘 되라.
	`

	log.Print(msg)
	http.ListenAndServe(":10000", nil)
}
