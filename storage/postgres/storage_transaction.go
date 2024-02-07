package postgres

import (
	"context"
	"developer/api/models"
	"developer/storage"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type storageTransactionRepo struct {
	DB *pgxpool.Pool
}

func NewStorageTransactionRepo(db *pgxpool.Pool) storage.IStorageTransaction {
	return &storageTransactionRepo{
		DB: db,
	}
}

func (s *storageTransactionRepo) Create(ctx context.Context, createStTrans models.CreateStorageTransaction) (string, error) {

	uid := uuid.New()

	query := `INSERT INTO storage_transactions (id, staff_id, product_id, storage_tranaction_type, 
		price, quantity) 
		values ($1, $2, $3, $4, $5, $6, $7)`

	_, err := s.DB.Exec(ctx, query,
		uid,
		createStTrans.StaffID,
		createStTrans.ProductID,
		createStTrans.StorageTransactionType,
		createStTrans.Price,
		createStTrans.Quantity,
	)
	if err != nil {
		fmt.Println("error while inserting storage transaction data", err.Error())
		return "", err
	}
	return uid.String(), nil
}

func (s *storageTransactionRepo) GetByID(ctx context.Context, id models.PrimaryKey) (models.StorageTransaction, error) {

	storageTransaction := models.StorageTransaction{}

	query := `SELECT id, staff_id, product_id, storage_transaction_type, 
	price, quantity, created_at, updated_at from storage_transaction
	 where id = $1`

	row := s.DB.QueryRow(ctx, query, id)

	err := row.Scan(
		&storageTransaction.ID,
		&storageTransaction.StaffID,
		&storageTransaction.ProductID,
		&storageTransaction.StorageTransactionType,
		&storageTransaction.Price,
		&storageTransaction.Quantity,
		&storageTransaction.CreatedAt,
		&storageTransaction.UpdatedAt,
	)

	if err != nil {
		fmt.Println("error while selecting storage transaction data", err.Error())
		return models.StorageTransaction{}, err
	}

	return storageTransaction, nil
}

func (s *storageTransactionRepo) GetList(ctx context.Context, request models.GetListRequest) (models.StorageTransactionsResponse, error) {

	var (
		storageTransactions = []models.StorageTransaction{}
		count               = 0
		query, countQuery   string
		page                = request.Page
		offset              = (page - 1) * request.Limit
		search              = request.Search
	)

	countQuery = `select count(1) from storage_transactions`

	if search != "" {
		countQuery += fmt.Sprintf(` where storage_transaction_type ilike '%%%s%%'`, search)
	}
	err := s.DB.QueryRow(ctx, countQuery).Scan(&count)
	if err != nil {
		fmt.Println("error is while selecting storage_transaction count", err.Error())
		return models.StorageTransactionsResponse{}, err
	}

	query = `select id, staff_id, product_id, storafe_transaction_type, 
	price, quantity, created_at, updated_at
	from storage_transactions`

	if search != "" {
		query += fmt.Sprintf(` where storage_transaction_type ilike '%%%s%%'`, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := s.DB.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting storage transaction", err.Error())
		return models.StorageTransactionsResponse{}, err
	}

	for rows.Next() {
		storageTransaction := models.StorageTransaction{}
		if err = rows.Scan(
			&storageTransaction.ID,
			&storageTransaction.StaffID,
			&storageTransaction.ProductID,
			&storageTransaction.StorageTransactionType,
			&storageTransaction.Price,
			&storageTransaction.Quantity,
			&storageTransaction.CreatedAt,
			&storageTransaction.UpdatedAt,
		); err != nil {
			fmt.Println("error is while scanning storage transaction data", err.Error())
			return models.StorageTransactionsResponse{}, err
		}

		storageTransactions = append(storageTransactions, storageTransaction)

	}

	return models.StorageTransactionsResponse{
		StorageTransactions: storageTransactions,
		Count:               count,
	}, nil
}

func (s *storageTransactionRepo) Update(ctx context.Context, updateStTrans models.UpdateStorageTransaction) (string, error) {

	query := `update storage_transaction
   set staff_id = $1, product_id = $2, storage_transaction_type = $3,
   price = $4, quantity = $5, updated_at = now()
   where id = $7
   `
	_, err := s.DB.Exec(ctx, query,
		updateStTrans.StaffID,
		updateStTrans.ProductID,
		updateStTrans.StorageTransactionType,
		updateStTrans.Price,
		updateStTrans.Quantity,
		updateStTrans.ID,
	)
	if err != nil {
		fmt.Println("error while updating storage_transaction data...", err.Error())
		return "", err
	}

	return updateStTrans.ID, nil
}

func (s *storageTransactionRepo) Delete(ctx context.Context, pKey models.PrimaryKey) error {

	query := `
	 DELETE FROM storage_transactions where id = $1
	`

	_, err := s.DB.Exec(ctx, query, pKey.ID)
	if err != nil {
		fmt.Println("error while deleting storage_transaction", err.Error())
		return err
	}

	return nil
}