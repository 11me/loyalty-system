package handler

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"loyalty-system/db"
	"loyalty-system/model"
	"loyalty-system/pkg/logger"
	"net/http"
)

func PostUser(db db.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userReq model.PostUserRequest
		log := logger.GetLogger()

		if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
			log.Warnf("Unable to decode request body %s", err.Error())
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		if err := validate.Struct(userReq); err != nil {
			log.Warnf("Request body validation failed %s", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		passwordHash, err := bcrypt.GenerateFromPassword(userReq.PasswordPlain, bcrypt.DefaultCost)
		if err != nil {
			log.Errorf("Failed to hash a password %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user := &model.User{
			FirstName:    userReq.FirstName,
			LastName:     userReq.LastName,
			Email:        userReq.Email,
			PasswordHash: passwordHash,
		}
		err = db.CreateUser(r.Context(), user)
		if err != nil {
			log.Errorf("Failed to create a user %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	}
}
