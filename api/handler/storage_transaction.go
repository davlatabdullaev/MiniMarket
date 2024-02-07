package handler

import (
	"context"
	"developer/api/models"
	"errors"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// CreateStorageTransaction godoc
// @Router       /storage_transaction [POST]
// @Summary      Create a new storage_transaction
// @Description  create a new  storage_transaction
// @Tags         storage_transaction
// @Accept       json
// @Produce      json
// @Param 		 storage_transaction body models.CreateStorageTransaction false "storage_transaction"
// @Success      200  {object}  models.StorageTransaction
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateStorageTransaction(c *gin.Context) {
	createStTrans := models.CreateStorageTransaction{}

	err := c.ShouldBindJSON(&createStTrans)
	if err != nil{
		handleResponse(c,"Error in handlers, while reading storage transaction json!",http.StatusBadRequest,err.Error())
	}

	pKey, err := h.Store.StorageTransaction().Create(context.Background(),createStTrans)
	if err != nil {
		handleResponse(c, "Error in handlers, while creating storage transaction!",http.StatusInternalServerError, err.Error())
		return
	}

	stTransaction, err := h.Store.StorageTransaction().GetByID(context.Background(),models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error while getting storage transaction by id!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusCreated, stTransaction)
}

// GetStorageTransaction godoc
// @Router       /storage_transaction/{id} [GET]
// @Summary      Get storage_transaction by id
// @Description  get storage_transaction by id
// @Tags         storage_transaction
// @Accept       json
// @Produce      json
// @Param 		 id path string true "storage_transaction_id"
// @Success      200  {object}  models.StorageTransaction
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetStorageTransactionByID(c *gin.Context) {
	 
	uid := c.Param("id")

	stTransaction, err := h.Store.StorageTransaction().GetByID(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting storage transaction by id!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, stTransaction)
}

// GetStorageTransactionList godoc
// @Router       /storage_transactions [GET]
// @Summary      Get storage_transaction list
// @Description  get storage_transaction list
// @Tags         storage_transaction
// @Accept       json
// @Produce      json
// @Param 		 page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.StorageTransactionsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetStorageTransactionList(c *gin.Context) {
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

	resp, err := h.Store.StorageTransaction().GetList(context.Background(),models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

		 
	if err != nil {
		handleResponse(c, "Error in handlers, while getting storage transactions!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, resp)
}

// UpdateStorageTransactions godoc
// @Router       /storage_transaction/{id} [PUT]
// @Summary      Update storage_transaction
// @Description  update storage_transaction
// @Tags         storage_transaction
// @Accept       json
// @Produce      json
// @Param 		 id path string true "storage_transaction_id"
// @Param 		 storage_transactions body models.UpdateStorageTransaction false "storage_transactions"
// @Success      200  {object}  models.StorageTransaction
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateStorageTransaction(c *gin.Context) {
	updateStTransaction := models.UpdateStorageTransaction{}

	uid := c.Param("id")

	if uid == ""{
		handleResponse(c, "invalid uuid!", http.StatusBadRequest,errors.New("uuid is not valid"))
		return
	 }
 
	 updateStTransaction.ID = uid

	 err := c.ShouldBindJSON(&updateStTransaction)
	 if err != nil{
		handleResponse(c,"Error in handlers, while reading storage transactions json from client!",http.StatusBadRequest,err)
		return
	 }

	pKey, err := h.Store.StorageTransaction().Update(context.Background(),updateStTransaction)
	if err != nil {
		handleResponse(c, "Error in handlers, while updating storage transaction!",http.StatusInternalServerError, err.Error())
		return
	}

	stTransaction, err := h.Store.StorageTransaction().GetByID(context.Background(),models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while selecting staff by id!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, stTransaction)
}

// DeleteStorageTransaction godoc
// @Router       /storage_transaction/{id} [DELETE]
// @Summary      Delete storage_transaction
// @Description  delete storage_transaction
// @Tags         storage_transaction
// @Accept       json
// @Produce      json
// @Param 		 id path string true "storage_transaction_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteStorageTransaction(c *gin.Context) {
	uid := c.Param("id") 

	err := h.Store.StorageTransaction().Delete(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil{
		handleResponse(c, "Error in handlers, while deleting storage transaction!",http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "",http.StatusOK, "data successfully deleted")
}
