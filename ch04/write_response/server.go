package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User string
	Threads []string
}

// body
func writeExample(w http.ResponseWriter, r *http.Request){
    str := `<html><head><title>Go Web Programming</title></head><body><h1>Hello World</h1></body></html>`

    w.Write([]byte(str))
}

// status code
func writeHeaderExample(w http.ResponseWriter, r *http.Request){
	// cannot change headers after calling WriteHeader
	w.WriteHeader(501)
	fmt.Fprintln(w, "Not Implemented")
}

// update header
func headerExample(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Location", "https://google.com")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User: "ace",
		Threads: []string{"No.1", "No.2", "No.3"},
	}
	// encode
	json, _ := json.Marshal(post)

	w.Write(json)
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1:8090",
	}

	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)

	server.ListenAndServe()
}
