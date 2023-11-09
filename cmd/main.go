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

//Addshopcart
	shopcart, err := c.AddShopcart(&models.AddShopcart{
		User_id: "19a477c4-6f1e-4444-97eb-8a48b292c30c",
		Product_id: "af122f7e-e33b-4ace-b68a-8b8c3cd16caf",
		Count: 4,
	})
	if err != nil {
		log.Println("error:", err)
		return
	}
	fmt.Println(shopcart)

	user, err := c.GetByPkey(&models.UserPrimaryKey{Id: "19a477c4-6f1e-4444-97eb-8a48b292c30c"})
	if err != nil {
		log.Println("error:", err)
		return
	}
	product, err := c.GetProductByPkey(&models.ProductPrimaryKey{Id: "af122f7e-e33b-4ace-b68a-8b8c3cd16caf"})
	if err != nil {
		log.Println("error:", err)
		return
	}
	fmt.Println("Customer:",user.Name)
	fmt.Println("Product Name:", product.Name)
	fmt.Println("Total Price:", shopcart.Count * int(product.Price))
//Create Product:::::>
	// id, err := c.CreateProduct(
	// 	&models.CreateProduct{
	// 		Name:  "banana",
	// 		Price: 19000,
	// 	},
	// )
	// if err != nil {
	// 	log.Println("error:", err)
	// 	return
	// }
	// log.Println(id)

	//Create User:::::>
	// id, err := c.CreateUser(
	// 	&models.CreateUser{
	// 		Name: "Sasha",
	// 		Surname: "G'aniyev",
	// 		Birthday: "2006-10-30",
	// 	},
	// )
	// if err != nil {
	// 	log.Println("error while CreateUser:", err.Error())
	// 	return
	// }

	// fmt.Println(id)

	//GetListProducts::::>
	// products, err := c.GetListProducts(&models.GetListProductRequest{
	// 	Offset: 1,
	// 	Limit:  10,
	// 	Search: "",
	// })
	// if err != nil {
	// 	log.Println("error is coming from GetlistProducts...")
	// 	return
	// }

	// for _, product := range products.Products {
	// 	fmt.Println(product.Id, product.Name, product.Price)
	// }

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
	// user, err := c.GetByPkey(&models.UserPrimaryKey{
	// 	Id: "668f9621-62bb-45ea-af3f-82a40713b4d7",
	// })
	// if err != nil {
	// 	log.Println("error: ", err)
	// 	return
	// }
	// fmt.Println(*user)

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
