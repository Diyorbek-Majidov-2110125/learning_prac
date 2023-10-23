package models



type User struct {
	Id      int
	Name    string
	Surname string
	Birthday string
}

type GetListRequest struct {
	Offset int
	Limit  int
	Search string
}

type GetListDate struct {
	FromDate string
	ToDate   string
}
