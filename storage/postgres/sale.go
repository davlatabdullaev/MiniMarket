package postgres

import (
	"context"
	"developer/api/models"
	"developer/storage"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type saleRepo struct {
	DB *pgxpool.Pool
}

func NewSaleRepo(db *pgxpool.Pool) storage.ISale{
	return &saleRepo{
		DB: db,
	}
}

func (s *saleRepo) Create(ctx context.Context,sale models.CreateSale)(string, error){
	uid := uuid.New()
	query := `INSERT INTO sales (id, branch_id, shop_assistent_id, chashier_id,
			payment_type, price, status, client_name) 
	values ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := s.DB.Exec(ctx,query,
		uid,
		sale.BranchID,
		sale.ShopAssistantID,
		sale.CashierID,
		sale.PaymentType,
		sale.Price,
		sale.Status,
		sale.ClientName,
	)
	if err != nil{
		fmt.Println("Error while inserting to sales!", err.Error())
		return "", err
	}

	return uid.String(), nil
}

func (s *saleRepo) GetByID(ctx context.Context,pKey models.PrimaryKey)(models.Sale, error){
	sale := models.Sale{}
	query :=  `SELECT id, branch_id, shop_assistent_id,
			chashier_id, payment_type, price, status, client_name,
			created_at, updated_at from sales where id = $1`
	err := s.DB.QueryRow(ctx,query,pKey.ID).Scan(
		 &sale.ID,
		 &sale.BranchID,
		 &sale.ShopAssistantID,
		 &sale.CashierID,
		 &sale.PaymentType,
		 &sale.Price,
		 &sale.Status,
		 &sale.ClientName,
		 &sale.CreatedAt,
		 &sale.UpdatedAt,
	)

	if err != nil{
		fmt.Println("Error while selecting sale by id!", err.Error())
		return models.Sale{},err
	}
	return sale, nil
}

func (s *saleRepo) GetList(ctx context.Context,request models.GetListRequest) (models.SalesResponse, error){
	var (
		sales = []models.Sale{}
		count = 0
		countQuery, query string
		page   = request.Page
		offset = (page - 1) * request.Limit
		search = request.Search
	)

	countQuery = `
	SELECT count(1) from sales`

	if search != ""{
		countQuery += fmt.Sprintf(` WHERE (client_name ilike '%%%s%%' OR payment_type ilike '%%%s%%')`, search, search)
	}

	err := s.DB.QueryRow(ctx,countQuery).Scan(&count)
	if err != nil{
		fmt.Println("Error while scanning count of sales!")
		return models.SalesResponse{}, err
	}

	query = `SELECT id, branch_id, shop_assistent_id,
	chashier_id, payment_type, price, status, client_name,
	created_at, updated_at from sales`
		
	
	if search != ""{
		query += fmt.Sprintf(` WHERE (clent_name ilike '%%%s%%' OR payment_type ilike '%%%s%%)`,search, search)
	}

	query += ` LIMIT $1 OFFSET $2`

	rows, err := s.DB.Query(ctx,query,request.Limit,offset)
	if err != nil{
		fmt.Println("Error while selecting sales!", err.Error())
		return models.SalesResponse{},err
	}

	for rows.Next(){
		sale := models.Sale{}

		 err := rows.Scan(
			&sale.ID,
			&sale.BranchID,
			&sale.ShopAssistantID,
			&sale.CashierID,
			&sale.PaymentType,
			&sale.Price,
			&sale.Status,
			&sale.ClientName,
			&sale.CreatedAt,
			&sale.UpdatedAt,
		)
		if err != nil{
			fmt.Println("Error while scanning sales!", err.Error())
			return models.SalesResponse{}, err
		}
		 sales = append(sales, sale)
	}
	return models.SalesResponse{
		
		Sales: sales,
		Count: count,
	}, nil
}

func (s *saleRepo) Update(ctx context.Context,sale models.UpdateSale)(string, error){
	query :=  `UPDATE sales set branch_id = $1, shop_assistent_id = $2,
			chashier_id = $3, payment_type = $4, price = $5,
			status = $6, client_name = $7, updated_at = now()
			WHERE id = $8`
	_, err := s.DB.Exec(ctx,query,
		sale.BranchID,
		sale.ShopAssistantID,
		sale.CashierID,
		sale.PaymentType,
		sale.Price,
		sale.Status,
		sale.ClientName,
		sale.ID,
	)
	if err != nil{
		fmt.Println("Error while updating sales!", err.Error())
		return "", err
	}

	return sale.ID, nil
}

func (s *saleRepo) Delete(ctx context.Context, pKey models.PrimaryKey) error{
	query := `DELETE from sales where id = $1`
	_, err := s.DB.Exec(ctx,query,pKey.ID)
	if err != nil{
		fmt.Println("Error while deleting sale!", err.Error())
		return err
	}
	return nil
}