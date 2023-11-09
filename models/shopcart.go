package models

type ShopcartPrimaryKey struct{
	Id string `json:"id"`
}

type Shopcart struct {
	Id string `json:"id"`
	User_id string `json:"user_id"`
	Product_id string `json:"product_id"`
	Count int `json:"count"`
}

type AddShopcart struct {
	User_id string `json:"user_id"`
	Product_id string `json:"product_id"`
	Count int `json:"count"`
}

type RemoveShopcart struct {
	User_id string `json:"user_id"`
	Product_id string `json:"product_id"`
}



