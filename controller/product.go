package controller

import (
	"app/models"
	"app/pkg/util"
	"errors"
	"log"
)

func (c *Controller) CreateProduct(req *models.CreateProduct) (res string, err error) {
	res, err = c.store.Product().Create(req)
	if err != nil {
		log.Println("Calling Create method in Controller:", err)
		return
	}

	return res, nil
}

func (c *Controller) GetListProducts(req *models.GetListProductRequest) (res *models.GetListProductResponse, err error) {
	res, err = c.store.Product().GetList(req)
	if err != nil {
		log.Println("Calling Getlist method in Controller:", err)
		return
	}

	return res, nil
}

func (c *Controller) UpdateProduct(req *models.UpdateProduct) (res string, err error) {
	
	if !util.IsValidUUID(req.Id) {
		return req.Id, errors.New("invalid Id")
	}

	res, err = c.store.Product().Update(req)
	if err != nil {
		log.Println("Calling UpdateProduct method in Controller:", err)
		return
	}
	
	return res, nil
}

func (c *Controller) GetProductByPkey(req *models.ProductPrimaryKey) (res *models.Product, err error) {

	if !util.IsValidUUID(req.Id) {
		return nil, errors.New("Invalid ID")
	}

	res, err = c.store.Product().GetPkey(req)
	if err != nil {
		log.Println("Calling GetProductPkey in controller:", err)
		return
	}

	return res, nil
}

func (c *Controller) DeleteProduct(req *models.ProductPrimaryKey) (res int, err error) {

	if !util.IsValidUUID(req.Id) {
		return 1, errors.New("Invalid ID")
	}

	res, err = c.store.Product().Delete(req)
	if err != nil {
		log.Println("Calling DeleteProduct method in Controller:", err)
		return
	}

	return res, nil
}