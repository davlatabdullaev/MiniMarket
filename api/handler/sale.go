package handler

import (
	"developer/api/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (h Handler) Sale(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateSale(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			h.GetSaleList(w, r)
		} else {
			h.GetSaleByID(w, r)
		}
	case http.MethodPut:
		h.UpdateSale(w, r)
	case http.MethodDelete:
		h.DeleteSale(w, r)
	}
}

func (h Handler) CreateSale(w http.ResponseWriter, r *http.Request) {
	createSale := models.CreateSale{}
	
	if err := json.NewDecoder(r.Body).Decode(&createSale); err != nil {
		hanldeResponse(w, http.StatusBadRequest, err)
		return
	}

	pKey, err := h.Store.Sale().Create(createSale)
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	sale, err := h.Store.Sale().GetByID(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusCreated, sale)
}

func (h Handler) GetSaleByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		hanldeResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]
	var err error

	sale, err := h.Store.Sale().GetByID(models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusOK, sale)
}

func (h Handler) GetSaleList(w http.ResponseWriter, r *http.Request) {
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
	resp, err := h.Store.Sale().GetList(models.GetListRequest{
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

func (h Handler) UpdateSale(w http.ResponseWriter, r *http.Request) {
	updateSale := models.UpdateSale{}

	if err := json.NewDecoder(r.Body).Decode(&updateSale); err != nil {
		hanldeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.Store.Sale().Update(updateSale)
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	sale, err := h.Store.Sale().GetByID(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusOK, sale)
}

func (h Handler) DeleteSale(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		hanldeResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]

	err := h.Store.Sale().Delete(models.PrimaryKey{
		ID: id,
	})
	if err != nil{
		hanldeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	hanldeResponse(w, http.StatusOK, "data successfully deleted")
}
