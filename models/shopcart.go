package models

type ShopcartPrimaryKey struct{
	Id string `json:"id"`
}

type Shopcart struct {
	Id string `json:"id"`
	User_id string `json:"user_id"`
	Product_id string `json:"product_id"`
	Count int `json:"count"`
	IsPaid bool `json:"is_paid"`
}

type AddShopcart struct {
	User_id string `json:"user_id"`
	Product_id string `json:"product_id"`
	Count int `json:"count"`
	IsPaid bool `json:"is_paid"`
}

type RemoveShopcart struct {
	User_id string `json:"user_id"`
	Product_id string `json:"product_id"`
}

type CalculateShop struct {
	UserId string 	`json:"userId"`
	Discount float64 `json:"discount"`
	DiscountStatus string `json:"discount_status"`
}

type UpdateStatus struct {
	User_id string `json:"user_id"`
	Product_id string `json:"product_id"`
	Count int `json:"count"`
	IsPaid bool `json:"is_paid"`
}



