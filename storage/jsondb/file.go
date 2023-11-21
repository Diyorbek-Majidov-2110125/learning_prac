package jsondb

import (
	"app/config"
	"app/storage"
	"os"
)

type Store struct {
	user *userRepo
	product *productRepo
	shopcart *shopcartRepo
	commission *commissionRepo
}

func NewFileJson(cfg *config.Config) (storage.StorageI, error) {


	userFile, err := os.Open(cfg.Path + cfg.UserFileName)
	if err != nil {
		return nil, err
	}

	productFile, err := os.Open(cfg.Path + cfg.ProductFileName)
	if err != nil {
		return nil, err
	}

	shopcartFile, err := os.Open(cfg.Path + cfg.ShopcartFileName)
	if err != nil {
		return nil, err
	}

	commissionFile, err := os.Open(cfg.Path + cfg.CommissionFileName)
	if err != nil {
		return nil, err
	}

	return &Store{
		user: NewUserRepo(cfg.Path+cfg.UserFileName, userFile),
		product: NewProductRepo(cfg.Path + cfg.ProductFileName, productFile),
		shopcart: NewShopcartRepo(cfg.Path + cfg.ShopcartFileName, shopcartFile),
		commission: NewCommissionRepo(cfg.Path + cfg.CommissionFileName, commissionFile),
	}, nil
}

func (s *Store) CloseDB() {
	s.user.file.Close()
}

func (s *Store) User() storage.UserRepoI {
	return s.user
}

func (s *Store) Product() storage.ProductRepoI {
	return s.product
}

func (s *Store) Shopcart() storage.ShopcartRepoI {
	return s.shopcart
}

func (s *Store) Commission() storage.CommissionRepoI {
	return s.commission
}


