package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"loyalty-system/pkg/logger"
	"net/http"
)

var (
	jsonContentType = []string{"application/json; charset=utf=8"}
	validate        = validator.New()
)

func writeJSON(w http.ResponseWriter, obj any) {
	log := logger.GetLogger()
	writeContentType(w, jsonContentType)
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		log.Errorf("Object marshaling failed %s", err.Error())
	}
	_, err = w.Write(jsonBytes)
	if err != nil {
		log.Errorf("Failed to write json response %s", err.Error())
	}
}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
