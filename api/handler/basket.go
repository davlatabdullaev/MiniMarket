package handler

import (
	"context"
	"developer/api/models"
	"errors"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)


// CreateBasket godoc
// @Router       /basket [POST]
// @Summary      Creates a new basket
// @Description  create a new basket
// @Tags         basket
// @Accept       json
// @Produce      json
// @Param        basket body models.CreateBasket false "basket"
// @Success      201  {object}  models.Basket
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateBasket(c *gin.Context) {
	createBasket := models.CreateBasket{}
	
	err := c.ShouldBindHeader(&createBasket)
	if err != nil{
		handleResponse(c,"Error in handler, while reading basket from json!",http.StatusBadRequest, err.Error())
	}

	pKey, err := h.Store.Basket().Create(context.Background(),createBasket)
	if err != nil {
		handleResponse(c,"Error in handlers, while creating basket!", http.StatusInternalServerError, err)
		return
	}

	basket, err := h.Store.Basket().GetByID(context.Background(),models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting user by id!" ,http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "",http.StatusCreated, basket)
}

// GetBasket godoc
// @Router       /basket/{id} [GET]
// @Summary      Get basket by id
// @Description  get basket by id
// @Tags         basket
// @Accept       json
// @Produce      json
// @Param        id path string true "basket_id"
// @Success      201  {object}  models.Basket
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBasketByID(c *gin.Context) {
	 
	uid := c.Param("id")

	basket, err := h.Store.Basket().GetByID(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting user by id" ,http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "" ,http.StatusOK, basket)
}

// GetBasketList godoc
// @Router       /baskets [GET]
// @Summary      Get basket list
// @Description  get basket list
// @Tags         basket
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Success      201  {object}  models.BasketsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBasketList(c *gin.Context) {
	var (
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
	resp, err := h.Store.Basket().GetList(context.Background(),models.GetListRequest{
		Page:   page,
		Limit:  limit,
	})

		 
	if err != nil {
		handleResponse(c, "Error in handlers, while getting baskets, " ,http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "",http.StatusOK, resp)
}

// UpdateBasket godoc
// @Router       /basket/{id} [PUT]
// @Summary      Update basket
// @Description  update basket
// @Tags         basket
// @Accept       json
// @Produce      json
// @Param        id path string true "basket_id"
// @Param        basket body models.UpdateBasket false "basket"
// @Success      201  {object}  models.Basket
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateBasket(c *gin.Context) {
	updateBasket := models.UpdateBasket{}

	uid := c.Param("id")
	if uid == ""{
	   handleResponse(c, "invalid uuid!", http.StatusBadRequest,errors.New("uuid is not valid"))
	   return
	}

	updateBasket.ID = uid

	err := c.ShouldBindJSON(&updateBasket)
	if err != nil{
		handleResponse(c,"Error in handlers, while reading body for udatebasket!", http.StatusBadRequest,err.Error())
	}

	pKey, err := h.Store.Basket().Update(context.Background(),updateBasket)
	if err != nil {
		handleResponse(c, "Error in handlers, while updating basket!",http.StatusInternalServerError, err.Error())
		return
	}

	basket, err := h.Store.Basket().GetByID(context.Background(),models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting basket by id!" ,http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, basket)
}

// DeleteBasket godoc
// @Router       /basket/{id} [Delete]
// @Summary      Delete basket
// @Description  delete basket
// @Tags         basket
// @Accept       json
// @Produce      json
// @Param        id path string true "basket_id"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteBasket(c *gin.Context) {
	uid := c.Param("id")

	err := h.Store.Basket().Delete(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil{
		handleResponse(c, "Error in handlers, while deleting basket!",http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c,"", http.StatusOK, "data successfully deleted")
}
