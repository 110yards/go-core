package api

import (
	"net/http"

	"110yards.ca/libs/go/core/logger"
	"github.com/go-chi/chi/v5"
)

type Method int

const (
	Method_Get Method = iota
	Method_Put
	Method_Post
	Method_Delete
)

type Route struct {
	Method         Method
	Path           string
	Handler        http.HandlerFunc
	AllowAnonymous bool
}

func AddRoutes(router *chi.Mux, routes []Route) {
	for _, route := range routes {
		AddRoute(router, route)
	}
}

func AddRoute(router *chi.Mux, route Route) {
	switch route.Method {
	case Method_Get:
		router.Get(route.Path, route.Handler)
		logger.Infof("GET %s", route.Path)
	case Method_Post:
		router.Post(route.Path, route.Handler)
		logger.Infof("POST %s", route.Path)
	case Method_Put:
		router.Put(route.Path, route.Handler)
		logger.Infof("PUT %s", route.Path)
	case Method_Delete:
		router.Delete(route.Path, route.Handler)
		logger.Infof("DELETE %s", route.Path)
	}

}

func Get(path string, handler http.HandlerFunc) Route {
	return Route{
		Method:  Method_Get,
		Path:    path,
		Handler: handler,
	}
}

func Put(path string, handler http.HandlerFunc) Route {
	return Route{
		Method:  Method_Put,
		Path:    path,
		Handler: handler,
	}
}

func Post(path string, handler http.HandlerFunc) Route {
	return Route{
		Method:  Method_Post,
		Path:    path,
		Handler: handler,
	}
}

func Delete(path string, handler http.HandlerFunc) Route {
	return Route{
		Method:  Method_Delete,
		Path:    path,
		Handler: handler,
	}
}
