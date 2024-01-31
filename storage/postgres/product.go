package postgres

import (
	"database/sql"
	"developer/api/models"
	"developer/storage"
	"fmt"

	"github.com/google/uuid"
)

type productRepo struct {
	DB *sql.DB
}

func NewProductRepo (db *sql.DB) storage.IProductStorage{
	return productRepo{
		DB: db,
	}
}

func (p productRepo) Create(product models.CreateProduct)(string, error){
	uid := uuid.New()
	query := `INSERT INTO products values ($1, $2, $3, $4, $5)`
	_, err := p.DB.Exec(query,
		uid,
		product.Name,
		product.Price,
		product.BarCode,
		product.CategoryID,
	)
	if err != nil{
		fmt.Println("error while inserting to products!")
		return "", err
	}

	return uid.String(), nil
}

func (p productRepo) GetByID(pKey models.PrimaryKey)(models.Product, error){
	product := models.Product{}
	query :=  `SELECT id, name, price, bar_code, category_id,
			created_at, updated_at, deleted_at 
		from products where id = $1`
	err := p.DB.QueryRow(query,pKey.ID).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.BarCode,
		&product.CategoryID,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.DeletedAt,
	)

	if err != nil{
		fmt.Println("error while selecting basket by id!")
		return models.Product{},err
	}
	return product, nil
}

func (p productRepo) GetList(request models.GetListRequest) (models.ProductsResponse, error){
	var (
		products = []models.Product{}
		count = 0
		countQuery, query string
		page   = request.Page
		offset = (page - 1) * request.Limit
		search = request.Search
	)

	countQuery = `
	SELECT count(1) from products`

	if search != ""{
		countQuery += fmt.Sprintf(` WHERE (name ilike '%%%s%%')`, search)
	}

	err := p.DB.QueryRow(countQuery).Scan(&count)
	if err != nil{
		fmt.Println("error while scanning count of products!")
		return models.ProductsResponse{}, err
	}

	query =  `SELECT id, name, price, bar_code, category_id,
			created_at, updated_at, deleted_at 
		from products where id = $1`

	
	if search != ""{
		query += fmt.Sprintf(`WHERE (name ilike '%%%s%%')`, search)
	}

	query += `LIMIT $1 OFFSET $2`

	rows, err := p.DB.Query(query,request.Limit,offset)
	if err != nil{
		fmt.Println("error while query rows!", err.Error())
		return models.
		ProductsResponse{},err
	}

	for rows.Next(){
		product := models.Product{}

	 err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.BarCode,
		&product.CategoryID,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.DeletedAt,
		)
		if err != nil{
			fmt.Println("error while scanning rows!", err.Error())
			return models.ProductsResponse{}, err
		}
		products = append(products, product)
	}
	return models.ProductsResponse{
		Products: products,
		Count: count,
	}, nil
}

func (p productRepo) Update(product models.UpdateProduct)(string, error){
	query :=  `UPDATE products set name = $1, price = $2, category_id = $3 where id = $4`
	uid, _ := uuid.Parse(product.ID)
	_, err := p.DB.Exec(query,
		product.Name,
		product.Price,
		product.CategoryID,
		uid,	
	)
	if err != nil{
		fmt.Println("error while updating products!", err.Error())
		return "", err
	}

	return product.ID, nil
}

func (b productRepo) Delete(pKey models.PrimaryKey) error{
	query := `DELETE from products where id = $1`
	_, err := b.DB.Exec(query,pKey)
	if err != nil{
		fmt.Println("error while deleting product!", err.Error())
		return err
	}
	return nil
}