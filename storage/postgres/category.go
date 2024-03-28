package postgres

import (
	"context"
	"developer/api/models"
	"developer/storage"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type categoryRepo struct {
	DB *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) storage.ICategory{
	return &categoryRepo{
		DB: db,
	}
}

func (c *categoryRepo) Create(ctx context.Context,category models.CreateCategory) (string, error) {

	uid := uuid.New()

	query := `INSERT INTO categories (id, name, parent_id) values ($1, $2, $3)`

	_, err := c.DB.Exec(ctx,query, 
		uid, 
		category.Name, 
		category.ParentID,
	)
	if err != nil{
		fmt.Println("Error while inserting into categories!",err.Error())
		return "", err
	}

	return uid.String(), nil
}

func (c *categoryRepo) GetByID(ctx context.Context,pKey models.PrimaryKey) (models.Category, error) {
	query := `SELECT id, name, parent_id, created_at, updated_at from categories
	WHERE id = $1`

	category := models.Category{}

	err := c.DB.QueryRow(ctx,query, pKey.ID).Scan(
		&category.ID,
		&category.Name,
		&category.ParentID,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil{
		fmt.Println("Error while selecting category by id!", err.Error())
		return models.Category{},err
	}
	
	return category, nil
}

func (c *categoryRepo) GetList(ctx context.Context,request models.GetListRequest) (models.CategoriesResponse, error) {
	var (
		categories        = []models.Category{}
		count             = 0
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `SELECT count(1) from categories`

	if search != "" {
		countQuery += fmt.Sprintf(` where name ilike '%%%s%%'`, search)
	}

	err := c.DB.QueryRow(ctx,countQuery).Scan(&count)
	if err != nil {
		fmt.Println("Error is while selecting count", err.Error())
		return models.CategoriesResponse{}, err
	}

	query = `SELECT id, name, parent_id, created_at, updated_at from categories`

	if search != "" {
		query += fmt.Sprintf(` where name ilike '%%%s%%'`, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := c.DB.Query(ctx,query, request.Limit, offset)
	if err != nil {
		fmt.Println("Error is while selecting category", err.Error())
		return models.CategoriesResponse{}, err
	}

	for rows.Next() {
		category := models.Category{}
		 err := rows.Scan(
			&category.ID,
		    &category.Name,
			&category.ParentID, 
			&category.CreatedAt, 
			&category.UpdatedAt,
		 )
		if err != nil {
			fmt.Println("Error is while scanning category data", err.Error())
			return models.CategoriesResponse{}, err
		}
		categories = append(categories, category)
	}

	return models.CategoriesResponse{
		Categories: categories,
		Count:      count,
	}, nil
}

func (c *categoryRepo) Update(ctx context.Context, updateCat models.UpdateCategory) (string, error) {

	query := `UPDATE categories
   set name = $1, parent_id = $2, updated_at = now()  where id = $3
   `

	_,err := c.DB.Exec(ctx,query, 
		 updateCat.Name,
		 updateCat.ParentID,
		 updateCat.ID,
	)
	if err != nil {
		fmt.Println("Error while updating category!", err.Error())
		return "", err
	}

	return updateCat.ID, nil
}

func (c *categoryRepo) Delete(ctx context.Context, pKey models.PrimaryKey) error {

	query := `DELETE from categories where id = $1`
	 

	_, err := c.DB.Exec(ctx,query, pKey.ID)
	if err != nil {
		fmt.Println("Error while deleting categories!", err.Error())
		return err
	}

	return nil
}