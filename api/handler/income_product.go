package handler

import (
	"context"
	"developer/api/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


// CreateIncomeProduct godoc
// @Router       /income_product [POST]
// @Summary      Creates a new income product
// @Description  create a new income product
// @Tags         income_product
// @Accept       json
// @Produce      json
// @Param        income_product body models.CreateIncomeProduct false "income_product"
// @Success      201  {object}  models.IncomeProduct
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateIncomeProduct(c *gin.Context) {
	createIncomeProduct := models.CreateIncomeProduct{}
	
	err := c.ShouldBindJSON(&createIncomeProduct)
	if err != nil{
		handleResponse(c,"Error in handlers, while reading incomeProduct!",http.StatusBadRequest,err.Error())
	}

	pKey, err := h.Store.IncomeProduct().Create(context.Background(),createIncomeProduct)
	if err != nil {
		handleResponse(c, "Error in handlers, while creating incomeProduct!" ,http.StatusInternalServerError, err)
		return
	}

	incomeProduct, err := h.Store.IncomeProduct().GetByID(context.Background(),models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting incomeProduct by id!",http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "" ,http.StatusCreated, incomeProduct)
}

// GetIncomeProduct godoc
// @Router       /income_product/{id} [GET]
// @Summary      Get income product by id
// @Description  get income product by id
// @Tags         income_product
// @Accept       json
// @Produce      json
// @Param        id path string true "income_product_id"
// @Success      201  {object}  models.IncomeProduct
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetIncomeProduct(c *gin.Context) {
	 uid := c.Param("id")

	 incomeProduct, err := h.Store.IncomeProduct().GetByID(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting incomeProduct by id!" ,http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "",http.StatusOK, incomeProduct)
}

// GetIncomeProductsList godoc
// @Router       /income_products [GET]
// @Summary      Get income product list
// @Description  get income product list
// @Tags         income_product
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      201  {object}  models.IncomeProductsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetIncomeProductsList(c *gin.Context) {
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

	

	resp, err := h.Store.IncomeProduct().GetList(context.Background(),models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

		 
	if err != nil {
		handleResponse(c, "Error in handlers, while getting incomeProducts!",http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "",http.StatusOK, resp)
}

// UpdateIncomeProduct godoc
// @Router       /income_product/{id} [PUT]
// @Summary      Update income product
// @Description  update income product
// @Tags         income_product
// @Accept       json
// @Produce      json
// @Param        id path string true "income_product_id"
// @Param        income_product body models.UpdateIncomeProduct false "income_product"
// @Success      201  {object}  models.IncomeProduct
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateincomeProduct(c *gin.Context) {
	updIncomeProduct := models.UpdateIncomeProduct{}

	uid := c.Param("id")

	if uid == ""{
		handleResponse(c, "invalid uuid!", http.StatusBadRequest,errors.New("uuid is not valid"))
		return
	 }
 
	 updIncomeProduct.ID = uid

	 err := c.ShouldBindJSON(&updIncomeProduct)
	 if err != nil{
		handleResponse(c,"Error in handlers, while reading incomeProduct json from client!",http.StatusBadRequest,err)
		return
	 }

	incomeProduct, err := h.Store.IncomeProduct().Update(context.Background(),updIncomeProduct)
	if err != nil {
		handleResponse(c, "Error in handlers, while updating incomeProduct!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, incomeProduct)
}

// DeleteIncomeProduct godoc
// @Router       /income_product/{id} [Delete]
// @Summary      Delete income_product
// @Description  delete income_product
// @Tags         income_product
// @Accept       json
// @Produce      json
// @Param        id path string true "income_product_id"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteIncomeProduct(c *gin.Context) {
	uid := c.Param("id")

	err := h.Store.IncomeProduct().Delete(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil{
		handleResponse(c, "Error in handlers, while deleting incomeProduct!",http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "",http.StatusOK, "data successfully deleted")
}
