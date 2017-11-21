package main

const (
	cityIdVancouver    = "13672f03-5cf2-4878-a2cb-1d4e453c56da"
	cityIdVictoria     = "050a94ab-728d-43bc-9be6-b54f79d330da"
	cityIdSanFrancisco = "7ce9f77c-32a0-4683-b910-5d82b9a656d6"
)

type company struct {
	ID           string `json:"id"`
	OfficialName string `json:"officialName"`
	CommonName   string `json:"commonName"`
	CityID       string `json:"cityId"`
}

var companies = []*company{
	{
		ID:           "acacde31-7a68-4a58-8a6c-4da16043f210",
		OfficialName: "Klue Labs Inc.",
		CommonName:   "Klue",
		CityID:       cityIdVancouver,
	},
	{
		ID:           "6ded79f9-9875-4c9e-bfe7-c3689c6e083a",
		OfficialName: "Teespring, Inc.",
		CommonName:   "Teespring",
		CityID:       cityIdSanFrancisco,
	},
	{
		ID:           "8bc1489f-a588-4861-8337-8de0a8a1f6d4",
		OfficialName: "Hootsuite Media Inc.",
		CommonName:   "Hootsuite",
		CityID:       cityIdVancouver,
	},
	{
		ID:           "ef145fab-c14e-4330-a592-c8dd135df2ba",
		OfficialName: "Metalogix Software Corp.",
		CommonName:   "Metalogix",
		CityID:       cityIdVancouver,
	},
	{
		ID:           "49117105-2e45-401d-a739-586ce3518e93",
		OfficialName: "ShipConstructor Software Inc.",
		CommonName:   "ShipConstructor",
		CityID:       cityIdVictoria,
	},
}

func find(id string) *company {
	for _, c := range companies {
		if c.ID == id {
			return c
		}
	}

	return nil
}
