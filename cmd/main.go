package main

import (
	"course/internal/handlers/wallet"
	"net/http"
)

func main() {
	http.HandleFunc("/wallet", wallet.Handle)

	http.ListenAndServe(":9001", nil)
}
