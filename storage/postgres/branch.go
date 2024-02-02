package postgres

import (
	"context"
	"developer/api/models"
	"developer/storage"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type branchRepo struct {
	DB *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) storage.IBranch{
	return &branchRepo{
		DB: db,
	}
}

func (b *branchRepo) Create(ctx context.Context, branch models.CreateBranch)(string, error){
	uid := uuid.New()
	query := `INSERT INTO branche (id, name, address) values ($1, $2, $3)`
	_, err := b.DB.Exec(ctx, query,
		uid,
		branch.Name,
		branch.Address,
	)
	if err != nil{
		fmt.Println("Error while inserting to branches!", err.Error())
		return "", err
	}

	return uid.String(), nil
}

func (b *branchRepo) GetByID(ctx context.Context,pKey models.PrimaryKey)(models.Branch, error){
	branch := models.Branch{}
	query := `SELECT id, name, address, created_at, 
	updated_at from branche where id = $1`
	err := b.DB.QueryRow(ctx,query,pKey.ID).Scan(
		&branch.ID,
		&branch.Name,
		&branch.Address,
		&branch.CreatedAt,
		&branch.UpdatedAt,
	)

	if err != nil{
		fmt.Println("Error while selecting branch by id!", err.Error())
		return models.Branch{},err
	}
	return branch, nil
}

func (b *branchRepo) GetList(ctx context.Context,request models.GetListRequest) (models.BranchesResponse, error){
	var (
		branches = []models.Branch{}
		count = 0
		countQuery, query string
		page   = request.Page
		offset = (page - 1) * request.Limit
		search = request.Search
	)

	countQuery = `
	SELECT count(1) from branche`

	if search != ""{
		countQuery += fmt.Sprintf(` WHERE (name ilike '%%%s%%' OR address ilike '%%%s%%')`, search, search)
	}

	err := b.DB.QueryRow(ctx,countQuery).Scan(&count)
	if err != nil{
		fmt.Println("Error while scanning count of branches!", err.Error())
		return models.BranchesResponse{}, err
	}

	query = `SELECT id, name, address, created_at, 
	updated_at from branche`
		
	
	if search != ""{
		query += fmt.Sprintf(`WHERE (name ilike '%%%s%%' OR address ilike '%%%s%%')`, search, search)
	}

	query += `LIMIT $1 OFFSET $2`

	rows, err := b.DB.Query(ctx,query,request.Limit,offset)
	if err != nil{
		fmt.Println("Error while selecting branches!", err.Error())
		return models.
		BranchesResponse{},err
	}

	for rows.Next(){
		branch := models.Branch{}

		 err := rows.Scan(
			&branch.ID,
			&branch.Name,
			&branch.Address,
			&branch.CreatedAt,
			&branch.UpdatedAt,
		)
		if err != nil{
			fmt.Println("Error while scanning branches!", err.Error())
			return models.BranchesResponse{}, err
		}
		branches = append(branches, branch)
	}
	return models.BranchesResponse{
		Branches: branches,
		Count: count,
	}, nil
}

func (b *branchRepo) Update(ctx context.Context,branch models.UpdateBranch)(string, error){
	query :=  `UPDATE branche set name = $1, address = $2 where id = $3`
	uid, _ := uuid.Parse(branch.ID)
	_, err := b.DB.Exec(ctx, query,
		branch.Name,
		branch.Address,
		uid,
	)
	if err != nil{
		fmt.Println("Error while updating branches!", err.Error())
		return "", err
	}

	return branch.ID, nil
}

func (b *branchRepo) Delete(ctx context.Context,pKey models.PrimaryKey) error{
	query := `DELETE from branche where id = $1`
	_, err := b.DB.Exec(ctx,query,pKey)
	if err != nil{
		fmt.Println("Error while deleting branch!", err.Error())
		return err
	}
	return nil
}