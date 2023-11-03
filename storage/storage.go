package storage

import "app/models"

type StorageI interface {
	CloseDB()
	User() UserRepoI
}

type UserRepoI interface {
	Create(*models.CreateUser) (string, error)
	GetPkey(*models.UserPrimaryKey) (*models.User, error)
	GetList(*models.GetListRequest) (*models.GetListResponse, error)
	Update(*models.UpdateUser) (string, error)
	Delete(*models.UserPrimaryKey) (int, error)
}
