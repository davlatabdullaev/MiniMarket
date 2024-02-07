package handler

import (
	"context"
	"developer/api/models"
	"errors"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// CreateStaff godoc
// @Router       /staff [POST]
// @Summary      Create a new staff
// @Description  create a new staff
// @Tags         staff
// @Accept       json
// @Produce      json
// @Param 		 staff body models.CreateStaff false "staff"
// @Success      200  {object}  models.Staff
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateStaff(c *gin.Context) {
	createStaff := models.CreateStaff{}

	err := c.ShouldBindJSON(&createStaff)
	if err != nil{
		handleResponse(c,"Error in handlers, while reading staff json!",http.StatusBadRequest,err.Error())
	}

	pKey, err := h.Store.Staff().Create(context.Background(),createStaff)
	if err != nil {
		handleResponse(c, "Error in handlers, while creating staff!",http.StatusInternalServerError, err.Error())
		return
	}

	staff, err := h.Store.Staff().GetByID(context.Background(),models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error while getting staff by id!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusCreated, staff)
}

// GetStaff godoc
// @Router       /staff/{id} [GET]
// @Summary      Get staff by id
// @Description  get staff by id
// @Tags         staff
// @Accept       json
// @Produce      json
// @Param 		 id path string true "staff_id"
// @Success      200  {object}  models.Staff
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetStaffByID(c *gin.Context) {
	 
	uid := c.Param("id")

	staff, err := h.Store.Staff().GetByID(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting staff by id!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, staff)
}

// GetStaffList godoc
// @Router       /staffs [GET]
// @Summary      Get staff list
// @Description  get staff list
// @Tags         staff
// @Accept       json
// @Produce      json
// @Param 		 page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.StaffsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetStaffList(c *gin.Context) {
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

	resp, err := h.Store.Staff().GetList(context.Background(),models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

		 
	if err != nil {
		handleResponse(c, "Error  in handlers, while getting staffs!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, resp)
}

// UpdateStaff godoc
// @Router       /staff/{id} [PUT]
// @Summary      Update staff
// @Description  update staff
// @Tags         staff
// @Accept       json
// @Produce      json
// @Param 		 id path string true "staff_id"
// @Param 		 staff body models.UpdateStaff false "staff"
// @Success      200  {object}  models.Staff
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateStaff(c *gin.Context) {
	updateStaff := models.UpdateStaff{}

	uid := c.Param("id")

	if uid == ""{
		handleResponse(c, "invalid uuid!", http.StatusBadRequest,errors.New("uuid is not valid"))
		return
	 }
 
	 updateStaff.ID = uid

	 err := c.ShouldBindJSON(&updateStaff)
	 if err != nil{
		handleResponse(c,"Error in handlers, while reading staff json from client!",http.StatusBadRequest,err)
		return
	 }

	pKey, err := h.Store.Staff().Update(context.Background(),updateStaff)
	if err != nil {
		handleResponse(c, "Error in handlers, while updating staff!",http.StatusInternalServerError, err.Error())
		return
	}

	staff, err := h.Store.Staff().GetByID(context.Background(),models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while selecting staff by id!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, staff)
}

// DeleteStaff godoc
// @Router       /staff/{id} [DELETE]
// @Summary      Delete staff
// @Description  delete staff
// @Tags         staff
// @Accept       json
// @Produce      json
// @Param 		 id path string true "staff_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteStaff(c *gin.Context) {
	uid := c.Param("id") 

	err := h.Store.Staff().Delete(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil{
		handleResponse(c, "Error in handlers, while deleting staff!",http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "",http.StatusOK, "data successfully deleted")
}
