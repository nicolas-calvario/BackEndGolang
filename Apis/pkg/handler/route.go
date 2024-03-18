package handler

import (
	"Api-Go/pkg/middleware"
	"net/http"
)

func RouterUser(mux *http.ServeMux, storage Storage) {
	h := newUser(storage)
	mux.HandleFunc("/v1/users/create", middleware.Log(middleware.Authentication(h.created)))
	mux.HandleFunc("/v1/users/all", middleware.Log(h.getAll))
	mux.HandleFunc("/v1/users/delete", middleware.Log(middleware.Authentication(h.delete)))
	mux.HandleFunc("/v1/users/byId", middleware.Log(h.getById))
	mux.HandleFunc("/v1/users/update", middleware.Log(h.update))

}

func RouterLogin(mux *http.ServeMux, storage Storage) {
	h := newLogin(storage)
	mux.HandleFunc("/v1/login", h.login)
}
