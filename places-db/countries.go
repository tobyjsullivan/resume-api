package main

type country struct {
	ID           string `json:"id"`
	OfficialName string `json:"officialName"`
	CommonName   string `json:"commonName"`
}

var countries = []*country{
	{
		ID:           "a944ee85-b10a-45b7-b749-2a32733ca26f",
		OfficialName: "Canada",
		CommonName:   "Canada",
	},
	{
		ID:           "e6b38bc7-1e14-4515-9e69-e1e385446ceb",
		OfficialName: "United States of America",
		CommonName:   "United States",
	},
}

func findCountry(id string) (*country) {
	for _, c := range countries {
		if c.ID == id {
			return c
		}
	}

	return nil
}
