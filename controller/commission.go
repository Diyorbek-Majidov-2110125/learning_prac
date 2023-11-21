package controller

import "app/models"

func (c *Controller) CreateCommission(req *models.CreateCommission) (res *models.Commission, err error) {
	res, err = c.store.Commission().CreateCommission(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Controller) GetCommission(req *models.GetCommission)(res *models.Commission, err error) {
	res, err = c.store.Commission().GetCommission(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Controller) GetCommissionList(req *models.GetListCommissionRequest) (res *models.GetListCommissionResponse, err error) {
	res, err = c.store.Commission().GetCommissionList(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Controller) CalculateTotalTransactionFee(req *models.GetListCommissionRequest) (res float64, err error) {
	commissions, err := c.store.Commission().GetCommissionList(req)
	if err != nil {
		return 1, err
	}
	for _, com := range commissions.Commission {
		res += com.Transaction_fee
	}
	return res, nil
}