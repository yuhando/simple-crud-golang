package main

import (
	"net/http"
	"text/template"

	"github.com/yuhando/tniscodingtest3/model"
)

var pathform = "templates/*"
var tmpl = template.Must(template.ParseGlob(pathform))

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Index", nil)
}

func Registration(w http.ResponseWriter, r *http.Request) {
	Account := make(map[string]interface{})
	Account["Title"] = "Registration"
	Account["Slug"] = "registration/add"
	Account["ID"] = ""
	Account["Name"] = ""
	Account["Address"] = ""
	Account["IdentityNumber"] = ""
	Account["Disabled"] = ""

	tmpl.ExecuteTemplate(w, "Registration", Account)
}

func Account(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT id, name, address, identity_number FROM account ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	account := model.Accounts{}
	resAccount := []model.Accounts{}

	for selDB.Next() {
		var id int
		var name, address, identityNumber string

		err = selDB.Scan(&id, &name, &address, &identityNumber)
		if err != nil {
			panic(err.Error())
		}
		account.ID = id
		account.Name = name
		account.Address = address
		account.IdentityNumber = identityNumber

		resAccount = append(resAccount, account)
	}
	tmpl.ExecuteTemplate(w, "Account", resAccount)
	defer db.Close()
}

func Deposit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	id := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT id, name, address, identity_number FROM account WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}

	Account := make(map[string]interface{})
	for selDB.Next() {
		var id int
		var name, address, identityNumber string
		err = selDB.Scan(&id, &name, &address, &identityNumber)
		if err != nil {
			panic(err.Error())
		}

		Account["Title"] = "Deposit"
		Account["Slug"] = "deposit/update"
		Account["ID"] = id
		Account["Name"] = name
		Account["Address"] = address
		Account["IdentityNumber"] = identityNumber
		Account["Disabled"] = "disabled"
	}

	tmpl.ExecuteTemplate(w, "Registration", Account)
	defer db.Close()
}

func HistoryDeposit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("select dep.id as IDTransaction, acc.id as AccountID, acc.name, dep.amount from tniscodingtest.account acc inner join tniscodingtest.deposit dep on acc.id = dep.account_id ORDER BY dep.id DESC")
	if err != nil {
		panic(err.Error())
	}
	deposit := model.DepositHistories{}
	resDeposit := []model.DepositHistories{}

	for selDB.Next() {
		var idTransaction, accountID, amount int
		var name string

		err = selDB.Scan(&idTransaction, &accountID, &name, &amount)
		if err != nil {
			panic(err.Error())
		}
		deposit.IDTransaction = idTransaction
		deposit.AccountID = accountID
		deposit.Name = name
		deposit.Amount = amount

		resDeposit = append(resDeposit, deposit)
	}
	tmpl.ExecuteTemplate(w, "Deposit", resDeposit)
	defer db.Close()
}

func Balance(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("select acc.id as AccountID, acc.name, bal.balance from tniscodingtest.account acc inner join tniscodingtest.balance bal on acc.id = bal.account_id ORDER BY bal.id DESC")
	if err != nil {
		panic(err.Error())
	}
	balanceModel := model.Balances{}
	resBalance := []model.Balances{}

	for selDB.Next() {
		var accountID, balance int
		var name string

		err = selDB.Scan(&accountID, &name, &balance)
		if err != nil {
			panic(err.Error())
		}

		balanceModel.AccountID = accountID
		balanceModel.Name = name
		balanceModel.Balance = balance

		resBalance = append(resBalance, balanceModel)
	}
	tmpl.ExecuteTemplate(w, "Balance", resBalance)
	defer db.Close()
}
