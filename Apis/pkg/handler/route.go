package handler

import "net/http"

func RouterUser(mux *http.ServeMux, storage Storage) {
	h := newUser(storage)
	mux.HandleFunc("/v1/users/create", h.created)

}
