package controller

import (
	"app/models"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bxcodec/faker/v3"
)

var Users []models.User

func CreateUser(data models.User) {
	Users = append(Users, data)
}

func GetListUsers(req models.GetListRequest) (res []models.User, err bool) {

	if req.Offset > len(Users) {
		return []models.User{}, true
	}

	if req.Offset+req.Limit > len(Users) {
		return Users[req.Offset:], false
	}

	return Users[req.Offset : req.Offset+req.Limit], false
}

func GetByUserId(id int) (res models.User, err bool) {
	for _, user := range Users {
		if id == user.Id {
			return user, false
		}
	}
	return models.User{}, true
}

func Delete(id int) []models.User {
	for ind, user := range Users {
		if id == user.Id {
			if ind < len(Users)-1 {
				Users = append(Users[:ind], Users[ind+1:]...)
			} else {
				Users = Users[:ind]
			}
		}
	}
	return Users
}

func Update(data models.User, users []models.User) []models.User {
	for ind, user := range users {
		if data.Id == user.Id {
			users[ind].Name = data.Name
			users[ind].Surname = data.Surname
			users[ind].Birthday = data.Birthday
			return users
		}
	}
	return []models.User{}
}

func GenerateUser(count int) []models.User {
	for i := 0; i < count; i++ {
		Users = append(Users, models.User{
			Id:       i + 1,
			Name:     faker.FirstName(),
			Surname:  faker.LastName(),
			Birthday: faker.Date(),
		})
	}
	return Users
}

func AllPages() int {
	if len(Users)%10 != 0 {
		return len(Users)/10 + 1
	}
	return len(Users) / 10
}

func GetPageByNumber(number int) []models.User {
	Users, err := GetListUsers(models.GetListRequest{
		Offset: (number - 1) * 10,
		Limit:  10,
	})

	if err {
		panic(err)
	}

	return Users
}

func GetByName(users []models.User, req models.GetListRequest) []models.User {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your full name: ")

	fullName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	req.Search = strings.TrimSpace(fullName)

	results := []models.User{}
	for _, user := range users {
		if strings.Contains(strings.ToLower(user.Name)+" "+strings.ToLower(user.Surname), strings.ToLower(req.Search)) || strings.Contains(strings.ToLower(user.Surname)+" "+strings.ToLower(user.Name), strings.ToLower(req.Search)) {
			results = append(results, user)
			continue
		}
		if strings.Contains(strings.ToLower(user.Name), strings.ToLower(req.Search)) || strings.Contains(strings.ToLower(user.Surname), strings.ToLower(req.Search)) {
			results = append(results, user)
		}
	}
	return results
}
func SortByDate(req models.GetListDate) ([]models.User, error) {
	users := []models.User{}

	fromDate, err1 := time.Parse("2006-01-02", req.FromDate)
	toDate, err2 := time.Parse("2006-01-02", req.ToDate)

	if err1 != nil || err2 != nil {
		return nil, fmt.Errorf("Error in time parsing: %v, %v", err1, err2)
	}

	if fromDate.After(toDate) {
		return nil, fmt.Errorf("Error: 'fromDate' is after 'toDate'")
	}

	for _, user := range Users {
		birthday, err := time.Parse("2006-01-02", user.Birthday)
		if err != nil {
			return nil, fmt.Errorf("Error in parsing: %v", err)
		}
		if (birthday.After(fromDate) && birthday.Before(toDate)) || birthday.Equal(fromDate) || birthday.Equal(toDate) {
			users = append(users, user)
		}
	}

	return users, nil
}


//update, delete, getbyid, generateuser, pages, scrolling pages
