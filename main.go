package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("hello world"))
		if err != nil {
			log.Println("Error writing", err)
		}

	})
	log.Println("Running")
	log.Println(http.ListenAndServe(":8080", nil))

}
