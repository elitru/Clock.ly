package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	router      *mux.Router
	controllers []Controller
	isRunning   bool
	Config      Config
}

func (server *Server) IsRunning() bool {
	return server.isRunning
}

type Config struct {
	Port int
	Host string
}

type Controller struct {
	BaseRoute string
	Routes    []Route
}

type Route struct {
	Route       string
	Interceptor func(w http.ResponseWriter, r *http.Request, next func(w http.ResponseWriter, r *http.Request, additional map[string]interface{})) func(w http.ResponseWriter, r *http.Request)
	Handler func(w http.ResponseWriter, r *http.Request, additional map[string]interface{})
	Method  string
}

func (route *Route) HasInterceptor() bool {
	return route.Interceptor != nil
}
