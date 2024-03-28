package storage

import (
	"context"
	"developer/api/models"
)



type IfStorage interface {
	Close()
	Branch() IBranch
	Sale()   ISale
	Basket() IBasket
	Product() IProduct
	Storage() IStorage
	Tarif() ITarif
	Category() ICategory
	Staff() IStaff
	Transaction() ITransaction
	StorageTransaction() IStorageTransaction
	Income() IIncome
	IncomeProduct() IIncomeProduct
}

type IBranch interface {
	Create(context.Context, models.CreateBranch)(string, error)
	GetByID(context.Context,models.PrimaryKey) (models.Branch, error)
	GetList(context.Context,models.GetListRequest) (models.BranchesResponse, error)
	Update(context.Context,models.UpdateBranch) (string, error)
	Delete(context.Context,models.PrimaryKey) error
}

type ISale interface {
	Create(context.Context,models.CreateSale)(string, error)
	GetByID(context.Context,models.PrimaryKey) (models.Sale, error)
	GetList(context.Context,models.GetListRequest) (models.SalesResponse, error)
	Update(context.Context,models.UpdateSale) (string, error)
	Delete(context.Context,models.PrimaryKey) error
	UpdateSalePrice(context.Context, models.UpdateSaleForPrice) error 
}

type IBasket interface {
	Create(context.Context,models.CreateBasket)(string, error)
	GetByID(context.Context,models.PrimaryKey) (models.Basket, error)
	GetList(context.Context,models.GetListRequest) (models.BasketsResponse, error)
	Update(context.Context,models.UpdateBasket) (string, error)
	Delete(context.Context,models.PrimaryKey) error
}
type IProduct interface {
	Create(context.Context,models.CreateProduct)(string, error)
	GetByID(context.Context,models.PrimaryKey) (models.Product, error)
	GetList(context.Context,models.GetListRequest) (models.ProductsResponse, error)
	Update(context.Context,models.UpdateProduct) (string, error)
	Delete(context.Context,models.PrimaryKey) error
	GetByBarcode(context.Context, string)(models.Product, error)
}

type IStorage interface {
	Create(context.Context,models.CreateStorage)(string, error)
	GetByID(context.Context,models.PrimaryKey) (models.Storage, error)
	GetList(context.Context,models.GetListRequest) (models.StoragesResponse, error)
	Update(context.Context,models.UpdateStorage) (string, error)
	Delete(context.Context,models.PrimaryKey) error
	GetByProductID(context.Context,models.PrimaryKey) (models.Storage, error)
	UpdateCount(context.Context, string, int)error
}

type ICategory interface {
	Create(context.Context,models.CreateCategory) (string, error)
	GetByID(context.Context,models.PrimaryKey) (models.Category, error)
	GetList(context.Context,models.GetListRequest) (models.CategoriesResponse, error)
	Update(context.Context,models.UpdateCategory) (string, error)
	Delete(context.Context,models.PrimaryKey) error
}

type IStaff interface {
	Create(context.Context,models.CreateStaff) (string, error)
	GetByID(context.Context,models.PrimaryKey) (models.Staff, error)
	GetList(context.Context,models.GetListRequest) (models.StaffsResponse, error)
	Update(context.Context,models.UpdateStaff) (string, error)
	Delete(context.Context,models.PrimaryKey) error
	UpdateSalary(context.Context,models.PrimaryKey, int) error
}

type IStorageTransaction interface {
	Create(context.Context,models.CreateStorageTransaction) (string, error)
	GetByID(context.Context,models.PrimaryKey) (models.StorageTransaction, error)
	GetList(context.Context,models.GetListRequest) (models.StorageTransactionsResponse, error)
	Update(context.Context,models.UpdateStorageTransaction) (string, error)
	Delete(context.Context,models.PrimaryKey) error
}

type ITarif interface {
	Create(context.Context,models.CreateTarif) (string, error)
	GetByID(context.Context,models.PrimaryKey) (models.Tarif, error)
	GetList(context.Context,models.GetListRequest) (models.TarifsResponse, error)
	Update(context.Context,models.UpdateTarif) (string, error)
	Delete(context.Context,models.PrimaryKey) error
}

type ITransaction interface {
	Create(context.Context,models.CreateTransaction) (string, error)
	GetByID(context.Context,models.PrimaryKey) (models.Transaction, error)
	GetList(context.Context,models.GetListRequest) (models.TransactionsResponse, error)
	Update(context.Context,models.UpdateTransaction) (string, error)
	Delete(context.Context,models.PrimaryKey) error
	UpdateStaffBalanceAndCreateTransaction(context.Context, models.UpdateStaffBalanceAndCreateTransaction) error
}

type IIncome interface {
	Create(context.Context,models.CreateIncome) (string, error)
	GetByID(context.Context,models.PrimaryKey) (models.Income, error)
	GetList(context.Context,models.GetListRequest) (models.IncomesResponse, error)
	Update(context.Context,models.UpdateIncome) (models.Income, error)
	Delete(context.Context,models.PrimaryKey) error
}

type IIncomeProduct interface {
	Create(context.Context,models.CreateIncomeProduct) (string, error)
	GetByID(context.Context,models.PrimaryKey) (models.IncomeProduct, error)
	GetList(context.Context,models.GetListRequest) (models.IncomeProductsResponse, error)
	Update(context.Context,models.UpdateIncomeProduct) (models.IncomeProduct, error)
	Delete(context.Context,models.PrimaryKey) error
}