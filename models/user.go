package models


type UserPrimaryKey struct {
	Id string `json:"id"`
}

type CreateUser struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Birthday string `json:"birthday"`
	Balance float64 `json:"balance"`
}

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Birthday string `json:"birthday"`
	Balance float64 `json:"balance"`
}

type UpdateUser struct {
	Id 		string  `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Birthday string `json:"birthday"`
	Balance float64 `json:"balance"`
}

type GetListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Search string `json:"search"`
}

type GetListResponse struct {
	Count int `json:"count"`
	Users []*User `json:"users"`
}

type GetListDate struct {
	FromDate string `json:"fromdate"`
	ToDate   string `json:"todate"`
}

type TransferBalance struct {
	SenderId string `json:"sender_id"`
	ReceiverId string `json:"receiver_id"`
	Money float64 `json:"money"`
	Service_fee_percentage float64 `json:"service_fee_percentage"`
}
