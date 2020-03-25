package internal

import "net/http"

type StatusView struct{}

func (v *StatusView) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		v.status(w, r)
	}
}

func (v *StatusView) status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
