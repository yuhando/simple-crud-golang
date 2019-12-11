package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/registration", Registration)
	http.HandleFunc("/registration/add", RegistartionInsert)
	http.HandleFunc("/account", Account)
	http.HandleFunc("/deposit", Deposit)
	http.HandleFunc("/deposit/update", DepositUpdate)
	http.HandleFunc("/history", HistoryDeposit)
	http.HandleFunc("/balance", Balance)

	http.ListenAndServe(":8080", nil)
}
