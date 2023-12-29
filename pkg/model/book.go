package model

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Genre  Genre  `json:"genre"`
	Author string `json:"author"`
	Status Status `json:"status"`
}

type NewBook struct {
	Title  string `json:"title"`
	Genre  Genre  `json:"genre"`
	Author string `json:"author"`
}
