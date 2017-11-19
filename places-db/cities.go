package main

var cities = []city{
	{
		ID:        "13672f03-5cf2-4878-a2cb-1d4e453c56da",
		Name:      "Vancouver",
		CountryID: "a944ee85-b10a-45b7-b749-2a32733ca26f",
	},
	{
		ID:        "050a94ab-728d-43bc-9be6-b54f79d330da",
		Name:      "Victoria",
		CountryID: "a944ee85-b10a-45b7-b749-2a32733ca26f",
	},
	{
		ID:        "e60e1202-b1f5-45a2-8153-90c7474dad51",
		Name:      "Lake Cowichan",
		CountryID: "a944ee85-b10a-45b7-b749-2a32733ca26f",
	},
}

func findCity(id string) *city {
	for _, c := range cities {
		if c.ID == id {
			return &c
		}
	}

	return nil
}

type city struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CountryID string `json:"countryId"`
}
