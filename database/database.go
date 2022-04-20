package database

import (
	"fmt"
	"task_management/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

var DSN = "host=localhost user=postgres password=password dbname=task_management port=5432 sslmode=disable"

func Migration() {
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to the databse")
	DB.AutoMigrate(&model.User{}, &model.Preference{}, &model.Team{}, &model.Workspace{},&model.Task{})
}
