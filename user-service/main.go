package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/user-bank", UserBank())
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func UserBank() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uri := fmt.Sprintf("http://%s/bank", os.Getenv("PAYMENT_HOST"))
		res, err := http.Get(uri)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		decode := json.NewDecoder(res.Body)
		var payment PaymentResponse
		err = decode.Decode(&payment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(fmt.Sprintf("bank milik user: %s", payment.Bank)))
	}
}

type PaymentResponse struct {
	Bank string `json:"bank"`
}
