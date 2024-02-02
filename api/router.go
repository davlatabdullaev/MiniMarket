package api

import (
	"developer/api/handler"
	"developer/storage"
	_"developer/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
func New(store storage.IfStorage) *gin.Engine {
	h := handler.New(store)

	r := gin.New()

	// r.POST("/branch", h.CreateBranch)
	// r.GET("/branch/:id", h.GetBranchByID)
	// r.GET("/branches", h.GetBranchList)
	// r.PUT("/branch/:id", h.UpdateBranch)
	// r.DELETE("/branch/:id", h.DeleteBranch)

	// r.POST("/product", h.CreateProduct)
	// r.GET("/product/:id", h.GetProductByID)
	// r.GET("/products", h.GetProductList)
	// r.PUT("/product/:id", h.UpdateProduct)
	// r.DELETE("/product/:id", h.DeleteProduct)

	r.POST("/basket", h.CreateBasket)
	r.GET("/basket/:id", h.GetBasketByID)
	r.GET("/baskets", h.GetBasketList)
	r.PUT("/basket/:id", h.UpdateBasket)
	r.DELETE("/basket/:id", h.DeleteBasket)

	// r.POST("/sale", h.CreateSale)
	// r.GET("/sale/:id", h.GetSaleByID)
	// r.GET("/sales", h.GetSaleList)
	// r.PUT("/sale/:id", h.UpdateSale)
	// r.DELETE("/sale/:id", h.DeleteSale)

	// r.POST("/storage", h.CreateStorage)
	// r.GET("/storage/:id", h.GetStorageByID)
	// r.GET("/storages", h.GetStorageList)
	// r.PUT("/storage/:id", h.UpdateStorage)
	// r.DELETE("/storage/:id", h.DeleteStorage)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}