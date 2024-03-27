package main

import (
	"fmt"
	"github.com/run100/go_module_test/models"
	"gorm.io/gorm"
)
import "gorm.io/driver/mysql"

func main() {
	//server := NewServer("0.0.0.0", "8888")
	//server.Start()

	dsn := "lolga:2zb0U1&oOMmlUYpo@tcp(223.247.159.5:53306)/go_chat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Printf("%v, %v\n", db, err)

	db.AutoMigrate(&models.GroupBasic{})

	//user := &models.UserBasic{}
	//user.Name = "逍遥"
	//db.Create(user)

	//fmt.Println(db.First(user, 1))
	//
	//db.Model(user).Update("name", "333")

}
