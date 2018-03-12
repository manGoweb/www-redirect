package main

import (
	"net/http"
	"log"
	"os"
	"fmt"
	"strings"
)

const (
	HostVar = "$host"
	PathVar = "$path"
)

var format string

func handler(w http.ResponseWriter, r *http.Request) {
	target := format
	target = strings.Replace(target, HostVar, r.Host, -1)
	target = strings.Replace(target, PathVar, r.RequestURI, -1)

	w.Header().Set("Location", target)
	w.WriteHeader(301)
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

	format = strings.Replace(format, "%s", HostVar, 1)
	format = strings.Replace(format, "%s", PathVar, 1)
	log.Println("final format:", format)

	http.HandleFunc("/", handler)
	http.ListenAndServe(addr, nil)
}
