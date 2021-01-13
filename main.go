package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env files")
	}
}
func main() {
	name := os.Getenv("FIRST_NAME")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var lastName string = "sipatoha"
		fmt.Fprintf(w, "hai world %s %s! \n", name, lastName)
	})
	http.ListenAndServe(":8095", nil)
}
