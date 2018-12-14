package auth

import (
	"CarShar/dal"
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

	phone := r.FormValue("phone")
	pass := r.FormValue("password")

	if phone == "" || pass == "" {
		w.WriteHeader(400)
		return
	}

	user, ok, err := h.db.FindUser(phone)
	if err != nil {
		log.Println(err)
		w.WriteHeader(502)
		return
	}

	if !ok {
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

type UserPass struct {
	Phone string
	Pass  string
}
