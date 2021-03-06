package main

import (
	"log"
)

const (
	personIdToby              = "40d8ed45-6977-47b5-92fa-7c4c4fe214c6"
	companyIdKlue             = "acacde31-7a68-4a58-8a6c-4da16043f210"
	companyIdTeespring        = "6ded79f9-9875-4c9e-bfe7-c3689c6e083a"
	companyIdHootsuite        = "8bc1489f-a588-4861-8337-8de0a8a1f6d4"
	companyIdMetalogix        = "ef145fab-c14e-4330-a592-c8dd135df2ba"
	companyIdShipConstructor  = "49117105-2e45-401d-a739-586ce3518e93"
	cityIdVancouver           = "13672f03-5cf2-4878-a2cb-1d4e453c56da"
	cityIdVictoria            = "050a94ab-728d-43bc-9be6-b54f79d330da"
	jobIdKlue                 = "2c64fac7-c65a-4b28-ab0c-2aca30236fdd"
	jobIdTeespring            = "ae5be743-9ef4-4fe4-bde3-b70bb3b7e625"
	jobIdHootsuite            = "a0f38dd0-2e69-4db4-921f-20c96ea669a5"
	jobIdHootsuiteContract    = "e6404e1f-e32f-4b4d-af56-651144999cd1"
	jobIdMetalogix            = "71f6bd07-594a-4f70-a7bd-ed8d9e997f95"
	jobIdShipConstructorTerm1 = "49cbf403-2c0d-4bf8-adbe-6ce7810e955d"
	jobIdShipConstructorTerm2 = "15e0e842-fccd-4e53-bcbf-815da4e17ad2"
)

type job struct {
	ID                string  `json:"id"`
	EmployeePersonID  string  `json:"employeePersonId"`
	EmployerCompanyID string  `json:"employerCompanyId"`
	LocationCityID    string  `json:"locationCityId"`
	Remote            bool    `json:"remote"`
	StartDate         *string `json:"startDate"`
	EndDate           *string `json:"endDate"`
}

func strVal(s string) *string {
	return &s
}

var jobs = []*job{
	{
		// Klue
		ID:                jobIdKlue,
		EmployeePersonID:  personIdToby,
		EmployerCompanyID: companyIdKlue,
		LocationCityID:    cityIdVancouver,
		Remote:            false,
		StartDate:         strVal("2017-04-24"),
		EndDate:           nil,
	},
	{
		// Teespring
		ID:                jobIdTeespring,
		EmployeePersonID:  personIdToby,
		EmployerCompanyID: companyIdTeespring,
		LocationCityID:    cityIdVancouver,
		Remote:            true,
		StartDate:         strVal("2016-01-03"),
		EndDate:           strVal("2016-12-21"),
	},
	{
		// Hootsuite (Full-time)
		ID:                jobIdHootsuite,
		EmployeePersonID:  personIdToby,
		EmployerCompanyID: companyIdHootsuite,
		LocationCityID:    cityIdVancouver,
		Remote:            false,
		StartDate:         strVal("2013-11-04"),
		EndDate:           strVal("2015-12-24"),
	},
	{
		// Hootsuite (Contract)
		ID:                jobIdHootsuiteContract,
		EmployeePersonID:  personIdToby,
		EmployerCompanyID: companyIdHootsuite,
		LocationCityID:    cityIdVancouver,
		Remote:            false,
		StartDate:         strVal("2013-01-07"),
		EndDate:           strVal("2013-04-05"),
	},
	{
		// Metalogix
		ID:                jobIdMetalogix,
		EmployeePersonID:  personIdToby,
		EmployerCompanyID: companyIdMetalogix,
		LocationCityID:    cityIdVancouver,
		Remote:            false,
		StartDate:         strVal("2008-05-01"),
		EndDate:           strVal("2012-12-24"),
	},
	{
		// ShipConstructor (term 1)
		ID:                jobIdShipConstructorTerm1,
		EmployeePersonID:  personIdToby,
		EmployerCompanyID: companyIdShipConstructor,
		LocationCityID:    cityIdVictoria,
		Remote:            false,
		StartDate:         strVal("2006-05-01"),
		EndDate:           strVal("2006-08-31"),
	},
	{
		// ShipConstructor (term 2)
		ID:                jobIdShipConstructorTerm2,
		EmployeePersonID:  personIdToby,
		EmployerCompanyID: companyIdShipConstructor,
		LocationCityID:    cityIdVictoria,
		Remote:            false,
		StartDate:         strVal("2007-05-01"),
		EndDate:           strVal("2007-08-31"),
	},
}

func findJob(id string) *job {
	for _, j := range jobs {
		if j.ID == id {
			return j
		}
	}

	return nil
}

func findJobsByPersonId(personId string) []*job {
	out := []*job{}

	for _, j := range jobs {
		if j.EmployeePersonID == personId {
			log.Println("Responding with job:", j.ID)
			out = append(out, j)
		}
	}

	return out
}
