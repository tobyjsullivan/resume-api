package main

type role struct {
	ID        string  `json:"id"`
	JobID     string  `json:"jobId"`
	Title     string  `json:"title"`
	StartDate *string `json:"startDate"`
	EndDate   *string `json:"endDate"`
}

var roles = []*role{
	{
		ID:        "5b950ec5-9e27-400d-a593-ca129dd4afa8",
		JobID:     jobIdKlue,
		Title:     "Software Developer",
		StartDate: strVal("2017-04-24"),
		EndDate:   nil,
	},
	{
		ID:        "d2c9446a-b916-4e55-9136-cca8f534cf7a",
		JobID:     jobIdTeespring,
		Title:     "Software Developer",
		StartDate: strVal("2016-01-03"),
		EndDate:   strVal("2016-12-21"),
	},
	{
		ID:        "c3247f3b-beb9-48da-8de6-8da0664c41b1",
		JobID:     jobIdHootsuite,
		Title:     "Software Developer",
		StartDate: strVal("2013-11-04"),
		EndDate:   nil,
	},
	{
		ID:        "a9e65d32-b9d9-4ecc-8535-e5b52d0d5dca",
		JobID:     jobIdHootsuite,
		Title:     "Lead Software Developer",
		StartDate: nil,
		EndDate:   strVal("2015-12-24"),
	},
	{
		ID:        "24577895-8ea3-4aa0-aebd-cf7fa38b8b75",
		JobID:     jobIdHootsuiteContract,
		Title:     "Software Developer (Contract)",
		StartDate: strVal("2013-01-07"),
		EndDate:   strVal("2013-04-05"),
	},
	{
		ID:        "8eee03ca-29a0-471d-b261-1913fa8da4f9",
		JobID:     jobIdMetalogix,
		Title:     "Software Developer Co-op",
		StartDate: strVal("2008-05-01"),
		EndDate:   strVal("2008-12-24"),
	},
	{
		ID:        "12e0c977-89c5-446a-a81a-c02061c933c7",
		JobID:     jobIdMetalogix,
		Title:     "Software Developer",
		StartDate: strVal("2008-12-24"),
		EndDate:   nil,
	},
	{
		ID:        "59bb93b1-c3ed-475b-8e14-ca98e0625763",
		JobID:     jobIdMetalogix,
		Title:     "Lead Software Developer",
		StartDate: nil,
		EndDate:   strVal("2012-06-29"),
	},
	{
		ID:        "9d95e0c5-2e69-4960-a555-2f13b24ff71a",
		JobID:     jobIdMetalogix,
		Title:     "Software Developer (Contract)",
		StartDate: strVal("2012-06-29"),
		EndDate:   strVal("2012-12-24"),
	},
	{
		ID:        "5582b637-9055-4c4a-80cf-aeca0c901088",
		JobID:     jobIdShipConstructorTerm1,
		Title:     "Quality Assurance Co-op",
		StartDate: strVal("2006-05-01"),
		EndDate:   strVal("2006-08-31"),
	},
	{
		ID:        "be638494-5264-4502-a289-6dca87eb328b",
		JobID:     jobIdShipConstructorTerm2,
		Title:     "Software Developer Co-op",
		StartDate: strVal("2007-05-01"),
		EndDate:   strVal("2007-08-31"),
	},
}

func findRole(id string) *role {
	for _, r := range roles {
		if r.ID == id {
			return r
		}
	}

	return nil
}

func findRolesByJobId(jobId string) []*role {
	out := []*role{}

	for _, r := range roles {
		if r.JobID == jobId {
			out = append(out, r)
		}
	}

	return out
}
