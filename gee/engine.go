package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func NewEngine() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

func (t *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	t.router[key] = handler
}
func (t *Engine) Get(pattern string, handler HandlerFunc) {
	t.addRoute("GET", pattern, handler)
}
func (t *Engine) Post(pattern string, handler HandlerFunc) {
	t.addRoute("POST", pattern, handler)
}
func (t *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, t)
}
func (t *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := t.router[key]; ok {
		handler(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}
