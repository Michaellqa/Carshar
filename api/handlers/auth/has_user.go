package auth

import (
	"net/http"
	"strconv"
)

func UserToken(r *http.Request) (int, error) {
	idStr := r.Header.Get("Authorization")
	uid, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		return 0, err
	}
	return int(uid), nil
}
