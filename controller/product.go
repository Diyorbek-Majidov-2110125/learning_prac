package controller

import (
	"app/models"
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
	
	res, err = c.store.Product().Update(req)
	if err != nil {
		log.Println("Calling UpdateProduct method in Controller:", err)
		return
	}
	
	return res, nil
}

func (c *Controller) GetProductByPkey(req *models.ProductPrimaryKey) (res *models.Product, err error) {
	res, err = c.store.Product().GetPkey(req)
	if err != nil {
		log.Println("Calling GetProductPkey in controller:", err)
		return
	}

	return res, nil
}

func (c *Controller) DeleteProduct(req *models.ProductPrimaryKey) (res int, err error) {
	res, err = c.store.Product().Delete(req)
	if err != nil {
		log.Println("Calling DeleteProduct method in Controller:", err)
		return
	}

	return res, nil
}