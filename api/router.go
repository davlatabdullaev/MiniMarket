package api

import (
	"developer/api/handler"
	"net/http"
)

func New(h handler.Handler) {
	http.HandleFunc("/sale", h.Sale)
	http.HandleFunc("/branch", h.Branch)
	http.HandleFunc("/basket",h.Basket)
	http.HandleFunc("/product",h.Product)
	http.HandleFunc("/repository", h.Repository)
}