package handler

import (
	"context"
	"developer/api/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @Router       /product [POST]
// @Summary      Create a new product
// @Description  create a new product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 product body models.CreateProduct false "product"
// @Success      200  {object}  models.Product
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateProduct(c *gin.Context) {
	createProd := models.CreateProduct{}
	
	err := c.ShouldBindJSON(&createProd)

	if err != nil{
		handleResponse(c,"Error in handlers, while reading json product!",http.StatusBadRequest, err.Error())
	}

	pKey, err := h.Store.Product().Create(context.Background(),createProd)
	if err != nil {
		handleResponse(c, "Error in handlers, while creating product!",http.StatusInternalServerError, err.Error())
		return
	}

	product, err := h.Store.Product().GetByID(context.Background(),models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting product by id!" ,http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusCreated, product)
}

// GetProduct godoc
// @Router       /product/{id} [GET]
// @Summary      Get product by id
// @Description  get product by id
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 id path string true "product_id"
// @Success      200  {object}  models.Product
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetProduct(c *gin.Context) {
	uid := c.Param("id")

	product, err := h.Store.Product().GetByID(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting product by id!" ,http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, product)
}

// GetProductList godoc
// @Router       /products [GET]
// @Summary      Get product list
// @Description  get product list
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.ProductsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetProductList(c *gin.Context) {
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

	resp, err := h.Store.Product().GetList(context.Background(),models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

		 
	if err != nil {
		handleResponse(c, "Error in handlers, while getting list of products!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, resp)
}

// UpdateProduct godoc
// @Router       /product/{id} [PUT]
// @Summary      Update product
// @Description  update product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 id path string true "product_id"
// @Param 		 product body models.UpdateProduct false "product"
// @Success      200  {object}  models.Product
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateProduct(c *gin.Context) {
	updateProd := models.UpdateProduct{}

	uid := c.Param("id")

	if uid == ""{
		handleResponse(c, "invalid uuid!", http.StatusBadRequest,errors.New("uuid is not valid"))
		return
	 }
 
	 updateProd.ID = uid

	 err := c.ShouldBindJSON(&updateProd)
	 if err != nil{
		handleResponse(c,"Error in handlers, while reading product json from client!",http.StatusBadRequest,err)
		return
	 }

	pKey, err := h.Store.Product().Update(context.Background(),updateProd)
	if err != nil {
		handleResponse(c, "Error in handlers, while updating product!",http.StatusInternalServerError, err.Error())
		return
	}

	product, err := h.Store.Product().GetByID(context.Background(),models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error in hanlders, while getting product by id!",http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "",http.StatusOK, product)
}

// DeleteProduct godoc
// @Router       /product/{id} [DELETE]
// @Summary      Delete product
// @Description  delete product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 id path string true "product_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteProduct(c *gin.Context) {
	 	uid := c.Param("id")

	err := h.Store.Product().Delete(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil{
		handleResponse(c, "Error in handlers, while deleting product!",http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "",http.StatusOK, "data successfully deleted")
}
