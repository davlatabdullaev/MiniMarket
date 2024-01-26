package api

import (
	"developer/api/handler"
	"net/http"
)

func New(h handler.Handler) {
	http.HandleFunc("/sale", h.Sale)
	http.HandleFunc("/branch", h.Branch)
}