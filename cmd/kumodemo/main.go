package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/isAlive", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	http.HandleFunc("/isReady", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "KU-MO says hello @ %s", time.Now().Format(time.RFC1123))
		if err != nil {
			fmt.Println(err)
		}
	})

	fmt.Println("Starting...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
