package main

import (
	"trackly/controller"
	"trackly/server"
)

func main(){
	config := server.Config{
		Port: 80,
		Host: "",
	}

	server := server.New(config)

	server.AddController(controller.UserController)

	server.Start()
}