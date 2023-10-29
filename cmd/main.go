package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage"
	"fmt"
	"log"
)

func main() {

	cfg := config.Load()

	store, err := storage.NewFileJson(&cfg)
	if err != nil {
		panic("error while connect to json file: " + err.Error())
	}

	c := controller.NewController(&cfg, store)

//Create User:::::>
	// id, err := c.CreateUser(
	// 	&models.CreateUser{
	// 		Name: "Diyorbek",
	// 		Surname: "Majidov",
	// 		Birthday: "2001-10-30",
	// 	},
	// )
	// if err != nil {
	// 	log.Println("error while CreateUser:", err.Error())
	// 	return
	// }

	// fmt.Println(id)

//GetlistUsers::::>
	// users, err := c.GetListUsers(&models.GetListRequest{
	// 	Offset: 2,
	// 	Limit: 8,
	// 	Search: "",
	// })
	// if err != nil {
	// 	log.Println("error is coming from GetlistUsers...")
	// 	return
	// }

	// for _, user := range users.Users {
	// 	fmt.Println(user.Id,user.Name,user.Surname, user.Birthday)
	// }

//GetUserById::::>
	// user, err := c.GetByPkey(&models.UserPrimaryKey{
	// 	Id: 5,
	// })
	// if err != nil {
	// 	log.Println("error: ", err)
	// 	return
	// }
	// fmt.Println(*user)

//Update::::>

		// id , err := c.Update(4, &models.UpdateUser{
		// 	Name: "Jasurbek",
		// 	Surname: "Abdullaev",
		// 	Birthday: "23-01-2001",
		// })
		// if err != nil {
		// 	log.Println("error: ", err)
		// 	return
		// }

		// user, err := c.GetByPkey(&models.UserPrimaryKey{
		// 	Id: id,
		// })
		// if err != nil {
		// 	log.Println("error", err)
		// 	return
		// }
		// fmt.Println(*user)



// Delete :::::::>

			number, err := c.Delete(&models.UserPrimaryKey{
				Id: 6,
			})
			if err != nil {
				log.Println("error:", err)
				return
			}

			fmt.Println(number)

}
