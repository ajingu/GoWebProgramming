package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func process(w http.ResponseWriter, r *http.Request){
	dirPath := os.Getenv("GOPATH") + "/src/go_web_programming/templates/"
	var t *template.Template
	rand.Seed(time.Now().Unix())

	if rand.Intn(10) > 5{
		t, _ = template.ParseFiles(dirPath+"layout.html", dirPath+"red_hello.html")
	} else {
		t, _ = template.ParseFiles(dirPath+"layout.html", dirPath+"blue_hello.html")
	}

	t.ExecuteTemplate(w, "layout", "")
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1:8090",
	}

	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
