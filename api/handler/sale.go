package handler

import (
	"context"
	"developer/api/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateSale godoc
// @Router       /sale [POST]
// @Summary      Create a new sale
// @Description  create a new sale
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param 		 sale body models.CreateSale false "sale"
// @Success      200  {object}  models.Sale
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateSale(c *gin.Context) {
	createSale := models.CreateSale{}

	err := c.ShouldBindJSON(&createSale)
	if err != nil {
		handleResponse(c, "Error in handlers, while reading sale json!", http.StatusBadRequest, err.Error())
	}

	pKey, err := h.Store.Sale().Create(context.Background(), createSale)
	if err != nil {
		handleResponse(c, "Error in handlers, while creating sale!", http.StatusInternalServerError, err.Error())
		return
	}

	sale, err := h.Store.Sale().GetByID(context.Background(), models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error while getting sale by id!", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusCreated, sale)
}

// GetSale godoc
// @Router       /sale/{id} [GET]
// @Summary      Get sale by id
// @Description  get sale by id
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param 		 id path string true "sale_id"
// @Success      200  {object}  models.Sale
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetSaleByID(c *gin.Context) {

	uid := c.Param("id")

	sale, err := h.Store.Sale().GetByID(context.Background(), models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting sale by id!", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, sale)
}

// GetSaleList godoc
// @Router       /sales [GET]
// @Summary      Get sale list
// @Description  get sale list
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param 		 page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.SalesResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetSaleList(c *gin.Context) {
	var (
		page, limit = 1, 10
		err         error
		search      string
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

	resp, err := h.Store.Sale().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c, "Error  in handlers, while getting sales!", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

// UpdateSale godoc
// @Router       /sale/{id} [PUT]
// @Summary      Update sale
// @Description  update sale
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param 		 id path string true "sale_id"
// @Param 		 sale body models.UpdateSale false "sale"
// @Success      200  {object}  models.Sale
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateSale(c *gin.Context) {
	updateSale := models.UpdateSale{}

	uid := c.Param("id")

	if uid == "" {
		handleResponse(c, "invalid uuid!", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	updateSale.ID = uid

	err := c.ShouldBindJSON(&updateSale)
	if err != nil {
		handleResponse(c, "Error in handlers, while reading sale json from client!", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.Store.Sale().Update(context.Background(), updateSale)
	if err != nil {
		handleResponse(c, "Error in handlers, while updating sale!", http.StatusInternalServerError, err.Error())
		return
	}

	sale, err := h.Store.Sale().GetByID(context.Background(), models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while selecting sale by id!", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, sale)
}

// DeleteSale godoc
// @Router       /sale/{id} [DELETE]
// @Summary      Delete sale
// @Description  delete sale
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param 		 id path string true "sale_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteSale(c *gin.Context) {
	uid := c.Param("id")

	err := h.Store.Sale().Delete(context.Background(), models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while deleting sale!", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "", http.StatusOK, "data successfully deleted")
}

// NEW

// StartSale godoc
// @Router       /start_sale [POST]
// @Summary      Starting sale
// @Description  starting sale
// @Tags         start_sale
// @Accept       json
// @Produce      json
// @Param 		 start_sale body []models.StartSale false "start_sale"
// @Success      200  {object}  models.Sale
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) StartSale(c *gin.Context) {
	totalSum := 0

	startSale := []models.StartSale{}
	createSale := models.CreateSale{
		BranchID: "5b07248c-c631-4489-9dcb-3d0a4fa08917",
		ShopAssistantID: "abcdefghij",
		CashierID:   "123456789",
		PaymentType: "cash",
		Price:       0,
		Status:      "in_procces",
		ClientName:  "Salohiddin",
	}

	err := c.ShouldBindJSON(&startSale)
	if err != nil{
		handleResponse(c, "Error in handlers, reading start sale json from clients!", http.StatusBadRequest, err)
		return
	}

	saleID, err := h.Store.Sale().Create(context.Background(), createSale)
	if err != nil {
		handleResponse(c, "Error in handlers, while creating sale for start sell!", http.StatusInternalServerError, err)
		return
	}

	for _, sale := range startSale{
	product, err := h.Store.Product().GetByID(context.Background(), models.PrimaryKey{
		ID: sale.ProductId,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting product by id!", http.StatusInternalServerError, err)
		return
	}

	createBasket := models.CreateBasket{
		SaleID:    saleID,
		ProductID: sale.ProductId,
		Quantity:   sale.Quantity,
		Price:     uint(product.Price * sale.Quantity),
	}

	totalSum = totalSum + int(createBasket.Price)

	_, err = h.Store.Basket().Create(context.Background(), createBasket)
	if err != nil {
		handleResponse(c, "Error in handlers, while creating basket!", http.StatusInternalServerError, err)
		return
	}
}
	saleId, err := h.Store.Sale().Update(context.Background(), models.UpdateSale{
		ID: saleID,
		BranchID: createSale.BranchID,
		ShopAssistantID: "abcdefghij",
		CashierID: "123456789",
		PaymentType: createSale.PaymentType,
		Price:  uint(totalSum),
		Status: "succes",
		ClientName: createSale.ClientName,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while updating sale by id!", http.StatusInternalServerError, err)
		return
	}

	sale, err := h.Store.Sale().GetByID(context.Background(), models.PrimaryKey{ID: saleId})
	if err != nil{
		handleResponse(c, "Error in handlers, while getting sale by id!", http.StatusInternalServerError,err)
		return
	}

	handleResponse(c, "success", http.StatusOK, sale)

}
