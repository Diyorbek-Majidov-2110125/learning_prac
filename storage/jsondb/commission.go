package jsondb

import (
	"app/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type commissionRepo struct {
	fileName string
	file     *os.File
}

func NewCommissionRepo(fileName string, file *os.File) *commissionRepo {
	return &commissionRepo{
		fileName: fileName,
		file:     file,
	}
}

func (c *commissionRepo) CreateCommission(req *models.CreateCommission) (res *models.Commission, err error) {
	var commissions []*models.Commission

	err = json.NewDecoder(c.file).Decode(&commissions)
	if err != nil {
		return nil, err
	}
	res = &models.Commission{
		SenderId:        req.SenderId,
		ReceiverId:      req.ReceiverId,
		Transaction_fee: req.Transaction_fee,
		Transacton_time: req.Transacton_time,
	}
	commissions = append(commissions, res)

	body, err := json.MarshalIndent(commissions, " ", " ")
	if err != nil {
		log.Println("Marshalling:", err)
		return nil, err
	}

	err = os.WriteFile(c.fileName, body, os.ModePerm)
	if err != nil {
		fmt.Println("Writing File: ", err)
		return nil, err
	}

	return res, nil
}

func (c *commissionRepo) GetCommission(req *models.GetCommission) (res *models.Commission, err error) {
	var commissions *[]models.Commission
	fileContent, err := os.ReadFile(c.fileName)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fileContent, &commissions)
	if err != nil {
		return nil, err
	}

	for _, val := range *commissions {
		if req.SenderId == val.SenderId && res.ReceiverId == val.ReceiverId {
			res.SenderId = val.SenderId
			res.ReceiverId = val.ReceiverId
			res.Transaction_fee = val.Transaction_fee
			res.Transacton_time = val.Transacton_time
			break
		}
	}
	return res, nil
}

func (c *commissionRepo) GetCommissionList(req *models.GetListCommissionRequest) (res *models.GetListCommissionResponse, err error) {
	fileContent, err := os.ReadFile(c.fileName)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}

	var commissions []*models.Commission
	err = json.Unmarshal(fileContent, &commissions)
	if err != nil {
		fmt.Println("error: ", err)
		return nil, err
	}

	if req.Offset > len(commissions) {
		return nil, errors.New("out of range")
	}

	if req.Offset+req.Limit > len(commissions) {
		return &models.GetListCommissionResponse{
			Commission: commissions[req.Offset-1:],
		}, nil
	}

	return &models.GetListCommissionResponse{
		Commission: commissions[req.Offset-1 : req.Offset+req.Limit],
	}, nil
}
