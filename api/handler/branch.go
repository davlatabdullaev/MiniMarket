package handler

import (
	"developer/api/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (h Handler) Branch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateBranch(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			h.GetBranchList(w, r)
		} else {
			h.GetBranchByID(w, r)
		}
	case http.MethodPut:
		h.UpdateBranch(w, r)
	case http.MethodDelete:
		h.DeleteBranch(w, r)
	}
}

func (h Handler) CreateBranch(w http.ResponseWriter, r *http.Request) {
	createBranch := models.CreateBranch{}
	
	if err := json.NewDecoder(r.Body).Decode(&createBranch); err != nil {
		hanldeResponse(w, http.StatusBadRequest, err)
		return
	}

	pKey, err := h.Store.Branch().Create(createBranch)
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	sale, err := h.Store.Branch().GetByID(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusCreated, sale)
}

func (h Handler) GetBranchByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		hanldeResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}
	
	id := values["id"][0]
	var err error

	sale, err := h.Store.Branch().GetByID(models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusOK, sale)
}

func (h Handler) GetBranchList(w http.ResponseWriter, r *http.Request) {
	var (
		search string
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

	if values["search"][0] != ""{
		search = values["search"][0]
	}else{
		search = ""
	}

	

	resp, err := h.Store.Branch().GetList(models.GetListRequest{
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

func (h Handler) UpdateBranch(w http.ResponseWriter, r *http.Request) {
	updateBranch := models.UpdateBranch{}

	if err := json.NewDecoder(r.Body).Decode(&updateBranch); err != nil {
		hanldeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.Store.Branch().Update(updateBranch)
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	branch, err := h.Store.Branch().GetByID(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusOK, branch)
}

func (h Handler) DeleteBranch(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		hanldeResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]

	err := h.Store.Branch().Delete(models.PrimaryKey{
		ID: id,
	})
	if err != nil{
		hanldeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	hanldeResponse(w, http.StatusOK, "data successfully deleted")
}
