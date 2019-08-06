package model

type Json_T struct {
	Id_user int
	Connections []Connection_T
}

type Connection_T struct {
	Id_connection int
	Last_update string
	Accounts []Account
}

type Account struct {
	//loan Loan
	Id_account int
	Last_update string
	Currency Currency
	Original_name string
	Balance float32
	Account_type string
	Transactions []Transaction
}

type Loan struct {
	//TODO : traiter le cas des loans dans Account
}

type Currency struct {
	Name string
	Symbol string
}

type Transaction struct {
	Wording string
	Original_wording string
	//Id_cluster int
	Transaction_type string
	Date string
	Rdate string
}