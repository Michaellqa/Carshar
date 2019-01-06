package auth

import (
	"Carshar/dal"
	"encoding/json"
	"log"
	"net/http"
)

type CreateUserHandler struct {
	db dal.AuthRepository
}

func NewCreateUserHandler(db dal.AuthRepository) CreateUserHandler {
	return CreateUserHandler{db: db}
}

func (h CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json charset=utf-8")

	var user dal.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	if user.Name == "" || user.Password == "" || user.Phone == "" {
		w.WriteHeader(400)
		w.Write([]byte("All required fields should be filled"))
		return
	}

	ok, err := h.db.CreateUser(user)
	if err != nil {
		w.WriteHeader(502)
		return
	}

	if !ok {
		w.WriteHeader(409)
		return
	}
}
