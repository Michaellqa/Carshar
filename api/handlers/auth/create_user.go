package auth

import (
	"CarShar/dal"
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

	err := h.db.CreateUser(user)

	//TODO if user exist - conflict

	if err != nil {
		w.WriteHeader(502)
	}
}
