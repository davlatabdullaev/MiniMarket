package handler

import (
	"context"
	"developer/api/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


// CreateCategory godoc
// @Router       /category [POST]
// @Summary      Creates a new category
// @Description  create a new category
// @Tags         category
// @Accept       json
// @Produce      json
// @Param        category body models.CreateCategory false "category"
// @Success      201  {object}  models.Category
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateCategory(c *gin.Context) {
	createCategory := models.CreateCategory{}
	
	err := c.ShouldBindJSON(&createCategory)
	if err != nil{
		handleResponse(c,"Error in handlers, while reading category!",http.StatusBadRequest,err.Error())
	}

	pKey, err := h.Store.Category().Create(context.Background(),createCategory)
	if err != nil {
		handleResponse(c, "Error in handlers, while creating category!" ,http.StatusInternalServerError, err)
		return
	}

	category, err := h.Store.Category().GetByID(context.Background(),models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting category by id!",http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "" ,http.StatusCreated, category)
}

// GetCategory godoc
// @Router       /category/{id} [GET]
// @Summary      Get category by id
// @Description  get category by id
// @Tags         category
// @Accept       json
// @Produce      json
// @Param        id path string true "category_id"
// @Success      201  {object}  models.Category
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetCategoryByID(c *gin.Context) {
	 uid := c.Param("id")

	 category, err := h.Store.Category().GetByID(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting category by id!" ,http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "",http.StatusOK, category)
}

// GetCategoryList godoc
// @Router       /categories [GET]
// @Summary      Get category list
// @Description  get category list
// @Tags         category
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      201  {object}  models.CategoriesResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetCategoryList(c *gin.Context) {
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

	

	resp, err := h.Store.Category().GetList(context.Background(),models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

		 
	if err != nil {
		handleResponse(c, "Error in handlers, while getting categories!",http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "",http.StatusOK, resp)
}

// UpdateCategory godoc
// @Router       /category/{id} [PUT]
// @Summary      Update category
// @Description  update category
// @Tags         category
// @Accept       json
// @Produce      json
// @Param        id path string true "category_id"
// @Param        category body models.UpdateCategory false "category"
// @Success      201  {object}  models.Category
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateCategory(c *gin.Context) {
	updateCategory := models.UpdateCategory{}

	uid := c.Param("id")

	if uid == ""{
		handleResponse(c, "invalid uuid!", http.StatusBadRequest,errors.New("uuid is not valid"))
		return
	 }
 
	 updateCategory.ID = uid

	 err := c.ShouldBindJSON(&updateCategory)
	 if err != nil{
		handleResponse(c,"Error in handlers, while reading category json from client!",http.StatusBadRequest,err)
		return
	 }

	pKey, err := h.Store.Category().Update(context.Background(),updateCategory)
	if err != nil {
		handleResponse(c, "Error in handlers, while updating category!",http.StatusInternalServerError, err.Error())
		return
	}

	category, err := h.Store.Category().GetByID(context.Background(),models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting category!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, category)
}

// DeleteCategory godoc
// @Router       /category/{id} [Delete]
// @Summary      Delete category
// @Description  delete category
// @Tags         category
// @Accept       json
// @Produce      json
// @Param        id path string true "category_id"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteCategory(c *gin.Context) {
	uid := c.Param("id")

	err := h.Store.Category().Delete(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil{
		handleResponse(c, "Error in handlers, while deleting category!",http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "",http.StatusOK, "data successfully deleted")
}
