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

func (c *Controller) CalcTotal(req *models.CalculateShop) (res float64, err error) {

	var userShopcarts []models.Shopcart
	var total float64

	userShopcarts, err = c.store.Shopcart().GetUserShopcart(req)
	if err != nil {
		return 1, err
	}
	for _, val := range userShopcarts {
		product, err := c.store.Product().GetPkey(&models.ProductPrimaryKey{Id: val.Product_id})
		if err != nil {
			return 1, err
		}
		total += float64(val.Count) * product.Price
	}

	if req.DiscountStatus == "fixed" && req.Discount > total {
		total = total - req.Discount
	} else if req.DiscountStatus == "percentage"{
		if req.Discount > 0 && req.Discount <= 100 {
			total -= total * req.Discount/100
		} else {
			return 0, errors.New("out of range in percentage")
		}
		
	} else {
		return 0, errors.New("not allowed discount")
	}

	return total, nil
}

func (c *Controller) UpdateStatus(req *models.UpdateStatus) (status string, err error) {
	status, err = c.store.Shopcart().UpdateStatus(req)
	if err != nil {
		return "1", err
	}

	return 
}
