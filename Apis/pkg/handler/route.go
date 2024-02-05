package handler

import "net/http"

func RouterUser(mux *http.ServeMux, storage Storage) {
	h := newUser(storage)
	mux.HandleFunc("/v1/users/create", h.created)
	mux.HandleFunc("/v1/users/all", h.getAll)
	mux.HandleFunc("/v1/users/delete", h.delete)
	mux.HandleFunc("/v1/users/id", h.getById)
	mux.HandleFunc("/v1/users/update", h.update)

}
