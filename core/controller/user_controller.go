package controller

import (
	"fmt"
	"net/http"
	"trackly/server"
)

var UserController = server.Controller{
	BaseRoute: "/user",
	Routes: []server.Route{
		{
			Route:       "",
			Interceptor: nil,
			Handler:     login,
			Method:      "GET",
		},
	},
}

func login(w http.ResponseWriter, r *http.Request, additional map[string]interface{}) {
	fmt.Fprint(w, "Hello World")
}