package main

import (
	"fmt"
	"golang-mini-rest/Config"
	"golang-mini-rest/Models"
	"golang-mini-rest/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	//running
	r.Run(":8080")
}

// Testing Conflict Git
// Editing with branch feature 1
// Editing with branch feature 2
// Conflict Handle Git Branch
