package handler

import (
	"Api-Go/authorization"
	"Api-Go/pkg/model"
	"encoding/json"
	"net/http"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type":"error", "message":"Metodo no permitido"}`))
		return
	}
	data := model.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		res := newResponse(Error, "Estructura no valida", nil)
		responseJSON(w, http.StatusBadRequest, res)
		return
	}
	if !isLoginValid(&data) {
		res := newResponse(Error, "Correo o Password no valido", nil)
		responseJSON(w, http.StatusBadRequest, res)
		return
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		res := newResponse(Error, "No de puedo generar token", nil)
		responseJSON(w, http.StatusInternalServerError, res)
		return
	}
	resp := newResponse(Message, "Ok", map[string]string{"token": token})
	responseJSON(w, http.StatusOK, resp)
}

func isLoginValid(data *model.Login) bool {
	return data.Email == "nico@es.es" && data.Password == "hola"
}
