package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type PaymentResponse struct {
	Bank string `json:"bank"`
}

func main() {
	port := ":8055"
	http.HandleFunc("/", GetAll())
	http.HandleFunc("/bank", GetBank())
	log.Fatal(http.ListenAndServe(port, nil))
}

func GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		payment := []PaymentResponse{
			{
				Bank: "BCA",
			},
			{
				Bank: "Mandiri",
			},
		}

		encode, err := json.Marshal(payment)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(encode)
	}
}
func GetBank() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		payment := PaymentResponse{
			Bank: "BCA",
		}

		encode, err := json.Marshal(payment)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(encode)
	}
}
