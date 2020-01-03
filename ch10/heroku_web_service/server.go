package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
)

type Post struct{
	Id int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func main(){
	server := http.Server{
		Addr:":"+os.Getenv("PORT"),
	}

	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request){
	var err error

	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request)(err error){
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil{
		return
	}

	post, err := Retrieve(id)
	if err != nil{
		return
	}

	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil{
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	return
}

func handlePost(w http.ResponseWriter, r *http.Request)(err error){
    len := r.ContentLength
    body := make([]byte, len)
    r.Body.Read(body)
    fmt.Fprintln(os.Stdout, string(body))

    var post Post
    json.Unmarshal(body, &post)//ここで正常に読み込まれていない
	fmt.Fprintln(os.Stdout, post)
    err = post.Create()
    if err != nil{
    	return
	}

	w.WriteHeader(200)
    return
}

func handlePut(w http.ResponseWriter, r *http.Request)(err error){
	// Designate ID
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil{
		return
	}

	post, err := Retrieve(id)
	if err != nil{
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &post)
	err = post.Update()
	if err != nil{
		return
	}

	w.WriteHeader(200)

	return
}

func handleDelete(w http.ResponseWriter, r *http.Request)(err error){
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil{
		return
	}

	post, err := Retrieve(id)
	if err != nil {
		return
	}

	err = post.Delete()
	if err != nil{
		return
	}

	w.WriteHeader(200)

	return
}