package main

import (
    
	"net/http"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte( "<html><head><title>The Chat!</title></head><body><h1>Hi there, I love %s!</h1></body></html>" ))
}

func main() {
	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
   
}