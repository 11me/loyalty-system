package handler

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"loyalty-system/config"
	"loyalty-system/db"
	"loyalty-system/model"
	"loyalty-system/pkg/logger"
	"net/http"
	"strconv"
	"time"
)

func PostUser(db db.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userReq model.PostUserRequest
		log := logger.GetLogger()

		if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
			log.Warnf("unable to decode request body %s", err.Error())
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		if err := validate.Struct(userReq); err != nil {
			log.Warnf("request body validation failed %s", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		checkUser, err := db.GetUserByEmail(r.Context(), userReq.Email)
		if err != nil {
			log.Errorf("failed to get user %s with following error %s", userReq.Email, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if checkUser != nil {
			log.Warnf("user with email %s already exists", userReq.Email)
			w.WriteHeader(http.StatusConflict)
			return
		}

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(userReq.PasswordPlain), bcrypt.DefaultCost)
		if err != nil {
			log.Errorf("failed to hash a password %s", err.Error())
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
			log.Errorf("failed to create a user %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func AuthUser(db db.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var authUserReq model.AuthUserRequest
		log := logger.GetLogger()

		if err := json.NewDecoder(r.Body).Decode(&authUserReq); err != nil {
			log.Warningf("unable to decode token request body %s", err.Error())
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		if err := validate.Struct(authUserReq); err != nil {
			log.Warningf("invalid auth payload %s", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := db.GetUserByEmail(r.Context(), authUserReq.Email)
		if err != nil {
			log.Errorf("failed to retrieve a user %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if user == nil {
			log.Warnf("user with email %s not found", authUserReq.Email)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(authUserReq.PasswordPlain))
		if err != nil {
			log.Warnf("invalid user credentials for user with email %s", authUserReq.Email)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var token string
		token, err = generateJWT(user)
		if err != nil {
			log.Errorf("failed to generate jwt %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		writeJSON(w, model.TokenResponse{Token: token})
	}
}

func generateJWT(user *model.User) (string, error) {
	cfg := config.GetConfig()
	signingKey := []byte(cfg.Secret)

	tokenClaim := model.TokenClaim{
		strconv.Itoa(user.ID),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Issuer:    "LoyaltySystem",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaim)
	return token.SignedString(signingKey)
}
