package main

import "net/http"

func main(){
	server := http.Server{
		Addr: "127.0.0.1:8090",
		Handler: nil,
	}

	server.ListenAndServeTLS("cert.pem", "key.pem")
}