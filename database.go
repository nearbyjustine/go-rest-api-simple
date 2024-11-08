package main

type ClientProfile struct {
	Email string
	Id    string
	Name  string
	Token string
}

var database = map[string]ClientProfile{
	"001": {
		Name:  "Ae",
		Id:    "001",
		Email: "ae@gmail.com",
		Token: "AE01",
	},
	"002": {
		Name:  "Beau",
		Id:    "002",
		Email: "beau@gmail.com",
		Token: "BEAU01",
	},
}
