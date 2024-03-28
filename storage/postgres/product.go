package postgres

import (
	"context"
	"developer/api/models"
	"developer/storage"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type productRepo struct {
	DB *pgxpool.Pool
}

func NewProductRepo (db *pgxpool.Pool) storage.IProduct{
	return &productRepo{
		DB: db,
	}
}

func (p *productRepo) Create(ctx context.Context,product models.CreateProduct)(string, error){
	uid := uuid.New()
	query := `INSERT INTO products (id, name, price, bar_code, category_id) values ($1, $2, $3, $4, $5)`
	_, err := p.DB.Exec(ctx, query,
		uid,
		product.Name,
		product.Price,
		product.BarCode,
		product.CategoryID,
	)
	if err != nil{
		fmt.Println("Error while inserting to products!", err.Error())
		return "", err
	}

	return uid.String(), nil
}

func (p *productRepo) GetByID(ctx context.Context, pKey models.PrimaryKey)(models.Product, error){
	product := models.Product{}
	query :=  `SELECT id, name, price, barcode, category_id,
			created_at, updated_at from products where id = $1`
	err := p.DB.QueryRow(ctx,query,pKey.ID).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.BarCode,
		&product.CategoryID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil{
		fmt.Println("Error while selecting basket by id!", err.Error())
		return models.Product{},err
	}
	return product, nil
}

func (p *productRepo) GetList(ctx context.Context, request models.GetListRequest) (models.ProductsResponse, error){
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

	err := p.DB.QueryRow(ctx,countQuery).Scan(&count)
	if err != nil{
		fmt.Println("Error while scanning count of products!", err.Error())
		return models.ProductsResponse{}, err
	}

	query =  `SELECT id, name, price, barcode, category_id,
			created_at, updated_at from products `

	
	if search != ""{
		query += fmt.Sprintf(`WHERE (name ilike '%%%s%%')`, search)
	}

	query += ` LIMIT $1 OFFSET $2`

	rows, err := p.DB.Query(ctx,query,request.Limit,offset)
	if err != nil{
		fmt.Println("Error while selecting products!", err.Error())
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
		)
		if err != nil{
			fmt.Println("Error while scanning products!", err.Error())
			return models.ProductsResponse{}, err
		}
		products = append(products, product)
	}
	return models.ProductsResponse{
		Products: products,
		Count: count,
	}, nil
}

func (p *productRepo) Update(ctx context.Context, product models.UpdateProduct)(string, error){
	query :=  `UPDATE products set name = $1, price = $2, category_id = $3, updated_at = now() where id = $4`
	uid, _ := uuid.Parse(product.ID)
	_, err := p.DB.Exec(ctx,query,
		product.Name,
		product.Price,
		product.CategoryID,
		uid,	
	)
	if err != nil{
		fmt.Println("Error while updating products!", err.Error())
		return "", err
	}

	return product.ID, nil
}

func (p *productRepo) Delete(ctx context.Context,pKey models.PrimaryKey) error{
	query := `DELETE from products where id = $1`
	_, err := p.DB.Exec(ctx,query,pKey.ID)
	if err != nil{
		fmt.Println("Error while deleting product!", err.Error())
		return err
	}
	return nil
}

func (p *productRepo) GetByBarcode(ctx context.Context, barcode string)(models.Product, error){
	product := models.Product{}

	query := `SELECT id, name, price, barcode, category_id, created_at, updated_at
	from products where barcode = $1
	`

	err := p.DB.QueryRow(ctx,query,barcode).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.BarCode,
		&product.CategoryID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil{
		fmt.Println("error while selectin product by barcode!", err.Error())
		return models.Product{}, err
	}

	return product, nil
}