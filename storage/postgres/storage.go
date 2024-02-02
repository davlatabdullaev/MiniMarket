package postgres

import (
	"context"
	"developer/api/models"
	"developer/storage"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type storageRepo struct {
	DB *pgxpool.Pool
}

func NewStorageRepo (db *pgxpool.Pool) storage.IStorage{
	return &storageRepo{
		DB: db,
	}
}

func (s *storageRepo) Create(ctx context.Context,store models.CreateStorage)(string, error){
	uid := uuid.New()
	query := `INSERT INTO storage (id, product_id, branch_id, count) values ($1, $2, $3, $4)`
	_, err := s.DB.Exec(ctx,query,
		uid,
		store.ProductID,
		store.BranchID,
		store.Count,
	)
	if err != nil{
		fmt.Println("Error while inserting to storage!")
		return "", err
	}

	return uid.String(), nil
}

func (s *storageRepo) GetByID(ctx context.Context,pKey models.PrimaryKey)(models.Storage, error){
	store := models.Storage{}
	query :=  `SELECT id, product_id, branch_id, count,
			created_at, updated_at
		from storage where id = $1`
	err := s.DB.QueryRow(ctx,query,pKey.ID).Scan(
		&store.ID,
		&store.ProductID,
		&store.BranchID,
		&store.Count,
		&store.CreatedAt,
		&store.UpdatedAt,
	)

	if err != nil{
		fmt.Println("Error while selecting storage by id!")
		return models.Storage{},err
	}
	return store, nil
}

func (s *storageRepo) GetList(ctx context.Context,request models.GetListRequest) (models.StoragesResponse, error){
	var (
		storages = []models.Storage{}
		count = 0
		countQuery, query string
		page   = request.Page
		offset = (page - 1) * request.Limit
	)

	countQuery = `
	SELECT count(1) from storage`

	err := s.DB.QueryRow(ctx,countQuery).Scan(&count)
	if err != nil{
		fmt.Println("Error while scanning count of storages!")
		return models.StoragesResponse{}, err
	}

	query =  `SELECT id, product_id, branch_id, count,
		created_at, updated_at from storage`


	query += `LIMIT $1 OFFSET $2`

	rows, err := s.DB.Query(ctx,query,request.Limit,offset)
	if err != nil{
		fmt.Println("Error while selecting storages!", err.Error())
		return models.
		StoragesResponse{},err
	}

	for rows.Next(){
		store := models.Storage{}

	 err := rows.Scan(
		 &store.ID,
		 &store.ProductID,
		 &store.BranchID,
		 &store.Count,
		 &store.CreatedAt,
		 &store.UpdatedAt,
	)
		if err != nil{
			fmt.Println("Error while scanning storages!", err.Error())
			return models.StoragesResponse{}, err
		}
		storages = append(storages, store)
	}
	return models.StoragesResponse{
		Storages: storages,
		Count: count,
	}, nil
}

func (s *storageRepo) Update(ctx context.Context,updateStorage models.UpdateStorage)(string, error){
	query :=  `UPDATE storage SET product_id = $1, branch_id = $2, count = $3 where id = $4`
	uid, _ := uuid.Parse(updateStorage.ID)
	_, err := s.DB.Exec(ctx,query,
		 updateStorage.ProductID,
		 updateStorage.BranchID,
		 updateStorage.Count,
		 uid,
	)
	if err != nil{
		fmt.Println("Error while updating storages!", err.Error())
		return "", err
	}

	return updateStorage.ID, nil
}

func (s *storageRepo) Delete(ctx context.Context,pKey models.PrimaryKey) error{
	query := `DELETE from storage where id = $1`
	_, err := s.DB.Exec(ctx,query,pKey)
	if err != nil{
		fmt.Println("Error while deleting storage!", err.Error())
		return err
	}
	return nil
}