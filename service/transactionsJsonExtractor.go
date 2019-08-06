package service

type Account struct {
	id_account int `json:"id"`
	id_connection int `json:"id_connection"`
	id_user int `json:"id_user"`
	currency Currency `json:"currency"`
	loan Loan `json:"loan"`
	original_name string `json:"original_name"`
	account_type string `json:"type"`
	transactions Transactions `json:"transactions"`
	balance float32 `json:"balance"`
}

type Currency struct {
	name string `json:"name"`
	symbol string `json:"symbol"`
}

type Loan struct {
	//TODO : traiter le cas des loans, pour le moment tous null pour simplifier
}

type Transactions struct {
	transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	wording string `json:"wording"`
	original_wording string `json:"original_wording"`
	id_cluster int `json:"id_cluster"`
	original_value float32 `json:"original_value"`
	value float32 `json:"value"`
	date string `json:"date"`
	transaction_type string `json:"type"`
	id_account int `json:"id_account"`
	id_transaction int `json:"id"`

}

