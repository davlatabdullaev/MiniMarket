package postgres

import (
	"database/sql"
	"developer/api/models"
	"developer/storage"

	"fmt"

	"github.com/google/uuid"
)

type branchRepo struct {
	DB *sql.DB
}

func NewBranchRepo(db *sql.DB) storage.IBranchStorage{
	return branchRepo{
		DB: db,
	}
}

func (b branchRepo) Create(branch models.CreateBranch)(string, error){
	uid := uuid.New()
	query := `INSERT INTO branches values ($1, $2, $3)`
	_, err := b.DB.Exec(query,
		uid,
		branch.Name,
		branch.Address,
	)
	if err != nil{
		fmt.Println("error while inserting to branches!")
		return "", err
	}

	return uid.String(), nil
}

func (b branchRepo) GetByID(pKey models.PrimaryKey)(models.Branch, error){
	branch := models.Branch{}
	query := `SELECT id, name, address, created_at, 
	updated_at, deleted_at from branches where id = $1`
	err := b.DB.QueryRow(query,pKey.ID).Scan(
		&branch.ID,
		&branch.Name,
		&branch.Address,
		&branch.CreatedAt,
		&branch.UpdatedAt,
		&branch.DeletedAt,
	)

	if err != nil{
		fmt.Println("error while selecting branch by id!")
		return models.Branch{},err
	}
	return branch, nil
}

func (b branchRepo) GetList(request models.GetListRequest) (models.BranchesResponse, error){
	var (
		branches = []models.Branch{}
		count = 0
		countQuery, query string
		page   = request.Page
		offset = (page - 1) * request.Limit
		search = request.Search
	)

	countQuery = `
	SELECT count(1) from branches`

	if search != ""{
		countQuery += fmt.Sprintf(` WHERE (name ilike '%%%s%%' OR address ilike '%%%s%%')`, search, search)
	}

	err := b.DB.QueryRow(countQuery).Scan(&count)
	if err != nil{
		fmt.Println("error while scanning count of branches!")
		return models.BranchesResponse{}, err
	}

	query = `SELECT id, name, address, created_at, 
	updated_at, deleted_at from branches`
		
	
	if search != ""{
		query += fmt.Sprintf(`WHERE (name ilike '%%%s%%' OR address ilike '%%%s%%')`, search, search)
	}

	query += `LIMIT $1 OFFSET $2`

	rows, err := b.DB.Query(query,request.Limit,offset)
	if err != nil{
		fmt.Println("error while query rows!", err.Error())
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
			&branch.DeletedAt,
		)
		if err != nil{
			fmt.Println("error while scanning rows!", err.Error())
			return models.BranchesResponse{}, err
		}
		branches = append(branches, branch)
	}
	return models.BranchesResponse{
		Branches: branches,
		Count: count,
	}, nil
}

func (b branchRepo) Update(branch models.UpdateBranch)(string, error){
	query :=  `UPDATE branches set name = $1, address = $2 where id = $3`
	uid, _ := uuid.Parse(branch.ID)
	_, err := b.DB.Exec(query,
		branch.Name,
		branch.Address,
		uid,
	)
	if err != nil{
		fmt.Println("error while updating branches!", err.Error())
		return "", err
	}

	return branch.ID, nil
}

func (b branchRepo) Delete(pKey models.PrimaryKey) error{
	query := `DELETE from branches where id = $1`
	_, err := b.DB.Exec(query,pKey)
	if err != nil{
		fmt.Println("error while deleting branch!", err.Error())
		return err
	}
	return nil
}