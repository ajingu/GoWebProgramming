package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request){
	if file, _, err := r.FormFile("uploaded"); err == nil{
		if data, err := ioutil.ReadAll(file); err == nil{
			fmt.Fprintln(w, string(data))
		}
	}
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1:8090",
	}

	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
