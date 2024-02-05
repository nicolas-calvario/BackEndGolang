package handler

import (
	"Api-Go/pkg/model"
	"encoding/json"
	"net/http"
	"time"
)

type user struct {
	storage Storage
}

func newUser(storage Storage) user {
	return user{storage}
}

func (u *user) created(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type":"error", "message":"Metodo no permitido"}`))
		return
	}
	data := model.User{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type":"error", "message":"data no valida"}`))
		return
	}
	data.CreatedAt = time.Now()
	err = u.storage.Create(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type":"error", "message":"No se pudo guardar la informacion"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message_type":"message", "message":"Creado de manera correcta"}`))
	return
}
