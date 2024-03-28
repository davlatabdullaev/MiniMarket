package handler

import (
	"context"
	"developer/api/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


// CreateIncome godoc
// @Router       /income [POST]
// @Summary      Creates a new income
// @Description  create a new income
// @Tags         income
// @Accept       json
// @Produce      json
// @Param        income body models.CreateIncome false "income"
// @Success      201  {object}  models.Income
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateIncome(c *gin.Context) {
	createIncome := models.CreateIncome{}
	
	err := c.ShouldBindJSON(&createIncome)
	if err != nil{
		handleResponse(c,"Error in handlers, while reading income!",http.StatusBadRequest,err.Error())
	}

	pKey, err := h.Store.Income().Create(context.Background(),createIncome)
	if err != nil {
		handleResponse(c, "Error in handlers, while creating income!" ,http.StatusInternalServerError, err)
		return
	}

	income, err := h.Store.Income().GetByID(context.Background(),models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting income by id!",http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "" ,http.StatusCreated, income)
}

// GetIncome godoc
// @Router       /income/{id} [GET]
// @Summary      Get income by id
// @Description  get income by id
// @Tags         income
// @Accept       json
// @Produce      json
// @Param        id path string true "income_id"
// @Success      201  {object}  models.Income
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetIncomeByID(c *gin.Context) {
	 uid := c.Param("id")

	 income, err := h.Store.Income().GetByID(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting income by id!" ,http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "",http.StatusOK, income)
}

// GetIncomeList godoc
// @Router       /incomes [GET]
// @Summary      Get income list
// @Description  get income list
// @Tags         income
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      201  {object}  models.IncomesResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetIncomeList(c *gin.Context) {
	var (
		search string
		page, limit = 1, 10
		err         error
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

	

	resp, err := h.Store.Income().GetList(context.Background(),models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

		 
	if err != nil {
		handleResponse(c, "Error in handlers, while getting incomes!",http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "",http.StatusOK, resp)
}

// UpdateIncome godoc
// @Router       /income/{id} [PUT]
// @Summary      Update income
// @Description  update income
// @Tags         income
// @Accept       json
// @Produce      json
// @Param        id path string true "income_id"
// @Param        income body models.UpdateIncome false "income"
// @Success      201  {object}  models.Income
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) Updateincome(c *gin.Context) {
	updIncome := models.UpdateIncome{}

	uid := c.Param("id")

	if uid == ""{
		handleResponse(c, "invalid uuid!", http.StatusBadRequest,errors.New("uuid is not valid"))
		return
	 }
 
	 updIncome.ID = uid

	 err := c.ShouldBindJSON(&updIncome)
	 if err != nil{
		handleResponse(c,"Error in handlers, while reading income json from client!",http.StatusBadRequest,err)
		return
	 }

	income, err := h.Store.Income().Update(context.Background(),updIncome)
	if err != nil {
		handleResponse(c, "Error in handlers, while updating income!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, income)
}

// DeleteIncome godoc
// @Router       /income/{id} [Delete]
// @Summary      Delete income
// @Description  delete income
// @Tags         income
// @Accept       json
// @Produce      json
// @Param        id path string true "income_id"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteIncome(c *gin.Context) {
	uid := c.Param("id")

	err := h.Store.Income().Delete(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil{
		handleResponse(c, "Error in handlers, while deleting income!",http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "",http.StatusOK, "data successfully deleted")
}
