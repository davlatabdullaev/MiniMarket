package handler

import (
	"context"
	"developer/api/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateStorage godoc
// @Router       /storage [POST]
// @Summary      Create a new storage
// @Description  create a new storage
// @Tags         storage
// @Accept       json
// @Produce      json
// @Param 		 storage body models.CreateStorage false "storage"
// @Success      200  {object}  models.Storage
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateStorage(c *gin.Context) {
	createStorage := models.CreateStorage{}

	err := c.ShouldBindJSON(&createStorage)
	if err != nil{
		handleResponse(c, "Error in handlers, whiel reading storage json!",http.StatusBadRequest,err.Error())
	}
	 

	pKey, err := h.Store.Storage().Create(context.Background(),createStorage)
	if err != nil {
		handleResponse(c, "Error in handlers, while creating storage!",http.StatusInternalServerError, err.Error())
		return
	}

	store, err := h.Store.Storage().GetByID(context.Background(),models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting storage by id!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusCreated, store)
}

// GetStorage godoc
// @Router       /storage/{id} [GET]
// @Summary      Get storage by id
// @Description  get storage by id
// @Tags         storage
// @Accept       json
// @Produce      json
// @Param 		 id path string true "storage_id"
// @Success      200  {object}  models.Storage
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetStorageByID(c *gin.Context) {
	 uid := c.Param("id")

	store, err := h.Store.Storage().GetByID(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting storage by id!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, store)
}

// GetStorageList godoc
// @Router       /storages [GET]
// @Summary      Get storage list
// @Description  get storage list
// @Tags         storage
// @Accept       json
// @Produce      json
// @Param 		 page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.StoragesResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetStorageList(c *gin.Context) {
	var (
		page, limit = 1, 10
		err         error
		search 		string
	)
	
	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	search = c.Query("search")

	resp, err := h.Store.Storage().GetList(context.Background(),models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

		 
	if err != nil {
		handleResponse(c, "Error in handlers,while getting storages!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, resp)
}

// UpdateStorage godoc
// @Router       /storage/{id} [PUT]
// @Summary      Update storage
// @Description  update storage
// @Tags         storage
// @Accept       json
// @Produce      json
// @Param 		 id path string true "storage_id"
// @Param 		 storage body models.UpdateStorage false "storage"
// @Success      200  {object}  models.Storage
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateStorage(c *gin.Context) {
	updateStorage := models.UpdateStorage{}
	
	uid := c.Param("id")

	if uid == ""{
		handleResponse(c, "invalid uuid!", http.StatusBadRequest,errors.New("uuid is not valid"))
		return
	 }
 
	 updateStorage.ID = uid

	 err := c.ShouldBindJSON(&updateStorage)
	 if err != nil{
		handleResponse(c,"Error in handlers, while reading storage json from client!",http.StatusBadRequest,err)
		return
	 }

	pKey, err := h.Store.Storage().Update(context.Background(),updateStorage)
	if err != nil {
		handleResponse(c, "Error in handlers, while updating storage!",http.StatusInternalServerError, err.Error())
		return
	}

	store, err := h.Store.Storage().GetByID(context.Background(),models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting storage by id!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, store)
}


// DeleteStorage godoc
// @Router       /storage/{id} [DELETE]
// @Summary      Delete storage
// @Description  delete storage
// @Tags         storage
// @Accept       json
// @Produce      json
// @Param 		 id path string true "storage_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteStorage(c *gin.Context) {

	uid := c.Param("id")

	err := h.Store.Storage().Delete(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil{
		handleResponse(c, "Error in handlers, while deleting storage!",http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "",http.StatusOK, "data successfully deleted")
}
