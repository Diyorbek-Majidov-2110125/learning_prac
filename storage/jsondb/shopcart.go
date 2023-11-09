package jsondb

import (
	"app/models"
	"encoding/json"
	"os"

	"github.com/google/uuid"
)

type shopcartRepo struct {
	fileName string
	file *os.File
}

func NewShopcartRepo(fileName string, file *os.File) *shopcartRepo {
	return &shopcartRepo{
		fileName: fileName,
		file: file,
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
		Id: id,
		User_id: req.User_id,
		Product_id: req.Product_id,
		Count: req.Count,
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
			shopcartList = append(shopcartList[:ind], shopcartList[ind + 1:]...)
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

// func (sh *shopcartRepo) GetShopcartById(req *models.ShopcartPrimaryKey) (res *models.Shopcart, err error) {

// }
