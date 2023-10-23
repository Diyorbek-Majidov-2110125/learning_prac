package main

import (
	"app/controller"
	"app/models"
	"fmt"
)

func main() {

	users := controller.GenerateUser(76)
	for _, user := range users {
		fmt.Println(user)
	}
	// users, err := controller.GetListUsers(models.GetListRequest{
	// 	Offset: 20,
	// 	Limit:  15,
	// })
	// if err {
	// 	panic(users)
	// }

	//page
	// fmt.Println("The number of pages: ", controller.AllPages())
	// var number int
	// fmt.Println("Enter the number of page:")
	// fmt.Scanln(&number)
	// users = controller.GetPageByNumber(number)

	// users =  controller.GetByName(users,models.GetListRequest{})
	// for _, user := range users {
	// 	fmt.Println(user)
	// }

	fmt.Println("\nEnter the duration(from and to..e.x \"2023-10-17\"):")
	var from,to string
	fmt.Print("\nfromDate: ")
	fmt.Scanln(&from)
	fmt.Print("toDate:")
	fmt.Scanln(&to)
	fmt.Println()
	users, err := controller.SortByDate(models.GetListDate{
		ToDate: to,
		FromDate: from,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(users) == 0 {
		fmt.Println("result: No match")
	}
	for _, user := range users {
		fmt.Println(user)
	}
}
