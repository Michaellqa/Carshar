package auth

import (
	"Carshar/dal"
	"encoding/json"
	"log"
	"net/http"
)

type AuthorizeHandler struct {
	db dal.AuthRepository
}

func NewAuthHandler(db dal.AuthRepository) AuthorizeHandler {
	return AuthorizeHandler{db: db}
}

func (h AuthorizeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json charset=utf-8")

	phone := r.FormValue("Phone")
	pass := r.FormValue("Password")

	if phone == "" || pass == "" {
		w.WriteHeader(400)
		return
	}

	user, found, err := h.db.FindUser(phone)
	if err != nil {
		log.Println(err)
		w.WriteHeader(502)
		return
	}

	if !found {
		w.WriteHeader(404)
		return
	}

	if pass != user.Password {
		w.WriteHeader(403)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Println(err)
		w.WriteHeader(502)
	}
}
