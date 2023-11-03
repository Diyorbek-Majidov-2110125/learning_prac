package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsondb"
	"fmt"
	"log"
)

func main() {

	cfg := config.Load()

	jsondb, err := jsondb.NewFileJson(&cfg)
	if err != nil {
		panic("error while connect to json file: " + err.Error())
	}
	defer jsondb.CloseDB()

	c := controller.NewController(&cfg, jsondb)

//Create User:::::>
	// id, err := c.CreateUser(
	// 	&models.CreateUser{
	// 		Name: "Mustafo",
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
	// 	Offset: 1,
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
	user, err := c.GetByPkey(&models.UserPrimaryKey{
		Id: "668f9621-62bb-45ea-af3f-82a40713b4d7",
	})
	if err != nil {
		log.Println("error: ", err)
		return
	}
	fmt.Println(*user)

//Update::::>

		// id , err := c.Update(&models.UpdateUser{
		// 	Id: "",
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

			// number, err := c.Delete(&models.UserPrimaryKey{
			// 	Id: 6,
			// })
			// if err != nil {
			// 	log.Println("error:", err)
			// 	return
			// }

			// fmt.Println(number)

}
