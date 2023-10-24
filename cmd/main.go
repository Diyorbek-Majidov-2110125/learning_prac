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

	id, err := c.CreateUser(
		&models.CreateUser{
			Name: "Diyorbek",
			Surname: "Majidov",
			Birthday: "2001-10-30",
		},
	)
	if err != nil {
		log.Println("error while CreateUser:", err.Error())
		return
	}

	fmt.Println(id)

}
