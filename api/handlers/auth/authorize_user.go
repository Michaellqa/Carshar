package auth

import (
	"Carshar/dal"
	"log"
	"net/http"
	"strconv"
)

type AuthorizeHandler struct {
	db *dal.UserDb
}

func NewAuthHandler(db *dal.UserDb) AuthorizeHandler {
	return AuthorizeHandler{db: db}
}

func (h AuthorizeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain charset=utf-8")

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

	id := strconv.FormatInt(int64(user.Id), 10)
	w.Write([]byte(id))
}
