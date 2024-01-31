package handler

import (
	"developer/api/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (h Handler) Product(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateProduct(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			h.GetProductList(w, r)
		} else {
			h.GetProductByID(w, r)
		}
	case http.MethodPut:
		h.UpdateProduct(w, r)
	case http.MethodDelete:
		h.DeleteProduct(w, r)
	}
}

func (h Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	createProd := models.CreateProduct{}
	
	if err := json.NewDecoder(r.Body).Decode(&createProd); err != nil {
		hanldeResponse(w, http.StatusBadRequest, err)
		return
	}

	pKey, err := h.Store.Product().Create(createProd)
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	product, err := h.Store.Product().GetByID(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusCreated, product)
}

func (h Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		hanldeResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]
	var err error

	product, err := h.Store.Product().GetByID(models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusOK, product)
}

func (h Handler) GetProductList(w http.ResponseWriter, r *http.Request) {
	var (
		page, limit = 1, 10
		err         error
		search 		string
	)
	values := r.URL.Query()

	if len(values["page"]) > 0 {
		page, err = strconv.Atoi(values["page"][0])
		if err != nil {
			page = 1
		}
	}

	if len(values["limit"]) > 0 {
		limit, err = strconv.Atoi(values["limit"][0])
		if err != nil {
			fmt.Println("limit", values["limit"])
			limit = 10
		}
	}

	if values["search"][0] != "" {
		search = values["search"][0]
	}else{
		search = ""
	}
	resp, err := h.Store.Product().GetList(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

		 
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusOK, resp)
}

func (h Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	updateProd := models.UpdateProduct{}

	if err := json.NewDecoder(r.Body).Decode(&updateProd); err != nil {
		hanldeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.Store.Product().Update(updateProd)
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	product, err := h.Store.Product().GetByID(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusOK, product)
}

func (h Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		hanldeResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]

	err := h.Store.Product().Delete(models.PrimaryKey{
		ID: id,
	})
	if err != nil{
		hanldeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	hanldeResponse(w, http.StatusOK, "data successfully deleted")
}
