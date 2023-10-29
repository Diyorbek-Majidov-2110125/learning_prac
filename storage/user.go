package storage

import (
	"app/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type userRepo struct {
	fileName string
	file     *os.File
}

func NewUserRepo(fileName string, file *os.File) *userRepo {
	return &userRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *userRepo) Create(req *models.CreateUser) (id int, err error) {

	var users []*models.User
	err = json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return 0, err
	}

	if len(users) > 0 {
		id = users[len(users)-1].Id + 1
		users = append(users, &models.User{
			Id:       id,
			Name:     req.Name,
			Surname:  req.Surname,
			Birthday: req.Birthday,
		})
	} else {
		id = 1
		users = append(users, &models.User{
			Id:       id,
			Name:     req.Name,
			Surname:  req.Surname,
			Birthday: req.Birthday,
		})
	}

	body, err := json.MarshalIndent(users, " ", " ")
	if err != nil {
		log.Println(err)
		return
	}
	err = os.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		log.Println(err)
		return
	}
	return id, nil
}

func (u *userRepo) GetListUsers(req *models.GetListRequest) (res *models.GetListResponse, err error) {

	fileContent, err := os.ReadFile(u.fileName)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}

	var users []*models.User
	err = json.Unmarshal(fileContent, &users)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	if req.Offset > len(users) {
		return nil, errors.New("out of range")
	}

	if req.Offset+req.Limit > len(users) {
		return &models.GetListResponse{
			Users: users[req.Offset-1:],
		}, nil
	}

	return &models.GetListResponse{
		Users: users[req.Offset-1 : req.Offset+req.Limit],
	}, nil
}

func (u *userRepo) GetByPkey(req *models.UserPrimaryKey) (res *models.User, err error) {
	fileContent, err := os.ReadFile(u.fileName)
	if err != nil {
		return nil, err
	}

	var users *[]models.User
	err = json.Unmarshal(fileContent, &users)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}

	res = &models.User{}
	for _, user := range *users {
		if req.Id == user.Id {
			res.Id = user.Id
			res.Name = user.Name
			res.Surname = user.Surname
			res.Birthday = user.Birthday
			break
		}
	}
	return res, err
}

func (u userRepo) Update(id int, req *models.UpdateUser) (res int, err error) {

	fileContent, err := os.ReadFile(u.fileName)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	var users []*models.User
	err = json.Unmarshal(fileContent, &users)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	for ind, user := range users {
		if user.Id == id {
			res = id
			users[ind].Name = req.Name
			users[ind].Surname = req.Surname
			users[ind].Birthday = req.Birthday
			break
		}
	}

	updatedData, err := json.MarshalIndent(users," ", " ")
	if err != nil {
		return 1, err
	}

	err = os.WriteFile(u.fileName, updatedData, 0644)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	return res, nil
}

func (u *userRepo) Delete(req *models.UserPrimaryKey) (res int, err error) {

	fileContent, err := os.ReadFile(u.fileName)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	var users []*models.User
	err = json.Unmarshal(fileContent, &users)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	for ind, user := range users {
		if user.Id == req.Id {
			users = append(users[:ind], users[ind+1:]...)
			break
		}
	}

	updatedData, err := json.MarshalIndent(users," ", " ")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = os.WriteFile(u.fileName, updatedData, 0644)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	return 0, nil
}
