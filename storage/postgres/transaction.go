package postgres

import (
	"context"
	"developer/api/models"
	"developer/storage"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type transactionRepo struct {
	DB *pgxpool.Pool
}

func NewTransactionRepo(db *pgxpool.Pool) storage.ITransaction {
	return &transactionRepo{
		DB: db,
	}
}

func (t *transactionRepo) Create(ctx context.Context, request models.CreateTransaction) (string, error) {

	uid := uuid.New()

	query := `INSERT INTO transactions (id, sale_id, staff_id, transaction_type,
		source_type, amount, description) 
	values 
	($1, $2, $3, $4, $5, $6, $7)`

	_, err := t.DB.Exec(ctx, query,
		uid,
		request.SaleID,
		request.StaffID,
		request.TransactionType,
		request.SourceType,
		request.Amount,
		request.Description,
	)
	if err != nil {
		fmt.Println("error while inserting transaction data", err.Error())
		return "", err
	}

	return uid.String(), nil
}

func (t *transactionRepo) GetByID(ctx context.Context, pKey models.PrimaryKey) (models.Transaction, error) {

	transaction := models.Transaction{}

	query := `select id, sale_id, staff_id, transaction_type,
	source_type, amount, description, 
	 created_at, updated_at from tarifs
	 where id = $1`

	row := t.DB.QueryRow(ctx, query, pKey)

	err := row.Scan(
		&transaction.ID,
		&transaction.SaleID,
		&transaction.StaffID,
		&transaction.TransactionType,
		&transaction.SourceType,
		&transaction.Amount,
		&transaction.Description,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)

	if err != nil {
		fmt.Println("error while selecting transaction data", err.Error())
		return models.Transaction{}, err
	}

	return transaction, nil
}

func (t *transactionRepo) GetList(ctx context.Context, request models.GetListRequest) (models.TransactionsResponse, error) {

	var (
		transactions      = []models.Transaction{}
		count             = 0
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `select count(1) from transactions`

	if search != "" {
		countQuery += fmt.Sprintf(` where description ilike '%%%s%%'`, search)
	}
	if err := t.DB.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting transaction count", err.Error())
		return models.TransactionsResponse{}, err
	}

	query = `select id, sale_id, staff_id, 
	transaction_type, source_type,
	amount, description,
	created_at, updated_at
	from tarifs`

	if search != "" {
		query += fmt.Sprintf(` where description ilike '%%%s%%'`, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := t.DB.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting transaction", err.Error())
		return models.TransactionsResponse{}, err
	}

	for rows.Next() {
		transaction := models.Transaction{}
		err := rows.Scan(
			&transaction.ID,
			&transaction.SaleID,
			&transaction.StaffID,
			&transaction.TransactionType,
			&transaction.SourceType,
			&transaction.Amount,
			&transaction.Description,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		)
		if err != nil {
			fmt.Println("error is while scanning transaction data", err.Error())
			return models.TransactionsResponse{}, err
		}

		transactions = append(transactions, transaction)

	}

	return models.TransactionsResponse{
		Transactions: transactions,
		Count:        count,
	}, nil
}

func (t *transactionRepo) Update(ctx context.Context, request models.UpdateTransaction) (string, error) {

	query := `update transaction
   set sale_id = $1, staff_id = $2, transaction_type = $3,
   source_type = $4, amount = $5, description = $6, updated_at = now()
   where id = $8
   `
	_, err := t.DB.Exec(ctx, query,
		request.SaleID,
		request.StaffID,
		request.TransactionType,
		request.SourceType,
		request.Amount,
		request.Description,
		request.ID,
	)
	if err != nil {
		fmt.Println("error while updating transaction data...", err.Error())
		return "", err
	}
	return request.ID, nil
}

func (t *transactionRepo) Delete(ctx context.Context, pKey models.PrimaryKey) error {

	query := `
	 DELETE FROM transactions where id = $1
	`

	_, err := t.DB.Exec(ctx, query, pKey.ID)
	if err != nil {
		fmt.Println("error while deleting transaction!", err.Error())
		return err
	}

	return nil
}

// NEW from assistant
func (t *transactionRepo) UpdateStaffBalanceAndCreateTransaction(ctx context.Context, req models.UpdateStaffBalanceAndCreateTransaction) error {
	tr, err := t.DB.Begin(ctx)
	if err != nil {
		fmt.Println("error while starting transaction to db!", err.Error())
		return err
	}

	defer func() {
		if err != nil {
			tr.Rollback(ctx)
		}
		tr.Commit(ctx)
	}()

	cashierUpdateQuery := `
	UPDATE staffs set balance = balance + $1 where id = $2`

	_, err = tr.Exec(ctx, cashierUpdateQuery,
		req.Cashier.Amount,
		req.Cashier.StaffID,
	)

	if err != nil {
		fmt.Println("error while updating balance of cashier!", err.Error())
		return err
	}

	uid := uuid.New()

	query := `INSERT INTO transactions (id, sale_id, staff_id, transaction_type,
		source_type, amount, description) 
	values 
	($1, $2, $3, $4, $5, $6, $7) `

	_, err = tr.Exec(ctx, query,
		uid,
		req.SaleID,
		req.Cashier.StaffID,
		req.TransactionType,
		req.SourceType,
		req.Amount,
		req.Description,
	)
	if err != nil {
		fmt.Println("error while creating a transaction!", err.Error())
		return err
	}

	if req.ShopAssistant.StaffID != "" {
		shopAssistantUpdateQuery := `
	UPDATE staffs set balance = balance + $1 where id = $2`

	_, err = tr.Exec(ctx, shopAssistantUpdateQuery,
		req.Cashier.Amount,
		req.Cashier.StaffID,
	)

	if err != nil {
		fmt.Println("error while updating balance of cashier!", err.Error())
		return err
	}

	uid = uuid.New()

	query2 := `INSERT INTO transactions (id, sale_id, staff_id, transaction_type,
		source_type, amount, description) 
	values 
	($1, $2, $3, $4, $5, $6, $7) `

	_, err = t.DB.Exec(ctx, query2,
		uid,
		req.SaleID,
		req.ShopAssistant.StaffID,
		req.TransactionType,
		req.SourceType,
		req.Amount,
		req.Description,
	)
	if err != nil {
		fmt.Println("error while creating a transaction!", err.Error())
		return err
	}
	}
	return nil
}
