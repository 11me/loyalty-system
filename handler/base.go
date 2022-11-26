package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var (
	jsonContentType = []string{"application/json; charset=utf=8"}
	validate        = validator.New()
)

func JSONResponse(w http.ResponseWriter, obj any) error {
	writeContentType(w, jsonContentType)
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
