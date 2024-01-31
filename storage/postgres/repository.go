package postgres

import (
	"database/sql"
	"developer/api/models"
	"developer/storage"
	"fmt"

	"github.com/google/uuid"
)

type repositoryRepo struct {
	DB *sql.DB
}

func NewRepositoryRepo (db *sql.DB) storage.IRepositoryStorage{
	return repositoryRepo{
		DB: db,
	}
}

func (r repositoryRepo) Create(repository models.CreateRepository)(string, error){
	uid := uuid.New()
	query := `INSERT INTO repositories values ($1, $2, $3, $4)`
	_, err := r.DB.Exec(query,
		uid,
		repository.ProductID,
		repository.BranchID,
		repository.Count,
	)
	if err != nil{
		fmt.Println("error while inserting to repositories!")
		return "", err
	}

	return uid.String(), nil
}

func (r repositoryRepo) GetByID(pKey models.PrimaryKey)(models.Repository, error){
	repository := models.Repository{}
	query :=  `SELECT id, product_id, branch_id, count,
			created_at, updated_at, deleted_at 
		from repositories where id = $1`
	err := r.DB.QueryRow(query,pKey.ID).Scan(
		&repository.ID,
		&repository.ProductID,
		&repository.BranchID,
		&repository.Count,
		&repository.CreatedAt,
		&repository.UpdatedAt,
		&repository.DeletedAt,
	)

	if err != nil{
		fmt.Println("error while selecting repository by id!")
		return models.Repository{},err
	}
	return repository, nil
}

func (r repositoryRepo) GetList(request models.GetListRequest) (models.RepositoriesResponse, error){
	var (
		repositories = []models.Repository{}
		count = 0
		countQuery, query string
		page   = request.Page
		offset = (page - 1) * request.Limit
	)

	countQuery = `
	SELECT count(1) from repositories`

	err := r.DB.QueryRow(countQuery).Scan(&count)
	if err != nil{
		fmt.Println("error while scanning count of repositories!")
		return models.RepositoriesResponse{}, err
	}

	query =  `SELECT id, product_id, branch_id, count,
		created_at, updated_at, deleted_at 
	from repositories where id = $1`


	query += `LIMIT $1 OFFSET $2`

	rows, err := r.DB.Query(query,request.Limit,offset)
	if err != nil{
		fmt.Println("error while query rows!", err.Error())
		return models.
		RepositoriesResponse{},err
	}

	for rows.Next(){
		repository := models.Repository{}

	 err := rows.Scan(
		&repository.ID,
		&repository.ProductID,
		&repository.BranchID,
		&repository.Count,
		&repository.CreatedAt,
		&repository.UpdatedAt,
		&repository.DeletedAt,
	)
		if err != nil{
			fmt.Println("error while scanning rows!", err.Error())
			return models.RepositoriesResponse{}, err
		}
		repositories = append(repositories, repository)
	}
	return models.RepositoriesResponse{
		Repositories: repositories,
		Count: count,
	}, nil
}

func (r repositoryRepo) Update(repository models.UpdateRepository)(string, error){
	query :=  `UPDATE repositories set product_id = $1, branch_id = $2, count = $3 where id = $4`
	uid, _ := uuid.Parse(repository.ID)
	_, err := r.DB.Exec(query,
		 repository.ProductID,
		 repository.BranchID,
		 repository.Count,
		 uid,
	)
	if err != nil{
		fmt.Println("error while updating repositories!", err.Error())
		return "", err
	}

	return repository.ID, nil
}

func (r repositoryRepo) Delete(pKey models.PrimaryKey) error{
	query := `DELETE from repositories where id = $1`
	_, err := r.DB.Exec(query,pKey)
	if err != nil{
		fmt.Println("error while deleting repositories!", err.Error())
		return err
	}
	return nil
}