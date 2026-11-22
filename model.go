package main

type Cliente struct {
	Id        string `json:"id" db:"id"`
	NomeRazao string `json:"nome_razao" db:"nome_razao"`
}
