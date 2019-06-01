package csurl

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func IntIdParam(r *http.Request) (int, bool) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		return -1, false
	}
	carId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return -1, false
	}
	return int(carId), true
}
