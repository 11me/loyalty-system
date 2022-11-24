package handler

import (
	"loyalty-system/db"
	"net/http"
)

func PostUser(db db.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
