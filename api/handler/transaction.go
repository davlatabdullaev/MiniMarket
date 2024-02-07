package handler

import (
	"context"
	"developer/api/models"
	"errors"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// CreateTransaction godoc
// @Router       /transaction [POST]
// @Summary      Create a new transaction
// @Description  create a new transaction
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param 		 transaction body models.CreateTransaction false "transaction"
// @Success      200  {object}  models.Transaction
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateTransaction(c *gin.Context) {
	createTransaction := models.CreateTransaction{}

	err := c.ShouldBindJSON(&createTransaction)
	if err != nil{
		handleResponse(c,"Error in handlers, while reading transaction json!",http.StatusBadRequest,err.Error())
	}

	pKey, err := h.Store.Transaction().Create(context.Background(),createTransaction)
	if err != nil {
		handleResponse(c, "Error in handlers, while creating transaction!",http.StatusInternalServerError, err.Error())
		return
	}

	transaction, err := h.Store.Transaction().GetByID(context.Background(),models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error while getting transaction by id!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusCreated, transaction)
}

// GetTransaction godoc
// @Router       /transaction/{id} [GET]
// @Summary      Get transaction by id
// @Description  get transaction by id
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param 		 id path string true "transaction_id"
// @Success      200  {object}  models.Transaction
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetTransactionByID(c *gin.Context) {
	
	uid := c.Param("id")

	transaction, err := h.Store.Transaction().GetByID(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting transaction by id!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, transaction)
}

// GetTransactionList godoc
// @Router       /transactions [GET]
// @Summary      Get transactions list
// @Description  get transactions list
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param 		 page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.TransactionsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetTransactionList(c *gin.Context) {
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

	resp, err := h.Store.Transaction().GetList(context.Background(),models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

		 
	if err != nil {
		handleResponse(c, "Error  in handlers, while getting transactions!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, resp)
}

// UpdateTransaction godoc
// @Router       /transaction/{id} [PUT]
// @Summary      Update transaction
// @Description  update transaction
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param 		 id path string true "transaction_id"
// @Param 		 transaction body models.UpdateTransaction false "transaction"
// @Success      200  {object}  models.Transaction
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateTransaction(c *gin.Context) {
	updateTransaction := models.UpdateTransaction{}

	uid := c.Param("id")

	if uid == ""{
		handleResponse(c, "invalid uuid!", http.StatusBadRequest,errors.New("uuid is not valid"))
		return
	 }
 
	 updateTransaction.ID = uid

	 err := c.ShouldBindJSON(&updateTransaction)
	 if err != nil{
		handleResponse(c,"Error in handlers, while reading transaction json from client!",http.StatusBadRequest,err)
		return
	 }

	pKey, err := h.Store.Transaction().Update(context.Background(),updateTransaction)
	if err != nil {
		handleResponse(c, "Error in handlers, while updating transaction!",http.StatusInternalServerError, err.Error())
		return
	}

	transaction, err := h.Store.Transaction().GetByID(context.Background(),models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while selecting transaction by id!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, transaction)
}

// DeleteTransaction godoc
// @Router       /transaction/{id} [DELETE]
// @Summary      Delete transaction
// @Description  delete transaction
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param 		 id path string true "transaction_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteTransaction(c *gin.Context) {
	uid := c.Param("id") 

	err := h.Store.Transaction().Delete(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil{
		handleResponse(c,"Error in handlers, while deleting transaction!",http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "",http.StatusOK, "data successfully deleted")
}
