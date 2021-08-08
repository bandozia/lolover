package handler

import (
	"encoding/json"
	"net/http"
)

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

func WriteErr(w http.ResponseWriter, c int, m string) {
	w.WriteHeader(c)
	w.Write([]byte(m))
}

func WriteJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	if e := json.NewEncoder(w).Encode(data); e != nil {
		WriteErr(w, http.StatusInternalServerError, "internal error")
	}
}
