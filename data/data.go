package data

import (
	"net/url"
	"time"
)

const (
	dateFmt = "2006-01-02"
)

type Person struct {
	FirstName string
	MiddleName string
	LastName string
	Website *url.URL
	Jobs []*Job
}

type Company struct {
	Name string
}

type Job struct {
	Company *Company
	Roles []*Role
}

type Role struct {
	StartDate *time.Time
	EndDate *time.Time
	Title string
}

type Database struct {
	Me *Person
}

var (
	website, _ = url.Parse("https://tobysullivan.net")
	database = &Database{
		Me: &Person{
			FirstName: "Toby",
			MiddleName: "Jay",
			LastName: "Sullivan",
			Website: website,
			Jobs: []*Job{
				{
					Company: &Company{Name: "Klue"},
					Roles: []*Role{
						{
							Title: "Software Developer",
							StartDate: date("2017-04-24"),
							EndDate: nil,
						},
					},
				},
				{
					Company: &Company{Name: "Teespring"},
					Roles: []*Role{
						{
							Title: "Software Developer",
							StartDate: date("2016-01-03"),
							EndDate: date("2016-12-21"),
						},
					},
				},
				{
					Company: &Company{Name: "Hootsuite"},
					Roles: []*Role{
						{
							Title: "Software Developer (Contract)",
							StartDate: date("2013-01-01"),
							EndDate: date("2013-03-31"),
						},
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
					Company: &Company{Name: "Metalogix"},
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
							Title: "Software Developer (Freelance)",
							StartDate: date("2012-06-29"),
							EndDate: date("2012-12-24"),
						},
					},
				},
				{
					Company: &Company{Name: "ShipConstructor"},
					Roles: []*Role{
						{
							Title: "Quality Assurance Co-op",
							StartDate: date("2006-05-01"),
							EndDate: date("2006-08-31"),
						},
						{
							Title: "Software Developer Co-op",
							StartDate: date("2007-05-01"),
							EndDate: date("2007-08-31"),
						},
					},
				},
			},
		},
	}
)

func date(value string) *time.Time {
	t, err := time.Parse(dateFmt, value)
	if err != nil {
		return nil
	}
	return &t
}

func LoadDatabase() *Database {
	return database
}
