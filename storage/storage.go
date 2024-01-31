package storage

import "developer/api/models"



type IStorage interface {
	Close()
	Branch() IBranchStorage
	Sale()   ISaleStorage
	Basket() IBasketStorage
	Product() IProductStorage
	Repository() IRepositoryStorage
}

type IBranchStorage interface {
	Create(models.CreateBranch)(string, error)
	GetByID(models.PrimaryKey) (models.Branch, error)
	GetList(models.GetListRequest) (models.BranchesResponse, error)
	Update(models.UpdateBranch) (string, error)
	Delete(models.PrimaryKey) error
}

type ISaleStorage interface {
	Create(models.CreateSale)(string, error)
	GetByID(models.PrimaryKey) (models.Sale, error)
	GetList(models.GetListRequest) (models.SalesResponse, error)
	Update(models.UpdateSale) (string, error)
	Delete(models.PrimaryKey) error
}

type IBasketStorage interface {
	Create(models.CreateBasket)(string, error)
	GetByID(models.PrimaryKey) (models.Basket, error)
	GetList(models.GetListRequest) (models.BasketsResponse, error)
	Update(models.UpdateBasket) (string, error)
	Delete(models.PrimaryKey) error
}
type IProductStorage interface {
	Create(models.CreateProduct)(string, error)
	GetByID(models.PrimaryKey) (models.Product, error)
	GetList(models.GetListRequest) (models.ProductsResponse, error)
	Update(models.UpdateProduct) (string, error)
	Delete(models.PrimaryKey) error
}

type IRepositoryStorage interface {
	Create(models.CreateRepository)(string, error)
	GetByID(models.PrimaryKey) (models.Repository, error)
	GetList(models.GetListRequest) (models.RepositoriesResponse, error)
	Update(models.UpdateRepository) (string, error)
	Delete(models.PrimaryKey) error
}