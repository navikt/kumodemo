package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "KU-MO says hello @ %s", time.Now().Format(time.RFC822))
		if err != nil {
			fmt.Println(err)
		}
	})
	fmt.Println("Starting...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
