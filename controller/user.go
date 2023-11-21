package controller

import (
	"app/models"
	"app/pkg/util"
	"errors"
	"fmt"
	"time"
)

// "bufio"
// "errors"
// "fmt"
// "log"
// "os"
// "strings"
// "time"

// "github.com/bxcodec/faker/v3"

// var Users []models.User

func (c *Controller) CreateUser(req *models.CreateUser) (id string, err error) {

	id, err = c.store.User().Create(req)
	if err != nil {
		return "error:", err
	}

	return id, nil
}

func (c *Controller) GetListUsers(req *models.GetListRequest) (res *models.GetListResponse, err error) {

	res, err = c.store.User().GetList(req)
	if err != nil {
		return nil, err
	}
	return
}

func (c *Controller) GetByPkey(req *models.UserPrimaryKey) (res *models.User, err error) {

	if !util.IsValidUUID(req.Id) {
		return nil, errors.New("invalid ID")
	}

	res, err = c.store.User().GetPkey(req)
	if err != nil {
		return nil, err
	}
	return
}

func (c *Controller) Update(req *models.UpdateUser) (res string, err error) {

	if !util.IsValidUUID(req.Id) {
		return "", errors.New("invalid ID")
	}

	res, err = c.store.User().Update(req)
	if err != nil {
		return "error", err
	}
	return res, nil
}

func (c *Controller) Delete(req *models.UserPrimaryKey) (res int, err error) {

	if !util.IsValidUUID(req.Id) {
		return 1, errors.New("invalid ID")
	}

	res, err = c.store.User().Delete(req)
	if err != nil {
		return 1, nil
	}
	return res, nil
}

func (c *Controller) GetByName(req *models.GetListRequest) (res *[]models.User, err error) {

	res, err = c.store.User().GetByName(req)
	if err != nil {
		return nil, err
	}
	return
}

func (c *Controller) ChooseByBirthDate(req *models.GetListDate) (res []models.User, err error) {
	res, err = c.store.User().ChooseByBirthDate(req)
	if err != nil {
		return nil, err
	}
	return
}

func (c *Controller) WithdrawUserBalance(id *models.UserPrimaryKey, balance float64) error {

	user, err := c.store.User().GetPkey(&models.UserPrimaryKey{Id: id.Id})
	if err != nil {
		return err
	}

	if user.Balance < balance {
		return errors.New("not available balance")
	}

	user.Balance = user.Balance - balance
	_, err = c.store.Shopcart().UpdateStatus(&models.UpdateStatus{
		User_id: user.Id,
	})
	if err != nil {
		return err
	}

	_, err = c.store.User().Update(&models.UpdateUser{
		Id:       user.Id,
		Name:     user.Name,
		Surname:  user.Surname,
		Birthday: user.Birthday,
		Balance:  user.Balance,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) TransferBalance(req *models.TransferBalance) (res string, err error) {

	sender, err := c.store.User().GetPkey(&models.UserPrimaryKey{
		Id: req.SenderId,
	})
	if err != nil {
		return "failure", errors.New("sender not found")
	}
	receiver, err := c.store.User().GetPkey(&models.UserPrimaryKey{
		Id: req.ReceiverId,
	})
	if err != nil {
		return "failure", errors.New("receiver not found")
	}

	if sender.Balance < req.Money+req.Money*req.Service_fee_percentage/100 {
		return "not enough money", errors.New("not enough balance available")
	}

	_, err = c.store.User().Update(&models.UpdateUser{
		Id:      req.SenderId,
		Balance: sender.Balance - (req.Money + req.Service_fee_percentage*req.Money/100),
	})
	if err != nil {
		return "Sender Updated unsuccessfully", err
	}

	_, err = c.store.Commission().CreateCommission(&models.CreateCommission{
		SenderId:        req.SenderId,
		ReceiverId:      req.ReceiverId,
		Transaction_fee: req.Money * req.Service_fee_percentage / 100.0,
		Transacton_time: time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		fmt.Println("error", err)
		return "", err
	}

	_, err = c.store.User().Update(&models.UpdateUser{
		Id:      req.ReceiverId,
		Balance: receiver.Balance + req.Money,
	})
	if err != nil {
		return "Reciever Updated unsuccessfully", err
	}
	return "Success", nil
}

// func CreateUser(data models.User) {
// 	Users = append(Users, data)
// }

// func GetByUserId(id int) (res models.User, err bool) {
// 	for _, user := range Users {
// 		if id == user.Id {
// 			return user, false
// 		}
// 	}
// 	return models.User{}, true
// }

// func Delete(id int) []models.User {
// 	for ind, user := range Users {
// 		if id == user.Id {
// 			if ind < len(Users)-1 {
// 				Users = append(Users[:ind], Users[ind+1:]...)
// 			} else {
// 				Users = Users[:ind]
// 			}
// 		}
// 	}
// 	return Users
// }

// func Update(data models.User, users []models.User) []models.User {
// 	for ind, user := range users {
// 		if data.Id == user.Id {
// 			users[ind].Name = data.Name
// 			users[ind].Surname = data.Surname
// 			users[ind].Birthday = data.Birthday
// 			return users
// 		}
// 	}
// 	return []models.User{}
// }

// func GenerateUser(count int) []models.User {
// 	for i := 0; i < count; i++ {
// 		Users = append(Users, models.User{
// 			Id:       i + 1,
// 			Name:     faker.FirstName(),
// 			Surname:  faker.LastName(),
// 			Birthday: faker.Date(),
// 		})
// 	}
// 	return Users
// }

// func AllPages() int {
// 	if len(Users)%10 != 0 {
// 		return len(Users)/10 + 1
// 	}
// 	return len(Users) / 10
// }

// func GetPageByNumber(number int) []models.User {
// 	Users, err := GetListUsers(models.GetListRequest{
// 		Offset: (number - 1) * 10,
// 		Limit:  10,
// 	})

// 	if err != nil {
// 		log.Println("info: ", err)
// 	}

// 	return Users
// }

// func SortByDate(req models.GetListDate) ([]models.User, error) {
// 	users := []models.User{}

// 	fromDate, err1 := time.Parse("2006-01-02", req.FromDate)
// 	toDate, err2 := time.Parse("2006-01-02", req.ToDate)

// 	if err1 != nil || err2 != nil {
// 		return nil, fmt.Errorf("error in time parsing: %v, %v", err1, err2)
// 	}

// 	if fromDate.After(toDate) {
// 		return nil, fmt.Errorf("error: 'fromDate' is after 'toDate'")
// 	}

// 	for _, user := range Users {
// 		birthday, err := time.Parse("2006-01-02", user.Birthday)
// 		if err != nil {
// 			return nil, fmt.Errorf("error in parsing: %v", err)
// 		}
// 		if (birthday.After(fromDate) && birthday.Before(toDate)) || birthday.Equal(fromDate) || birthday.Equal(toDate) {
// 			users = append(users, user)
// 		}
// 	}

// 	return users, nil
// }

//update, delete, getbyid, generateuser, pages, scrolling pages
