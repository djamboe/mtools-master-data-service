package models

type ProductModel struct {
	Id          string `bson:"_id"`
	ProductName string `json:"productname"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
