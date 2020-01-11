package models

type CustomerModel struct {
	Id              string `json:"_id"`
	CustomerName    string `json:"customername"`
	CustomerEmail   string `json:"customeremail"`
	CustomerPhone   string `json:"customerphone"`
	CustomerAddress string `json:"customeraddress"`
	Status          string `json:"status"`
}
