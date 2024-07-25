package models

type Request struct {
	Amount    int   `json:"amount"`
	Banknotes []int `json:"banknotes"`
}
