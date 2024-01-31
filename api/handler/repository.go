package handler

import (
	"developer/api/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (h Handler) Repository(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateRepository(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			h.GetRepositoryList(w, r)
		} else {
			h.GetRepositoryByID(w, r)
		}
	case http.MethodPut:
		h.UpdateRepository(w, r)
	case http.MethodDelete:
		h.DeleteRepository(w, r)
	}
}

func (h Handler) CreateRepository(w http.ResponseWriter, r *http.Request) {
	createRepo := models.CreateRepository{}
	
	if err := json.NewDecoder(r.Body).Decode(&createRepo); err != nil {
		hanldeResponse(w, http.StatusBadRequest, err)
		return
	}

	pKey, err := h.Store.Repository().Create(createRepo)
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	repository, err := h.Store.Repository().GetByID(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusCreated, repository)
}

func (h Handler) GetRepositoryByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		hanldeResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]
	var err error

	repository, err := h.Store.Repository().GetByID(models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusOK, repository)
}

func (h Handler) GetRepositoryList(w http.ResponseWriter, r *http.Request) {
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
	resp, err := h.Store.Repository().GetList(models.GetListRequest{
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

func (h Handler) UpdateRepository(w http.ResponseWriter, r *http.Request) {
	updateRepo := models.UpdateRepository{}

	if err := json.NewDecoder(r.Body).Decode(&updateRepo); err != nil {
		hanldeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.Store.Repository().Update(updateRepo)
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	repository, err := h.Store.Repository().GetByID(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusOK, repository)
}

func (h Handler) DeleteRepository(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		hanldeResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]

	err := h.Store.Repository().Delete(models.PrimaryKey{
		ID: id,
	})
	if err != nil{
		hanldeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	hanldeResponse(w, http.StatusOK, "data successfully deleted")
}
