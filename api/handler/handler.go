package handler

import (
	"developer/api/models"
	"developer/storage"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Store storage.IStorage
}

func New(store storage.IStorage) Handler {
	return Handler{
		Store: store,
	}
}

func hanldeResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	resp := models.Response{}

	switch code := statusCode; {
	case code < 400:
		resp.Description = "success"
	case code < 500:
		resp.Description = "bad request"
	default:
		resp.Description = "internal server error"
	}

	resp.StatusCode = statusCode
	resp.Data = data

	js, _ := json.Marshal(resp)

	w.WriteHeader(statusCode)
	w.Write(js)
}