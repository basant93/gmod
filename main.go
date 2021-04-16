package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Home Page!!"))
	})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Cannot start the server. Error : ", err)
	}

}
