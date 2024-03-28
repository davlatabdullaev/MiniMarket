package handler

import (
	"context"
	"developer/api/models"
	"errors"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// CreateTarif godoc
// @Router       /tarif [POST]
// @Summary      Create a new tarif
// @Description  create a new tarif
// @Tags         tarif
// @Accept       json
// @Produce      json
// @Param 		 tarif body models.CreateTarif false "tarif"
// @Success      200  {object}  models.Tarif
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateTarif(c *gin.Context) {
	createTarif := models.CreateTarif{}

	err := c.ShouldBindJSON(&createTarif)
	if err != nil{
		handleResponse(c,"Error in handlers, while reading tarif json!",http.StatusBadRequest,err.Error())
	}

	pKey, err := h.Store.Tarif().Create(context.Background(),createTarif)
	if err != nil {
		handleResponse(c, "Error in handlers, while creating tarif!",http.StatusInternalServerError, err.Error())
		return
	}

	tarif, err := h.Store.Tarif().GetByID(context.Background(),models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error while getting tarif by id!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusCreated, tarif)
}

// GetTarif godoc
// @Router       /tarif/{id} [GET]
// @Summary      Get tarif by id
// @Description  get tarif by id
// @Tags         tarif
// @Accept       json
// @Produce      json
// @Param 		 id path string true "tarif_id"
// @Success      200  {object}  models.Tarif
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetTarif(c *gin.Context) {
	
	uid := c.Param("id")

	tarif, err := h.Store.Tarif().GetByID(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting tarif by id!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, tarif)
}

// GetTarifList godoc
// @Router       /tarifs [GET]
// @Summary      Get tarifs list
// @Description  get tarifs list
// @Tags         tarifs
// @Accept       json
// @Produce      json
// @Param 		 page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.TarifsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetTarifList(c *gin.Context) {
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

	resp, err := h.Store.Tarif().GetList(context.Background(),models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

		 
	if err != nil {
		handleResponse(c, "Error  in handlers, while getting tarifs!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, resp)
}

// UpdateTarif godoc
// @Router       /tarif/{id} [PUT]
// @Summary      Update tarif
// @Description  update tarif
// @Tags         tarif
// @Accept       json
// @Produce      json
// @Param 		 id path string true "tarif_id"
// @Param 		 tarif body models.UpdateTarif false "tarif"
// @Success      200  {object}  models.Tarif
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateTarif(c *gin.Context) {
	updateTarif := models.UpdateTarif{}

	uid := c.Param("id")

	if uid == ""{
		handleResponse(c, "invalid uuid!", http.StatusBadRequest,errors.New("uuid is not valid"))
		return
	 }
 
	 updateTarif.ID = uid

	 err := c.ShouldBindJSON(&updateTarif)
	 if err != nil{
		handleResponse(c,"Error in handlers, while reading tarif json from client!",http.StatusBadRequest,err)
		return
	 }

	pKey, err := h.Store.Tarif().Update(context.Background(),updateTarif)
	if err != nil {
		handleResponse(c, "Error in handlers, while updating tarif!",http.StatusInternalServerError, err.Error())
		return
	}

	tarif, err := h.Store.Tarif().GetByID(context.Background(),models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while selecting tarif by id!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "",http.StatusOK, tarif)
}

// DeleteTarif godoc
// @Router       /tarif/{id} [DELETE]
// @Summary      Delete tarif
// @Description  delete tarif
// @Tags         tarif
// @Accept       json
// @Produce      json
// @Param 		 id path string true "tarif_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteTarif(c *gin.Context) {
	uid := c.Param("id") 

	err := h.Store.Tarif().Delete(context.Background(),models.PrimaryKey{
		ID: uid,
	})
	if err != nil{
		handleResponse(c,"Error in handlers, while deleting tarif!",http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "",http.StatusOK, "data successfully deleted")
}
