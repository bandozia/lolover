package handler

import "net/http"

var midlewares []func(w http.ResponseWriter)

type Handler struct {
	HandleFunc func(w http.ResponseWriter, r *http.Request)
}

func (h Handler) Handle(w http.ResponseWriter, r *http.Request) {
	for _, f := range midlewares {
		f(w)
	}
	h.HandleFunc(w, r)
}

func AddMiddleware(f func(w http.ResponseWriter)) {
	midlewares = append(midlewares, f)
}
