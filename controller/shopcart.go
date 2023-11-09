package controller

import (
	"app/models"
	"app/pkg/util"
	"errors"
)

func (c *Controller) AddShopcart(req *models.AddShopcart) (res *models.Shopcart, err error) {

	res, err = c.store.Shopcart().AddShopcart(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Controller) RemoveShopcart(req *models.RemoveShopcart) (res string, err error) {

	if !util.IsValidUUID(req.Product_id) || !util.IsValidUUID(req.User_id) {
		return "", errors.New("invalid Id")
	}

	res, err = c.store.Shopcart().RemoveShopcart(req)
	if err != nil {
		return "ERROR", err
	}
	return res, nil
}
