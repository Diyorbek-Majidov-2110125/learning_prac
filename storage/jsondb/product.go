package jsondb

import (
	"app/models"
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/google/uuid"
)

type productRepo struct {
	fileName string
	file     *os.File
}

func NewProductRepo(fileName string, file *os.File) *productRepo {
	return &productRepo{
		fileName: fileName,
		file:     file,
	}
}

func (p *productRepo) Create(req *models.CreateProduct) (id string, err error) {

	var products []*models.Product

	id = uuid.New().String()
	err = json.NewDecoder(p.file).Decode(&products)
	if err != nil {
		log.Println("here")
		return "error:", err
	}

	products = append(products, &models.Product{
		Id:    id,
		Name:  req.Name,
		Price: req.Price,
	})

	body, err := json.MarshalIndent(products, " ", " ")
	if err != nil {
		log.Println("Marshalling:", err)
		return
	}

	err = os.WriteFile(p.fileName, body, os.ModePerm)
	if err != nil {
		log.Println("Writing File: ", err)
		return
	}

	return id, nil

}

func (p *productRepo) GetList(req *models.GetListProductRequest) (res *models.GetListProductResponse, err error) {

	fileContent, err := os.ReadFile(p.fileName)
	if err != nil {
		log.Println("Reading file: ", err)
		return
	}

	var products []*models.Product
	err = json.Unmarshal(fileContent, &products)
	if err != nil {
		log.Println("Unmarshalling file: ", err)
		return
	}

	if req.Offset > len(products) {
		return nil, errors.New("out of range")
	}

	if req.Offset + req.Limit > len(products) {
		return &models.GetListProductResponse{
			Products: products[req.Offset-1:],
		}, nil
	}

	return &models.GetListProductResponse{
		Products: products[req.Offset-1: req.Offset + req.Limit],
	}, nil

}

func (p *productRepo) GetPkey(req *models.ProductPrimaryKey) (res *models.Product, err error) {
	
	fileContent, err := os.ReadFile(p.fileName)
	if err != nil {
		return nil, err
	}

	var products []*models.Product
	err = json.Unmarshal(fileContent, &products)
	if err != nil {
		log.Println("unmarshalling file:", err)
		return
	}

	res = &models.Product{}
	
	for _, product := range products {
		if req.Id == product.Id {
			res.Id = product.Id
			res.Name = product.Name
			res.Price = product.Price
			break
		}
	}

	return res, nil
}

func (p *productRepo) Update(req *models.UpdateProduct) (res string, err error) {

	fileContent, err := os.ReadFile(p.fileName)
	if err != nil {
		log.Println("Reading File: ", err)
		return
	}

	var products []*models.Product
	err = json.Unmarshal(fileContent, &products)
	if err != nil {
		log.Println("Unmarshalling file:", err)
		return 
	}

	for ind, product := range products {
		if product.Id == req.Id {
			res = req.Id
			products[ind].Name = req.Name
			products[ind].Price = req.Price
			break
		}
	}

	updatedData, err := json.MarshalIndent(products, " ", " ")
	if err != nil {
		log.Println("Marshalling file:", err)
		return
	}

	err = os.WriteFile(p.fileName, updatedData, 0644)
	if err != nil {
		log.Println("Writing File:", err)
		return
	}

	return res, nil
}

func (p *productRepo) Delete(req *models.ProductPrimaryKey) (res int, err error) {

	fileContent, err := os.ReadFile(p.fileName)
	if err != nil {
		log.Println("Reading File: ", err)
		return
	}

	var products []*models.Product
	err = json.Unmarshal(fileContent, &products)
	if err != nil {
		log.Println("Unmarshalling file:", err)
		return
	}

	for ind, product := range products {
		if product.Id == req.Id {
			products = append(products[:ind], products[ind + 1:]...)
			break
		}
	}

	updatedData, err := json.MarshalIndent(products, " ", " ")
	if err != nil {
		log.Println("marshalling file:", err)
		return
	}

	err = os.WriteFile(p.fileName, updatedData, 0644)
	if err != nil {
		log.Println("writing file:", err)
		return
	}

	return 0, nil
}
