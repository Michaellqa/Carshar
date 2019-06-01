package auth

import (
	"Carshar/api/handlers/csurl"
	"Carshar/dal"
	"encoding/json"
	"log"
	"net/http"
)

type UserInfoHandler struct {
	db *dal.UserDb
}

func NewUserInfoHandler(db *dal.UserDb) UserInfoHandler {
	return UserInfoHandler{db: db}
}

func (h UserInfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI)

	userId, ok := csurl.IntIdParam(r)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.db.GetUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
