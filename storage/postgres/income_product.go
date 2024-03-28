package postgres

import (
	"context"
	"developer/api/models"
	"developer/storage"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type incomeProductRepo struct {
	DB *pgxpool.Pool
}

func NewIncomeProductRepo(db *pgxpool.Pool)storage.IIncomeProduct{
	return &incomeProductRepo{
		DB: db,
	}
}

func (i *incomeProductRepo) Create(ctx context.Context, createIncomeProduct models.CreateIncomeProduct)(string, error){
	query := `INSERT INTO income_products (id, income_id, product_id, price) 
		values ($1, $2, $3, $4) `
	uid := uuid.New()
	_, err := i.DB.Exec(ctx, query,
		uid,
		createIncomeProduct.IncomeID,
		createIncomeProduct.ProductID,
		createIncomeProduct.Price,
	)
	if err != nil{
		fmt.Println("Error while inserting into incomes!", err.Error())
		return "",err
	}

	return uid.String(), nil
}

func (i *incomeProductRepo) GetByID(ctx context.Context, pKey models.PrimaryKey)(models.IncomeProduct, error){
	query := `SELECT id, income_id, product_id, price, created_at, updated_at 
	from income_products where id = $1 and deleted_at = 0 `

	incomeProduct := models.IncomeProduct{}

	err := i.DB.QueryRow(ctx, query,pKey.ID).Scan(
		&incomeProduct.ID,
		&incomeProduct.IncomeID,
		&incomeProduct.ProductID,
		&incomeProduct.Price,
		&incomeProduct.CreatedAt,
		&incomeProduct.UpdatedAt,
	)
	if err != nil{
		fmt.Println("Error while getting income by id!", err.Error())
		return models.IncomeProduct{}, err
	}

	return incomeProduct, nil
}

func (i *incomeProductRepo) GetList(ctx context.Context, req models.GetListRequest) (models.IncomeProductsResponse, error) {
	countQuery := `SELECT (1) count from income_products where deleted_at = 0 `
	tempCount := 0
	offset := (req.Page - 1) * req.Limit
	incomeProducts := []models.IncomeProduct{}

	err := i.DB.QueryRow(ctx,countQuery).Scan(&tempCount)
	if err != nil{
		fmt.Println("Error while scanning count of incomes!", err.Error())
		return models.IncomeProductsResponse{},err
	}

	query := `SELECT id, income_id, product_id, price, created_at, updated_at
		from income_products where deleted_at = 0 `
	
	query += ` LIMIT $1 OFFSET $2`

	rows, err := i.DB.Query(ctx, query, req.Limit, offset)
	for rows.Next(){
		incomeProduct := models.IncomeProduct{}
			rows.Scan(
				&incomeProduct.ID,
				&incomeProduct.IncomeID,
				&incomeProduct.ProductID,
				&incomeProduct.Price,
				&incomeProduct.CreatedAt,
				&incomeProduct.UpdatedAt,
		)
		if err != nil{
			fmt.Println("Error while scanning into incomes!",err.Error())
			return models.IncomeProductsResponse{},err
		}
		incomeProducts = append(incomeProducts, incomeProduct)
	}
	return models.IncomeProductsResponse{
		IncomeProducts: incomeProducts,
		Count: tempCount,
	},nil
}

func (i *incomeProductRepo) Update(ctx context.Context, updIncomeProduct models.UpdateIncomeProduct) (models.IncomeProduct, error) {
	incomeProduct := models.IncomeProduct{}

	query := `UPDATE income_products set income_id = $1, product_id = $2, price = $3, updated_at = now()
		where id = $4 and deleted_at = 0 
			returning id, income_id, product_id, price, updated_at `
	err := i.DB.QueryRow(ctx,query,
		updIncomeProduct.IncomeID,
		updIncomeProduct.ProductID,
		updIncomeProduct.Price,
		updIncomeProduct.ID,
	).Scan(
		&incomeProduct.ID,
		&incomeProduct.IncomeID,
		&incomeProduct.ProductID,
		&incomeProduct.Price,
		&incomeProduct.UpdatedAt,
	)
	if err != nil{
		fmt.Println("Error while updating incomes!",err.Error())
		return models.IncomeProduct{},err
	}

	return incomeProduct, nil
}

func (i *incomeProductRepo) Delete(ctx context.Context, pKey models.PrimaryKey) (error) {
	query := `UPDATE income_products set deleted_at = 1 where id = $1 and deleted_at = 0`
	_, err := i.DB.Exec(ctx, query, pKey.ID)
	if err != nil{
		fmt.Println("Error while deleting income!",err.Error())
		return err
	}
	return nil
}