package main

import (
	"time"
)

type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Ammount   int `json:"ammount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	Id        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastName string) *Account {

	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    0,
		Balance:   0,
		CreatedAt: time.Now().UTC(),
	}
}
