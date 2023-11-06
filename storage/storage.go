package storage

import "app/models"

type StorageI interface {
	CloseDB()
	User() UserRepoI
	Product() ProductRepoI
	Shopcart() ShopcartRepoI
}

type UserRepoI interface {
	Create(*models.CreateUser) (string, error)
	GetPkey(*models.UserPrimaryKey) (*models.User, error)
	GetList(*models.GetListRequest) (*models.GetListResponse, error)
	Update(*models.UpdateUser) (string, error)
	Delete(*models.UserPrimaryKey) (int, error)
}

type ProductRepoI interface {
	Create(*models.CreateProduct) (string, error)
	GetPkey(*models.ProductPrimaryKey) (*models.Product, error)
	GetList(*models.GetListProductRequest) (*models.GetListProductResponse, error)
	Update(*models.UpdateProduct) (string, error)
	Delete(*models.ProductPrimaryKey) (int, error)
}

type ShopcartRepoI interface{
	AddShopcart(*models.AddShopcart) (*models.Shopcart, error)
	RemoveShopcart(*models.RemoveShopcart) (string, error)
}