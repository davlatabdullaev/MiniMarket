package handler

import (
	"context"
	"developer/api/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


// CreateBranch godoc
// @Router       /branch [POST]
// @Summary      Creates a new branch
// @Description  create a new branch
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param        branch body models.CreateBranch false "branch"
// @Success      201  {object}  models.Branch
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateBranch(c *gin.Context) {
	createBranch := models.CreateBranch{}
	
	err := c.ShouldBindJSON(&createBranch)
	if err != nil{
		handleResponse(c,"Error in handlers, while reading branch!",http.StatusBadRequest,err.Error())
	}

	pKey, err := h.Store.Branch().Create(context.Background(),createBranch)
	if err != nil {
		handleResponse(c, "Error in handlers, while creating branch!" ,http.StatusInternalServerError, err)
		return
	}

	branch, err := h.Store.Branch().GetByID(context.Background(),models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting branch by id!",http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "" ,http.StatusCreated, branch)
}

// GetBranch godoc
// @Router       /branch/{id} [GET]
// @Summary      Get branch by id
// @Description  get branch by id
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param        id path string true "branch_id"
// @Success      201  {object}  models.Branch
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBranchByID(c *gin.Context) {
	 uid := c.Param("id")

	branch, err := h.Store.Branch().GetByID(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting branch by id!" ,http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "",http.StatusOK, branch)
}

// GetBranchList godoc
// @Router       /branches [GET]
// @Summary      Get branch list
// @Description  get branch list
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      201  {object}  models.BranchesResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBranchList(c *gin.Context) {
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

	

	resp, err := h.Store.Branch().GetList(context.Background(),models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

		 
	if err != nil {
		handleResponse(c, "Error in handlers, while getting branches!",http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "",http.StatusOK, resp)
}

// UpdateBranch godoc
// @Router       /branch/{id} [PUT]
// @Summary      Update branch
// @Description  update branch
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param        id path string true "branch_id"
// @Param        branch body models.UpdateBranch false "branch"
// @Success      201  {object}  models.Branch
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateBranch(c *gin.Context) {
	updateBranch := models.UpdateBranch{}

	uid := c.Param("id")

	if uid == ""{
		handleResponse(c, "invalid uuid!", http.StatusBadRequest,errors.New("uuid is not valid"))
		return
	 }
 
	 updateBranch.ID = uid

	 err := c.ShouldBindJSON(&updateBranch)
	 if err != nil{
		handleResponse(c,"Error in handlers, while reading branch json from client!",http.StatusBadRequest,err)
		return
	 }

	pKey, err := h.Store.Branch().Update(context.Background(),updateBranch)
	if err != nil {
		handleResponse(c, "Error in handlers, while updating branch!",http.StatusInternalServerError, err.Error())
		return
	}

	branch, err := h.Store.Branch().GetByID(context.Background(),models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting branch!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, branch)
}

// DeleteBranch godoc
// @Router       /branch/{id} [Delete]
// @Summary      Delete branch
// @Description  delete branch
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param        id path string true "branch_id"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteBranch(c *gin.Context) {
	uid := c.Param("id")

	err := h.Store.Branch().Delete(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil{
		handleResponse(c, "Error in handlers, while deleting branch!",http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "",http.StatusOK, "data successfully deleted")
}
