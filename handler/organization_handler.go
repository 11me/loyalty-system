package handler

import (
	"encoding/json"
	"loyalty-system/db"
	"loyalty-system/model"
	"loyalty-system/pkg/logger"
	"net/http"
)

func PostOrganization(db db.Organization) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var postOrgReq model.PostOrganizationRequest
		log := logger.GetLogger()

		if err := json.NewDecoder(r.Body).Decode(&postOrgReq); err != nil {
			log.Warnf("Unprocesseable organization payload %s.", err.Error())
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		if err := validate.Struct(postOrgReq); err != nil {
			log.Warnf("Invalid organization payload %s.", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		organization := &model.Organization{
			Name: postOrgReq.Name,
		}

		if err := db.CreateOrganization(r.Context(), organization); err != nil {
			log.Errorf("Failed to create organization %s.", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}
