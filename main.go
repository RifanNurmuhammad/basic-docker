package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var firstName string = "Roy"
		var lastName string = "sipatoha"
		fmt.Printf("hello world %s %s! \n", firstName, lastName)
	})
	http.ListenAndServe(":8095", nil)
}
