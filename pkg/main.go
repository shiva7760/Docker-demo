package main

import (
	"log"
	"net/http"

)

func main() {

	http.HandleFunc("/hello/", func(w http.ResponseWriter, r *http.Request) {
		
		w.Write([]byte("Hello World"))
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:4002", nil))
}
