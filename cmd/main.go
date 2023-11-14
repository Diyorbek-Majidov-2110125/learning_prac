package main

func main() {

}

func User() {

	// 	cfg := config.Load()

	// 	jsondb, err := jsondb.NewFileJson(&cfg)
	// 	if err != nil {
	// 		panic("error while connect to json file: " + err.Error())
	// 	}
	// 	defer jsondb.CloseDB()

	// 	c := controller.NewController(&cfg, jsondb)

	// //Create User:::::>
	// 	id, err := c.CreateUser(
	// 		&models.CreateUser{
	// 			Name: "Sardor",
	// 			Surname: "Anvarov",
	// 			Birthday: "2005-09-10",
	// 			Balance: 1000000,
	// 		},
	// 	)
	// 	if err != nil {
	// 		fmt.Println("error while CreateUser:", err.Error())
	// 		return
	// 	}

	// 	fmt.Println(id)

	// 	//GetlistUsers::::>

	// 	users, err := c.GetListUsers(&models.GetListRequest{
	// 		Offset: 1,
	// 		Limit: 8,
	// 		Search: "",
	// 	})
	// 	if err != nil {
	// 		log.Println("error is coming from GetlistUsers...")
	// 		return
	// 	}

	// 	for _, user := range users.Users {
	// 		fmt.Println(user.Id,user.Name,user.Surname, user.Birthday)
	// 	}

	// 	//GetUserById::::>
	// 	user, err := c.GetByPkey(&models.UserPrimaryKey{
	// 		Id: "668f9621-62bb-45ea-af3f-82a40713b4d7",
	// 	})
	// 	if err != nil {
	// 		log.Println("error: ", err)
	// 		return
	// 	}
	// 	fmt.Println(*user)

	// 	//Update::::>

	// 	id , err := c.Update(&models.UpdateUser{
	// 		Id: "",
	// 		Name: "Jasurbek",
	// 		Surname: "Abdullaev",
	// 		Birthday: "23-01-2001",
	// 	})
	// 	if err != nil {
	// 		log.Println("error: ", err)
	// 		return
	// 	}

	// 	user, err := c.GetByPkey(&models.UserPrimaryKey{
	// 		Id: id,
	// 	})
	// 	if err != nil {
	// 		log.Println("error", err)
	// 		return
	// 	}
	// 	fmt.Println(*user)

	// 	// Delete :::::::>

	// 	number, err := c.Delete(&models.UserPrimaryKey{
	// 		Id: 6,
	// 	})
	// 	if err != nil {
	// 		log.Println("error:", err)
	// 		return
	// 	}

	// 	fmt.Println(number)

	// 	//GetListUserByName
	// 	users, err := c.GetByName(&models.GetListRequest{
	// 		Offset: 1,
	// 		Limit:  0,
	// 		Search: "",
	// 	})
	// 	if err != nil {
	// 		fmt.Println("error", err)
	// 	}
	// 	for _, user := range *users {
	// 		fmt.Println(user)
	// 	}

	// 	//Choose User by BirthDate ::::::>

	// 	users, err := c.ChooseByBirthDate(&models.GetListDate{
	// 		FromDate: "1998-02-02",
	// 		ToDate: "2000-01-01",
	// 	})
	// 	if err != nil {
	// 		fmt.Println("error:", err)
	// 	}

	// 	for _, user :=range users {
	// 		fmt.Println(user)
	// 	// }

	// 	//Users Transfer Money:

	// 	status, err := c.TransferBalance(&models.TransferBalance{
	// 		SenderId:               "a0c86706-efda-41cb-9e25-d28bba85483c",
	// 		ReceiverId:             "791eb41b-4b42-4eaa-ac95-06df4ace43bf",
	// 		Money:                  1000,
	// 		Service_fee_percentage: 1,
	// 	})
	// 	if err != nil {
	// 		fmt.Println("error:", err)
	// 		return
	// 	}

	// 	fmt.Println(status)

	// 	// calculate total and withdraw money from user:

	// 	userId := "791eb41b-4b42-4eaa-ac95-06df4ace43bf"
	// 	totalPayment, err := c.CalcTotal(&models.CalculateShop{
	// 		UserId:         userId,
	// 		Discount:       10,
	// 		DiscountStatus: "percentage",
	// 	})
	// 	if err != nil {
	// 		fmt.Println("ERROR:", err)
	// 	}
	// 	fmt.Println(totalPayment)

	// 	err = c.WithdrawUserBalance(&models.UserPrimaryKey{Id: userId}, totalPayment)
	// 	if err != nil {
	// 		return
	// 	}

}

func Product() {

	// cfg := config.Load()

	// jsondb, err := jsondb.NewFileJson(&cfg)
	// if err != nil {
	// 	panic("error while connect to json file: " + err.Error())
	// }
	// defer jsondb.CloseDB()

	// c := controller.NewController(&cfg, jsondb)

	// //Create Product:::::>

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

	// //GetListProducts::::>

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

}

func Shopcart() {
	// cfg := config.Load()

	// jsondb, err := jsondb.NewFileJson(&cfg)
	// if err != nil {
	// 	panic("error while connect to json file: " + err.Error())
	// }
	// defer jsondb.CloseDB()

	// c := controller.NewController(&cfg, jsondb)

	// //Addshopcart

	// shopcart, err := c.AddShopcart(&models.AddShopcart{
	// 	User_id:    "791eb41b-4b42-4eaa-ac95-06df4ace43bf",
	// 	Product_id: "af122f7e-e33b-4ace-b68a-8b8c3cd16caf",
	// 	Count:      1,
	// })
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }
	// fmt.Println(shopcart)

	// //Total Price:

	// user, err := c.GetByPkey(&models.UserPrimaryKey{Id: "19a477c4-6f1e-4444-97eb-8a48b292c30c"})
	// if err != nil {
	// 	log.Println("error:", err)
	// 	return
	// }
	// product, err := c.GetProductByPkey(&models.ProductPrimaryKey{Id: "af122f7e-e33b-4ace-b68a-8b8c3cd16caf"})
	// if err != nil {
	// 	log.Println("error:", err)
	// 	return
	// }
	// fmt.Println("Customer:", user.Name)
	// fmt.Println("Product Name:", product.Name)
	// fmt.Println("Total Price:", shopcart.Count*int(product.Price))

}
