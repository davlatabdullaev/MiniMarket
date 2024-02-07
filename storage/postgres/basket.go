package postgres

import (
	"context"
	"developer/api/models"
	"developer/storage"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type basketRepo struct {
	DB *pgxpool.Pool
}

func NewBasketRepo(db *pgxpool.Pool) storage.IBasket {
	return &basketRepo{DB: db}
}

func (b *basketRepo) Create(ctx context.Context, basket models.CreateBasket)(string, error){
	uid := uuid.New()
	query := `INSERT INTO baskets (id, sale_id, product_id, quantity, price) values ($1, $2, $3, $4, $5)`
	_, err := b.DB.Exec(ctx,query,
		uid,
		basket.SaleID,
		basket.ProductID,
		basket.Quantity,
		basket.Price,
	)
	if err != nil{
		fmt.Println("Error while inserting to baskets!",err.Error())
		return "", err
	}

	return uid.String(), nil
}

func (b *basketRepo) GetByID(ctx context.Context, pKey models.PrimaryKey)(models.Basket, error){
	basket := models.Basket{}
	query := `SELECT id, sale_id, product_id, quantity, price,
			created_at, updated_at from baskets where id = $1`
	err := b.DB.QueryRow(ctx,query,pKey.ID).Scan(
		&basket.ID,
		&basket.SaleID,
		&basket.ProductID,
		&basket.Quantity,
		&basket.Price,
		&basket.CreatedAt,
		&basket.UpdatedAt,
	)

	if err != nil{
		fmt.Println("Error while selecting basket by id!", err.Error())
		return models.Basket{},err
	}
	return basket, nil
}

func (b *basketRepo) GetList(ctx context.Context, request models.GetListRequest) (models.BasketsResponse, error){
	var (
		baskets = []models.Basket{}
		count = 0
		countQuery, query string
		page   = request.Page
		offset = (page - 1) * request.Limit
	)

	countQuery = `
	SELECT count(1) from baskets`

	err := b.DB.QueryRow(ctx,countQuery).Scan(&count)
	if err != nil{
		fmt.Println("Error while scanning count of baskets!", err.Error())
		return models.BasketsResponse{}, err
	}

	query = `SELECT id, sale_id, product_id, quantity, price,
		created_at, updated_at from baskets `


	query += `LIMIT $1 OFFSET $2`

	rows, err := b.DB.Query(ctx,query,request.Limit,offset)
	if err != nil{
		fmt.Println("Error while selecting baskets!", err.Error())
		return models.
		BasketsResponse{},err
	}

	for rows.Next(){
		basket := models.Basket{}

		 err := rows.Scan(
			&basket.ID,
			&basket.SaleID,
			&basket.ProductID,
			&basket.Quantity,
			&basket.Price,
			&basket.CreatedAt,
			&basket.UpdatedAt,
		)
		if err != nil{
			fmt.Println("Error while scanning basket rows!", err.Error())
			return models.BasketsResponse{}, err
		}
		baskets = append(baskets, basket)
	}
	return models.BasketsResponse{
		Baskets: baskets,
		Count: count,
	}, nil
}

func (b *basketRepo) Update(ctx context.Context,basket models.UpdateBasket)(string, error){
	query :=  `UPDATE baskets set sale_id = $1, product_id = $2, quantity = $3, price = $4, updated_at = now() where id = $5`
	uid, _ := uuid.Parse(basket.ID)
	_, err := b.DB.Exec(ctx, query,
		basket.SaleID,
		basket.ProductID,
		basket.Quantity,
		basket.Price,
		uid,
	)
	if err != nil{
		fmt.Println("Error while updating baskets!", err.Error())
		return "", err
	}

	return basket.ID, nil
}

func (b *basketRepo) Delete(ctx context.Context, pKey models.PrimaryKey) error{
	query := `DELETE from baskets where id = $1`
	_, err := b.DB.Exec(ctx,query,pKey.ID)
	if err != nil{
		fmt.Println("Error while deleting basket!", err.Error())
		return err
	}
	return nil
}
 

