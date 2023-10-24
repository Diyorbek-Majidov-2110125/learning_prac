package storage

import (
	"app/models"
	"encoding/json"
	"log"
	"os"
)

type userRepo struct {
	fileName string
	file *os.File
}

func NewUserRepo(fileName string, file *os.File) *userRepo{
	return &userRepo{
		fileName: fileName,
		file: file,
	}
}

func (u *userRepo) Create(req *models.CreateUser) (id int, err error) {

	var users []*models.User
	err = json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return 0, err
	}

	if len(users) > 0 {
		id = users[len(users) - 1].Id + 1
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
