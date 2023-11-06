package controller

import "app/models"

func (c *Controller) AddShopcart(req *models.AddShopcart) (res *models.Shopcart, err error) {

	res, err = c.store.Shopcart().AddShopcart(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Controller)  RemoveShopcart(req *models.RemoveShopcart) (res string, err error) {
	res, err = c.store.Shopcart().RemoveShopcart(req)
	if err != nil {
		return "ERROR", err
	}
	return res, nil
}