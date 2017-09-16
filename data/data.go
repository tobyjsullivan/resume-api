package data

import "net/url"

type Person struct {
	FirstName string
	MiddleName string
	LastName string
	Website *url.URL
}

type Database struct {
	Me *Person
}

func LoadDatabase() *Database {
	website, _ := url.Parse("https://tobysullivan.net")

	return &Database{
		Me: &Person{
			FirstName: "Toby",
			MiddleName: "Jay",
			LastName: "Sullivan",
			Website: website,
		},
	}
}
