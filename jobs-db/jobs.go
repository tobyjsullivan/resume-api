package main

import "log"

const (
	personIdToby = "40d8ed45-6977-47b5-92fa-7c4c4fe214c6"
	companyIdKlue = "acacde31-7a68-4a58-8a6c-4da16043f210"
	companyIdTeespring = "6ded79f9-9875-4c9e-bfe7-c3689c6e083a"
	companyIdHootsuite = "8bc1489f-a588-4861-8337-8de0a8a1f6d4"
	companyIdMetalogix = "ef145fab-c14e-4330-a592-c8dd135df2ba"
	cityIdVancouver = "13672f03-5cf2-4878-a2cb-1d4e453c56da"
)

type job struct {
	ID string `json:"id"`
	EmployeePersonID string `json:"employeePersonId"`
	EmployerCompanyID string `json:"employerCompanyId"`
	LocationCityID string `json:"locationCityId"`
	Remote bool `json:"remote"`
}

var jobs = []*job {
	{
		// Klue
		ID:                "2c64fac7-c65a-4b28-ab0c-2aca30236fdd",
		EmployeePersonID:  personIdToby,
		EmployerCompanyID: companyIdKlue,
		LocationCityID:    cityIdVancouver,
		Remote:            false,
	},
	{
		// Teespring
		ID:                "ae5be743-9ef4-4fe4-bde3-b70bb3b7e625",
		EmployeePersonID:  personIdToby,
		EmployerCompanyID: companyIdTeespring,
		LocationCityID:    cityIdVancouver,
		Remote:            true,
	},
	{
		// Hootsuite
		ID:                "a0f38dd0-2e69-4db4-921f-20c96ea669a5",
		EmployeePersonID:  personIdToby,
		EmployerCompanyID: companyIdHootsuite,
		LocationCityID: cityIdVancouver,
		Remote: false,
	},
	{
		// Metalogix
		ID: "71f6bd07-594a-4f70-a7bd-ed8d9e997f95",
		EmployeePersonID: personIdToby,
		EmployerCompanyID: companyIdMetalogix,
		LocationCityID:cityIdVancouver,
		Remote:false,
	},
}

func find(id string) (*job) {
	for _, j := range jobs {
		if j.ID == id {
			return j
		}
	}

	return nil
}

func findByPersonId(personId string) []*job {
	out := []*job{}

	for _, j := range jobs {
		if j.EmployeePersonID == personId {
			log.Println("Responding with job:", j.ID)
			out = append(out, j)
		}
	}

	return out
}
