package jsondb

import (
	"app/models"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
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

func (u *userRepo) Create(req *models.CreateUser) (id string, err error) {

	var users []*models.User

	id = uuid.New().String()
	err = json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return "Error:", err
	}

	users = append(users, &models.User{
		Id:       id,
		Name:     req.Name,
		Surname:  req.Surname,
		Birthday: req.Birthday,
		Balance:  req.Balance,
	})

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

func (u *userRepo) GetList(req *models.GetListRequest) (res *models.GetListResponse, err error) {

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

func (u *userRepo) GetPkey(req *models.UserPrimaryKey) (res *models.User, err error) {
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

	found := false
	res = &models.User{}
	for _, user := range *users {
		if req.Id == user.Id {
			res.Id = user.Id
			res.Name = user.Name
			res.Surname = user.Surname
			res.Birthday = user.Birthday
			res.Balance = user.Balance
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("user with ID %s not found", req.Id)
	}

	return res, nil
}


func (u userRepo) Update(req *models.UpdateUser) (res string, err error) {

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
			res = req.Id
			users[ind].Name = req.Name
			users[ind].Surname = req.Surname
			users[ind].Birthday = req.Birthday
			users[ind].Balance = req.Balance
			break
		}
	}

	updatedData, err := json.MarshalIndent(users, " ", " ")
	if err != nil {
		return "error:", err
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

	updatedData, err := json.MarshalIndent(users, " ", " ")
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

func (u *userRepo) GetByName(req *models.GetListRequest) (res *[]models.User, err error) {

	var users *[]models.User
	filContent, err := os.ReadFile(u.fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(filContent, &users)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your full name: ")

	fullName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	req.Search = strings.TrimSpace(fullName)

	results := []models.User{}
	for _, user := range *users {
		if strings.Contains(strings.ToLower(user.Name)+" "+strings.ToLower(user.Surname), strings.ToLower(req.Search)) || strings.Contains(strings.ToLower(user.Surname)+" "+strings.ToLower(user.Name), strings.ToLower(req.Search)) {
			results = append(results, user)
			continue
		}
		if strings.Contains(strings.ToLower(user.Name), strings.ToLower(req.Search)) || strings.Contains(strings.ToLower(user.Surname), strings.ToLower(req.Search)) {
			results = append(results, user)
		}
	}

	return &results, nil
}

func (u *userRepo) ChooseByBirthDate(req *models.GetListDate) (res []models.User, err error) {

	var users *[]models.User
	filContent, err := os.ReadFile(u.fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(filContent, &users)
	if err != nil {
		return nil, err
	}

	fromDate, err1 := time.Parse("2006-01-02", req.FromDate)
	toDate, err2 := time.Parse("2006-01-02", req.ToDate)

	if err1 != nil || err2 != nil {
		return nil, fmt.Errorf("error in time parsing: %v, %v", err1, err2)
	}

	if fromDate.After(toDate) {
		return nil, fmt.Errorf("error: 'fromDate' is after 'toDate'")
	}

	for _, user := range *users {
		birthday, err := time.Parse("2006-01-02", user.Birthday)
		if err != nil {
			return nil, fmt.Errorf("error in parsing: %v", err)
		}
		if (birthday.After(fromDate) && birthday.Before(toDate)) || birthday.Equal(fromDate) || birthday.Equal(toDate) {
			res = append(res, user)
		}
	}

	return res, nil
}


