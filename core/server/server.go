package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func New(config Config) *Server {
	router := mux.NewRouter()

	server := Server{
		router: router,
		isRunning: false,
		Config: config,
	}

	return &server
}

func (server *Server) AddController(controller Controller) {
	server.controllers = append(server.controllers, controller)

	for _, route := range controller.Routes {
		if !route.HasInterceptor() {
			server.router.HandleFunc(controller.BaseRoute + route.Route, func(w http.ResponseWriter, r *http.Request) {
				route.Handler(w, r, nil)
			}).Methods(route.Method)
			continue
		}

		server.router.HandleFunc(controller.BaseRoute + route.Route, func(w http.ResponseWriter, r *http.Request) {
			route.Interceptor(w, r, route.Handler)
		}).Methods(route.Method)
	}
}

func (server *Server) Start() error {
	server.isRunning = true
	fmt.Println("Server running on -> " + server.Config.Host + ":" + strconv.Itoa(server.Config.Port))
	return http.ListenAndServe(server.Config.Host + ":" + strconv.Itoa(server.Config.Port), server.router)
}