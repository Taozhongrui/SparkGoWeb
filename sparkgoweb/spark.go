package sparkgoweb

import "net/http"

type Engine struct {
	routers map[string]http.HandlerFunc
}

func Default() *Engine {
	e := Engine{
		routers: make(map[string]http.HandlerFunc, 0),
	}
	return &e
}
