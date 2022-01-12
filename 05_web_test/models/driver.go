package models

type Team struct {
	//Id      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Driver struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	CarNumber int    `json:"carNumber"`
	TeamId    int    `json:"teamId"`
	Owner     Team   `json:"owner"`
}
