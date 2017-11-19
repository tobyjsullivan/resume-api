package main

type person struct {
	ID             string   `json:"id"`
	GivenNames     []string `json:"givenNames"`
	Surname        string   `json:"surname"`
	HometownCityID string   `json:"hometownCityID"`
	CurrentCityID  string   `json:"currentCityID"`
}

var people = []*person{
	{
		ID:             "40d8ed45-6977-47b5-92fa-7c4c4fe214c6",
		GivenNames:     []string{"Toby", "Jay"},
		Surname:        "Sullivan",
		HometownCityID: "e60e1202-b1f5-45a2-8153-90c7474dad51",
		CurrentCityID:  "13672f03-5cf2-4878-a2cb-1d4e453c56da",
	},
}

func find(id string) (*person) {
	for _, p := range people {
		if p.ID == id {
			return p
		}
	}

	return nil
}
