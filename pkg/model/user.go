package model

type NewUser struct {
	Name string `json:"name"`
}

type User struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	BorrowedBooks []*Book `json:"borrowedBooks"`
}
