package handler

import (
	"developer/api/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (h Handler) Basket(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateBasket(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			h.GetBasketList(w, r)
		} else {
			h.GetBasketByID(w, r)
		}
	case http.MethodPut:
		h.UpdateBasket(w, r)
	case http.MethodDelete:
		h.DeleteBasket(w, r)
	}
}

func (h Handler) CreateBasket(w http.ResponseWriter, r *http.Request) {
	createBasket := models.CreateBasket{}
	
	if err := json.NewDecoder(r.Body).Decode(&createBasket); err != nil {
		hanldeResponse(w, http.StatusBadRequest, err)
		return
	}

	pKey, err := h.Store.Basket().Create(createBasket)
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	basket, err := h.Store.Basket().GetByID(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusCreated, basket)
}

func (h Handler) GetBasketByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		hanldeResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]
	var err error

	basket, err := h.Store.Basket().GetByID(models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusOK, basket)
}

func (h Handler) GetBasketList(w http.ResponseWriter, r *http.Request) {
	var (
		page, limit = 1, 10
		err         error
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
	resp, err := h.Store.Sale().GetList(models.GetListRequest{
		Page:   page,
		Limit:  limit,
	})

		 
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusOK, resp)
}

func (h Handler) UpdateBasket(w http.ResponseWriter, r *http.Request) {
	updateBasket := models.UpdateBasket{}

	if err := json.NewDecoder(r.Body).Decode(&updateBasket); err != nil {
		hanldeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.Store.Basket().Update(updateBasket)
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	basket, err := h.Store.Basket().GetByID(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusOK, basket)
}

func (h Handler) DeleteBasket(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		hanldeResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]

	err := h.Store.Basket().Delete(models.PrimaryKey{
		ID: id,
	})
	if err != nil{
		hanldeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	hanldeResponse(w, http.StatusOK, "data successfully deleted")
}
