package models

import "time"

type Transaction struct {
	Amount      float64   `json:"amount"`
	Beneficiary string    `json:"beneficiary"`
	Currency    string    `json:"currency"`
	Id          string    `json:"id"`
	Sender      string    `json:"sender"`
	Timestamp   time.Time `json:"timestamp"`
}
