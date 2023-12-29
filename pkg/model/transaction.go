package model

type Transaction struct {
	ID   string `json:"id"`
	User *User  `json:"user"`
	Book *Book  `json:"book"`
}

type NewBorrow struct {
	BookID string `json:"bookID"`
	UserID string `json:"userID"`
}
