package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request){
    h := r.Header
    fmt.Fprintln(w, h)
}

func body(w http.ResponseWriter, r *http.Request){
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1:8090",
	}

	http.HandleFunc("/headers", headers)
	http.HandleFunc("/body", body)

	server.ListenAndServe()
}
