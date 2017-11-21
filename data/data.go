package data

import (
	"net/url"
	"time"
	"github.com/satori/go.uuid"
)

const (
	dateFmt = "2006-01-02"
)

type Person struct {
	ID uuid.UUID
	FirstName string
	MiddleName string
	LastName string
	Website *url.URL
	JobHistory []*Job
}

type Company struct {
	ID uuid.UUID
	Name string
	City *City
}

type City struct {
	ID uuid.UUID
	Name string
	Province *Province
}

type Province struct {
	Name string
	Country *Country
}

type Country struct {
	Name string
}

type Job struct {
	Company *Company
	StartDate *time.Time
	EndDate *time.Time
	City *City
	Remote bool
	Roles []*Role
}

type Role struct {
	StartDate *time.Time
	EndDate *time.Time
	Title string
}

var (
	canada          = &Country{"Canada"}
	usa             = &Country{"USA"}
	britishColumbia = &Province{
		Name:    "British Columbia",
		Country: canada,
	}
	california = &Province{
		Name:    "California",
		Country: usa,
	}
	vancouver = &City{
		ID: uuid.NewV4(),
		Name: "Vancouver",
		Province: britishColumbia,
	}
	victoria = &City{
		ID: uuid.NewV4(),
		Name: "Victoria",
		Province: britishColumbia,
	}
	sanFrancisco = &City{
		ID: uuid.NewV4(),
		Name: "San Francisco",
		Province: california,
	}
	cities = []*City{
		vancouver,
		victoria,
		sanFrancisco,
	}
	klue = &Company{
		ID: uuid.NewV4(),
		Name: "Klue",
		City: vancouver,
	}
	teespring = &Company{
		ID: uuid.NewV4(),
		Name: "Teespring",
		City: sanFrancisco,
	}
	hootsuite = &Company{
		ID: uuid.NewV4(),
		Name: "Hootsuite",
		City: vancouver,
	}
	metalogix = &Company{
		ID: uuid.NewV4(),
		Name: "Metalogix",
		City: vancouver,
	}
	shipConstructor = &Company{
		ID: uuid.NewV4(),
		Name: "ShipConstructor",
		City: victoria,
	}
	companies = []*Company{
		klue,
		teespring,
		hootsuite,
		metalogix,
		shipConstructor,
	}

	website, _ = url.Parse("https://tobysullivan.net")

	me = &Person{
		ID: uuid.NewV4(),
		FirstName:  "Toby",
		MiddleName: "Jay",
		LastName:   "Sullivan",
		Website:    website,
		JobHistory: []*Job{
			{
				Company: klue,
				StartDate: date("2017-04-24"),
				EndDate: nil,
				City: vancouver,
				Roles: []*Role{
					{
						Title: "Software Developer",
						StartDate: date("2017-04-24"),
						EndDate: nil,
					},
				},
			},
			{
				Company: teespring,
				StartDate: date("2016-01-03"),
				EndDate: date("2016-12-21"),
				Remote: true,
				Roles: []*Role{
					{
						Title: "Software Developer",
						StartDate: date("2016-01-03"),
						EndDate: date("2016-12-21"),
					},
				},
			},
			{
				Company: hootsuite,
				StartDate: date("2013-11-04"),
				EndDate: date("2015-12-24"),
				City: vancouver,
				Roles: []*Role{
					{
						Title: "Software Developer",
						StartDate: date("2013-11-04"),
						EndDate: nil,
					},
					{
						Title: "Lead Software Developer",
						StartDate: nil,
						EndDate: date("2015-12-24"),
					},
				},
			},
			{
				Company: hootsuite,
				StartDate: date("2013-01-01"),
				EndDate: date("2013-03-31"),
				City: vancouver,
				Roles: []*Role{
					{
						Title: "Software Developer (Contract)",
						StartDate: date("2013-01-01"),
						EndDate: date("2013-03-31"),
					},
				},
			},
			{
				Company: metalogix,
				StartDate: date("2008-05-01"),
				EndDate: date("2012-12-24"),
				City: vancouver,
				Roles: []*Role{
					{
						Title: "Software Developer Co-op",
						StartDate: date("2008-05-01"),
						EndDate: date("2008-12-24"),
					},
					{
						Title: "Software Developer",
						StartDate: date("2008-12-24"),
						EndDate: nil,
					},
					{
						Title: "Lead Software Developer",
						StartDate: nil,
						EndDate: date("2012-06-29"),
					},
					{
						Title: "Software Developer (Contract)",
						StartDate: date("2012-06-29"),
						EndDate: date("2012-12-24"),
					},
				},
			},
			{
				Company: shipConstructor,
				StartDate: date("2006-05-01"),
				EndDate:   date("2006-08-31"),
				City: victoria,
				Roles: []*Role{
					{
						Title:     "Quality Assurance Co-op",
						StartDate: date("2006-05-01"),
						EndDate:   date("2006-08-31"),
					},
				},
			},
			{
				Company: shipConstructor,
				StartDate: date("2007-05-01"),
				EndDate: date("2007-08-31"),
				City: victoria,
				Roles: []*Role{
					{
						Title: "Software Developer Co-op",
						StartDate: date("2007-05-01"),
						EndDate: date("2007-08-31"),
					},
				},
			},
		},
	}
	people = []*Person{
		me,
	}
)

func date(value string) *time.Time {
	t, err := time.Parse(dateFmt, value)
	if err != nil {
		return nil
	}
	return &t
}

func Me()*Person {
	return me
}

func FindCompany(id uuid.UUID) *Company {
	for _, c := range companies {
		if c.ID == id {
			return c;
		}
	}

	return nil;
}

func FindCity(id uuid.UUID) *City {
	for _, c := range cities {
		if c.ID == id {
			return c;
		}
	}

	return nil;
}

func FindPerson(id uuid.UUID) *Person {
	for _, p := range people {
		if p.ID == id {
			return p;
		}
	}

	return nil;
}
