package main

import (
	"github.com/iamrahultanwar/friskco/Config"
	"github.com/iamrahultanwar/friskco/Models"
	"github.com/iamrahultanwar/friskco/Routes"
)

func main() {
	Config.ConnectDB()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	r.Run()
}
