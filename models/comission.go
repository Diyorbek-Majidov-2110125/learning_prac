package models


type Commission struct {
	SenderId string `json:"sender_id"`
	ReceiverId string `json:"receiver_id"`
	Transaction_fee float64 `json:"transaction_fee"`
	Transacton_time string `json:"transaction_time"`
}

type GetCommission struct {
	SenderId string `json:"sender_id"`
	ReceiverId string `json:"receiver_id"`
}
type CreateCommission struct {
	SenderId string `json:"sender_id"`
	ReceiverId string `json:"receiver_id"`
	Transaction_fee float64 `json:"transaction_fee"`
	Transacton_time string `json:"transaction_time"`
}

type GetListCommissionRequest struct {
	Offset int `json:"offset"`
	Limit int `json:"limit"`
	Search string `json:"search"`
}

type GetListCommissionResponse struct {
	Count int `json:"count"`
	Commission []*Commission `json:"commissions"`
}
