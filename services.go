package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func RegistartionInsert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		address := r.FormValue("address")
		identityNumber := r.FormValue("identity_number")
		amount := r.FormValue("amount")

		accountIns, errPrepareAccount := db.Prepare(`insert into account(name, address, identity_number, created_at) values (?,?,?,?)`)
		if errPrepareAccount != nil {
			panic(errPrepareAccount.Error())
		}

		res, errExecAccount := accountIns.Exec(name, address, identityNumber, time.Now())
		if errExecAccount != nil {
			panic(errExecAccount.Error())
		}

		log.Println("INSERT Account : Name: " + name + " | Address: " + address + " | IdentityNumber: " + identityNumber)

		lastID, errLastID := res.LastInsertId()
		if errLastID != nil {
			panic(errLastID.Error())
		}

		deposits, errDeposits := db.Prepare(`insert into deposit(account_id, amount, created_at) values (?,?,?)`)
		if errDeposits != nil {
			panic(errDeposits.Error())
		}

		deposits.Exec(lastID, amount, time.Now())

		accountID := strconv.FormatInt(lastID, 10)
		log.Println("INSERT Deposit : AccountID: " + accountID + " | Amount: " + amount)

		balances, errBalances := db.Prepare(`insert into balance(account_id, balance, created_at) values (?,?,?)`)
		if errBalances != nil {
			panic(errBalances.Error())
		}

		balances.Exec(lastID, amount, time.Now())

		log.Println("INSERT Balance : AccountID: " + accountID + " | Balance: " + amount)

	}
	defer db.Close()
	http.Redirect(w, r, "/account", 301)
}

func DepositUpdate(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		id := r.FormValue("uid")
		amount := r.FormValue("amount")

		deposits, errDeposits := db.Prepare(`insert into deposit(account_id, amount, created_at) values (?,?,?)`)
		if errDeposits != nil {
			panic(errDeposits.Error())
		}

		deposits.Exec(id, amount, time.Now())
		log.Println("INSERT Deposit : AccountID: " + id + " | Amount: " + amount)

		updateBalance, err := db.Prepare("UPDATE balance SET balance=balance+?, updated_at=? WHERE account_id=?")
		if err != nil {
			panic(err.Error())
		}
		updateBalance.Exec(amount, time.Now(), id)
		log.Println("UPDATE Balance : amount : " + amount + " | account number : " + id)
	}
	defer db.Close()
	http.Redirect(w, r, "/account", 301)
}
