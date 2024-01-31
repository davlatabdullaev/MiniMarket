package postgres

import (
	"database/sql"
	"developer/api/models"
	"developer/storage"
	"fmt"

	"github.com/google/uuid"
)

type basketRepo struct {
	DB *sql.DB
}

func NewBasketRepo (db *sql.DB) storage.IBasketStorage{
	return basketRepo{
		DB: db,
	}
}

func (b basketRepo) Create(basket models.CreateBasket)(string, error){
	uid := uuid.New()
	query := `INSERT INTO baskets values ($1, $2, $3, $4)`
	_, err := b.DB.Exec(query,
		uid,
		basket.SaleID,
		basket.ProductID,
		basket.Quantity,
		basket.Price,
	)
	if err != nil{
		fmt.Println("error while inserting to baskets!")
		return "", err
	}

	return uid.String(), nil
}

func (b basketRepo) GetByID(pKey models.PrimaryKey)(models.Basket, error){
	basket := models.Basket{}
	query := `SELECT id, sale_id, product_id, quantity, price,
			created_at, updated_at, deleted_at
	  from baskets where id = $1`
	err := b.DB.QueryRow(query,pKey.ID).Scan(
		&basket.ID,
		&basket.SaleID,
		&basket.ProductID,
		&basket.Quantity,
		&basket.Price,
		&basket.CreatedAt,
		&basket.UpdatedAt,
		&basket.DeletedAt,
	)

	if err != nil{
		fmt.Println("error while selecting basket by id!")
		return models.Basket{},err
	}
	return basket, nil
}

func (b basketRepo) GetList(request models.GetListRequest) (models.BasketsResponse, error){
	var (
		baskets = []models.Basket{}
		count = 0
		countQuery, query string
		page   = request.Page
		offset = (page - 1) * request.Limit
	)

	countQuery = `
	SELECT count(1) from baskets`

	err := b.DB.QueryRow(countQuery).Scan(&count)
	if err != nil{
		fmt.Println("error while scanning count of baskets!")
		return models.BasketsResponse{}, err
	}

	query = `SELECT id, sale_id, product_id, quantity, price,
		created_at, updated_at, deleted_at
	from baskets`


	query += `LIMIT $1 OFFSET $2`

	rows, err := b.DB.Query(query,request.Limit,offset)
	if err != nil{
		fmt.Println("error while query rows!", err.Error())
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
			&basket.DeletedAt,
		)
		if err != nil{
			fmt.Println("error while scanning rows!", err.Error())
			return models.BasketsResponse{}, err
		}
		baskets = append(baskets, basket)
	}
	return models.BasketsResponse{
		Baskets: baskets,
		Count: count,
	}, nil
}

func (b basketRepo) Update(basket models.UpdateBasket)(string, error){
	query :=  `UPDATE baskets set sale_id = $1, product_id = $2, quantity = $3, price = $4 where id = $5`
	uid, _ := uuid.Parse(basket.ID)
	_, err := b.DB.Exec(query,
		basket.SaleID,
		basket.ProductID,
		basket.Quantity,
		basket.Price,
		uid,
	)
	if err != nil{
		fmt.Println("error while updating baskets!", err.Error())
		return "", err
	}

	return basket.ID, nil
}

func (b basketRepo) Delete(pKey models.PrimaryKey) error{
	query := `DELETE from baskets where id = $1`
	_, err := b.DB.Exec(query,pKey)
	if err != nil{
		fmt.Println("error while deleting basket!", err.Error())
		return err
	}
	return nil
}