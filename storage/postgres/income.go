package postgres

import (
	"context"
	"developer/api/models"
	"developer/storage"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type incomeRepo struct {
	DB *pgxpool.Pool
}

func NewIncomeRepo(db *pgxpool.Pool)storage.IIncome{
	return &incomeRepo{
		DB: db,
	}
}

func (i *incomeRepo) Create(ctx context.Context, createIncome models.CreateIncome)(string, error){
	query := `INSERT INTO incomes (id, branch_id, price) 
		values ($1, $2, $3) `
	uid := uuid.New()
	_, err := i.DB.Exec(ctx, query,
		uid,
		createIncome.BranchId,
		createIncome.Price,
	)
	if err != nil{
		fmt.Println("Error while inserting into incomes!", err.Error())
		return "",err
	}

	return uid.String(), nil
}

func (i *incomeRepo) GetByID(ctx context.Context, pKey models.PrimaryKey)(models.Income, error){
	query := `SELECT id, branch_id, price, created_at, updated_at 
	from incomes where id = $1 and deleted_at = 0 `

	income := models.Income{}

	err := i.DB.QueryRow(ctx, query,pKey.ID).Scan(
		&income.ID,
		&income.BranchId,
		&income.Price,
		&income.CreatedAt,
		&income.UpdatedAt,
	)
	if err != nil{
		fmt.Println("Error while getting income by id!", err.Error())
		return models.Income{}, err
	}

	return income, nil
}

func (i *incomeRepo) GetList(ctx context.Context, req models.GetListRequest) (models.IncomesResponse, error) {
	countQuery := `SELECT (1) count from incomes where deleted_at = 0`
	tempCount := 0
	offset := (req.Page - 1) * req.Limit
	incomes := []models.Income{}

	err := i.DB.QueryRow(ctx,countQuery).Scan(&tempCount)
	if err != nil{
		fmt.Println("Error while scanning count of incomes!", err.Error())
		return models.IncomesResponse{},err
	}

	query := `SELECT id, branch_id, price, created_at, updated_at
		from incomes where deleted_at = 0 `
	
	query += ` LIMIT $1 OFFSET $2`

	rows, err := i.DB.Query(ctx, query, req.Limit, offset)
	for rows.Next(){
		income := models.Income{}
			rows.Scan(
			&income.ID,
			&income.BranchId,
			&income.Price,
			&income.CreatedAt,
			&income.UpdatedAt,
		)
		if err != nil{
			fmt.Println("Error while scanning into incomes!",err.Error())
			return models.IncomesResponse{},err
		}
		incomes = append(incomes, income)
	}
	return models.IncomesResponse{
		Incomes: incomes,
		Count: tempCount,
	},nil
}

func (i *incomeRepo) Update(ctx context.Context, updIncome models.UpdateIncome) (models.Income, error) {
	income := models.Income{}

	query := `UPDATE incomes set branch_id = $1, price = $2, updated_at = now()
		where id = $3 and deleted_at = 0
			returning id, branch_id, price, updated_at `
	err := i.DB.QueryRow(ctx,query,
		updIncome.BranchId,
		updIncome.Price,
		updIncome.ID,
	).Scan(
		&income.ID,
		&income.BranchId,
		&income.Price,
		&income.UpdatedAt,
	)
	if err != nil{
		fmt.Println("Error while updating incomes!",err.Error())
		return models.Income{},err
	}

	return income, nil
}

func (i *incomeRepo) Delete(ctx context.Context, pKey models.PrimaryKey) (error) {
	query := `UPDATE incomes set deleted_at = 1 where id = $1 and deleted_at = 0`
	_, err := i.DB.Exec(ctx, query, pKey.ID)
	if err != nil{
		fmt.Println("Error while deleting income!",err.Error())
		return err
	}
	return nil
}