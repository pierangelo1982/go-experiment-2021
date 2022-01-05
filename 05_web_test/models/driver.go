package models

type Driver struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	CarNumber int    `json:"carNumber"`
	TeamId    int    `son:"teamId"`
}
