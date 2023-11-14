package jsondb

import (
	"app/models"
	"encoding/json"
	"os"

	"github.com/google/uuid"
)

type shopcartRepo struct {
	fileName string
	file     *os.File
}

func NewShopcartRepo(fileName string, file *os.File) *shopcartRepo {
	return &shopcartRepo{
		fileName: fileName,
		file:     file,
	}
}

func (sh *shopcartRepo) AddShopcart(req *models.AddShopcart) (res *models.Shopcart, err error) {

	var shopcartList []*models.Shopcart

	err = json.NewDecoder(sh.file).Decode(&shopcartList)
	if err != nil {
		return nil, err
	}

	id := uuid.New().String()
	res = &models.Shopcart{
		Id:         id,
		User_id:    req.User_id,
		Product_id: req.Product_id,
		Count:      req.Count,
		IsPaid:     false,
	}

	shopcartList = append(shopcartList, res)

	body, err := json.MarshalIndent(shopcartList, " ", " ")
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(sh.fileName, body, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (sh *shopcartRepo) RemoveShopcart(req *models.RemoveShopcart) (res string, err error) {

	fileContent, err := os.ReadFile(sh.fileName)
	if err != nil {
		return "ERROR:", err
	}

	var shopcartList []*models.Shopcart
	err = json.Unmarshal(fileContent, &shopcartList)
	if err != nil {
		return "ERROR:", err
	}

	for ind, info := range shopcartList {
		if info.User_id == req.User_id && info.Product_id == req.Product_id {
			shopcartList = append(shopcartList[:ind], shopcartList[ind+1:]...)
			break
		}
	}

	updatedData, err := json.MarshalIndent(shopcartList, " ", " ")
	if err != nil {
		return "ERROR:", err
	}

	err = os.WriteFile(sh.fileName, updatedData, 0644)
	if err != nil {
		return "ERROR:", err
	}

	return "Successfully deleted", nil
}

func (sh *shopcartRepo) GetbyId(req *models.ShopcartPrimaryKey) (res *models.Shopcart, err error) {

	var shopcart models.Shopcart

	fileContent, err := os.ReadFile(sh.fileName)
	if err != nil {
		return nil, err
	}

	var list *[]models.Shopcart
	err = json.Unmarshal(fileContent, &list)
	if err != nil {
		return nil, err
	}

	for _, val := range *list {
		if req.Id == val.Id {
			shopcart.Id = val.Id
			shopcart.Product_id = val.Product_id
			shopcart.User_id = val.User_id
			shopcart.Count = val.Count
		}
	}

	res = &shopcart

	return res, nil
}

func (sh *shopcartRepo) GetUserShopcart(req *models.CalculateShop) (res []models.Shopcart, err error) {

	var shopcarts []models.Shopcart
	fileContent, err := os.ReadFile(sh.fileName)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fileContent, &shopcarts)
	if err != nil {
		return nil, err
	}

	var result []models.Shopcart
	for _, shopcart := range shopcarts {
		if shopcart.User_id == req.UserId && shopcart.IsPaid == false {
			result = append(result, shopcart)
		}
	}

	res = result
	return res, nil
}

func (sh *shopcartRepo) UpdateStatus(req *models.UpdateStatus) (res string, err error) {
	var shopcarts []models.Shopcart
	fileContent, err := os.ReadFile(sh.fileName)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(fileContent, &shopcarts)
	if err != nil {
		return "", err
	}

	for ind, shopcart := range shopcarts {
		if shopcart.User_id == req.User_id{
			shopcarts[ind].IsPaid = true
		}
	}

	updatedData, err := json.MarshalIndent(shopcarts, " ", " ")
	if err != nil {
		return "Fail", err
	}

	err = os.WriteFile(sh.fileName, updatedData, os.ModePerm)
	if err != nil {
		return "Fail", err
	}
	return "Success", nil
}
