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

	r.POST("/branch", h.CreateBranch)
	r.GET("/branch/:id", h.GetBranchByID)
	r.GET("/branches", h.GetBranchList)
	r.PUT("/branch/:id", h.UpdateBranch)
	r.DELETE("/branch/:id", h.DeleteBranch)

	r.POST("/product", h.CreateProduct)
	r.GET("/product/:id", h.GetProductByID)
	r.GET("/products", h.GetProductList)
	r.PUT("/product/:id", h.UpdateProduct)
	r.DELETE("/product/:id", h.DeleteProduct)

	r.POST("/basket", h.CreateBasket)
	r.GET("/basket/:id", h.GetBasketByID)
	r.GET("/baskets", h.GetBasketList)
	r.PUT("/basket/:id", h.UpdateBasket)
	r.DELETE("/basket/:id", h.DeleteBasket)

	r.POST("/sale", h.CreateSale)
	r.GET("/sale/:id", h.GetSaleByID)
	r.GET("/sales", h.GetSaleList)
	r.PUT("/sale/:id", h.UpdateSale)
	r.DELETE("/sale/:id", h.DeleteSale)

	r.POST("/storage", h.CreateStorage)
	r.GET("/storage/:id", h.GetStorageByID)
	r.GET("/storages", h.GetStorageList)
	r.PUT("/storage/:id", h.UpdateStorage)
	r.DELETE("/storage/:id", h.DeleteStorage)

	// New
	r.POST("/category", h.CreateCategory)
	r.GET("/category/:id", h.GetCategoryByID)
	r.GET("/categories", h.GetCategoryList)
	r.PUT("/category/:id", h.UpdateCategory)
	r.DELETE("/category/:id", h.DeleteCategory)

	r.POST("/tarif", h.CreateTarif)
	r.GET("/tarif/:id", h.GetTarifByID)
	r.GET("/tarifs", h.GetTarifList)
	r.PUT("/tarif/:id", h.UpdateTarif)
	r.DELETE("/tarif/:id", h.DeleteTarif)

	r.POST("/staff", h.CreateStaff)
	r.GET("/staff/:id", h.GetStaffByID)
	r.GET("/staff", h.GetStaffList)
	r.PUT("/staff/:id", h.UpdateStaff)
	r.DELETE("/staff/:id", h.DeleteStaff)

	r.POST("/transaction", h.CreateTransaction)
	r.GET("/transaction/:id", h.GetTransactionByID)
	r.GET("/transactions", h.GetTransactionList)
	r.PUT("/transaction/:id", h.UpdateTransaction)
	r.DELETE("/transaction/:id", h.DeleteTransaction)

	r.POST("/storage_transaction", h.CreateStorageTransaction)
	r.GET("/storage_transaction/:id", h.GetStorageTransactionByID)
	r.GET("/storage_transactions", h.GetStorageTransactionList)
	r.PUT("/storage_transaction/:id", h.UpdateStorageTransaction)
	r.DELETE("/storage_transaction/:id", h.DeleteStorageTransaction)

	r.POST("/start_sale", h.StartSale)


	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}