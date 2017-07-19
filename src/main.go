package main

import (
	"net/http"
	"log"
	"os"
	"fmt"
)

var format string

func handler(w http.ResponseWriter, r *http.Request) {
	target := fmt.Sprintf(format, r.Host, r.RequestURI)

	w.Header().Set("Location", target)
}

func main() {
	port, set := os.LookupEnv("PORT")
	var addr string
	if !set {
		addr = ":80"
	} else {
		addr = fmt.Sprintf(":%s", port)
	}

	format, set = os.LookupEnv("FORMAT")
	if !set {
		log.Panic("FORMAT env unset")
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(addr, nil)
}
