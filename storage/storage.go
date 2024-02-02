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
}

type IStorage interface {
	Create(context.Context,models.CreateStorage)(string, error)
	GetByID(context.Context,models.PrimaryKey) (models.Storage, error)
	GetList(context.Context,models.GetListRequest) (models.StoragesResponse, error)
	Update(context.Context,models.UpdateStorage) (string, error)
	Delete(context.Context,models.PrimaryKey) error
}