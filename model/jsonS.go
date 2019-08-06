package model

type Json_S struct {
	Id_user     int
	Connections []Connection_S
}

type Connection_S struct {
	Id_connection int
	Last_update   string
	Id_accounts   []int
}