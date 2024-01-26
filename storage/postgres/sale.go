package postgres

import (
	"database/sql"
	"developer/api/models"
	"developer/storage"
	"time"

	"fmt"

	"github.com/google/uuid"
)

type saleRepo struct {
	DB *sql.DB
}

func NewSaleRepo(db *sql.DB) storage.ISaleStorage{
	return &saleRepo{
		DB: db,
	}
}

func (s saleRepo) Create(sale models.CreateSale)(string, error){
	uid := uuid.New()
	query := `INSERT INTO sales values ($1,$2,$3,$4,45,$6,$7,$8)`
	_, err := s.DB.Exec(query,
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
		fmt.Println("error while inserting to sales!")
		return "", err
	}

	return uid.String(), nil
}

func (s saleRepo) GetByID(pKey models.PrimaryKey)(models.Sale, error){
	sale := models.Sale{}
	query :=  `SELECT id, branch_id, shop_assistant_id,
			cashier_id, payment_type, price, status, client_name,
			created_at, updated_at, deleted_at from sales where id = $1`
	err := s.DB.QueryRow(query,pKey.ID).Scan(
		 &sale.ID,
		 &sale.BranchID,
		 &sale.ShopAssistantID,
		 &sale.CashierID,
		 &sale.PaymentType,
		 &sale.Price,
		 &sale.Status,
		 &sale.ClientName,
		 &sale.CreatedAt,
		 &sale.DeletedAt,
	)

	if err != nil{
		fmt.Println("error while selecting sale by id!")
		return models.Sale{},err
	}
	return sale, nil
}

func (s saleRepo) GetList(request models.GetListRequest) (models.SalesResponse, error){
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
		countQuery += fmt.Sprintf(` WHERE (client_name ilike '%%%s%%')`, search)
	}

	err := s.DB.QueryRow(countQuery).Scan(&count)
	if err != nil{
		fmt.Println("error while scanning count of sales!")
		return models.SalesResponse{}, err
	}

	query = `SELECT id, branch_id, shop_assistant_id,
	cashier_id, payment_type, price, status, client_name,
	created_at, updated_at, deleted_at from sales where id = $1`
		
	
	if search != ""{
		query += fmt.Sprintf(`WHERE (clent_name ilike '%%%s%%')`,search)
	}

	query += `LIMIT $1 OFFSET $2`

	rows, err := s.DB.Query(query,request.Limit,offset)
	if err != nil{
		fmt.Println("error while query rows!", err.Error())
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
			&sale.DeletedAt,
		)
		if err != nil{
			fmt.Println("error while scanning rows!", err.Error())
			return models.SalesResponse{}, err
		}
		 sales = append(sales, sale)
	}
	return models.SalesResponse{
		
		Sales: sales,
		Count: count,
	}, nil
}

func (s saleRepo) Update(sale models.UpdateSale)(string, error){
	sale.UpdatedAt = time.Now()
	query :=  `UPDATE sales set branch_id = $1, shop_assistant_id = $2
			cashier_id = $3, payment_type = $4, price = $5,
			status = $6, client_name = $7, updated_at = $8
			WHERE id = $9`
	_, err := s.DB.Exec(query,
		sale.BranchID,
		sale.ShopAssistantID,
		sale.CashierID,
		sale.PaymentType,
		sale.Price,
		sale.Status,
		sale.ClientName,
		sale.UpdatedAt,
		sale.ID,
	)
	if err != nil{
		fmt.Println("error while updating sales!", err.Error())
		return "", err
	}

	return sale.ID, nil
}

func (s saleRepo) Delete(pKey models.PrimaryKey) error{
	query := `DELETE from sales where id = $1`
	_, err := s.DB.Exec(query,pKey)
	if err != nil{
		fmt.Println("error while deleting sale!", err.Error())
		return err
	}
	return nil
}