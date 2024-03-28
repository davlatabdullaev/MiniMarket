package postgres

import (
	"context"
	"developer/api/models"
	"developer/pkg/check"
	"developer/storage"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type staffRepo struct {
	DB *pgxpool.Pool
}

func NewStaffRepo(db *pgxpool.Pool) storage.IStaff {
	return &staffRepo{
		DB: db,
	}
}

func (s *staffRepo) Create(ctx context.Context,createStaff models.CreateStaff) (string, error) {

	uid := uuid.New()

	query := `INSERT INTO staffs (id, branch_id, tarif_id, type_staff, name, balance, 
		birth_date, age, gender, login, password) 
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := s.DB.Exec(ctx,query,
		uid,
		createStaff.BranchID,
		createStaff.TarifID,
		createStaff.TypeStaff,
		createStaff.Name,
		createStaff.Balance,
		createStaff.BirthDate,
		check.CalculateAge(createStaff.BirthDate),
		createStaff.Gender,
		createStaff.Login,
		createStaff.Password,
	)
	if err != nil {
		fmt.Println("Error while inserting into staff!", err.Error())
		return "", err
	}

	return uid.String(), nil
}

func (s *staffRepo) GetByID(ctx context.Context,pKey models.PrimaryKey) (models.Staff, error) {

	staff := models.Staff{}

	query := `SELECT id, branch_id, tarif_id, type_staff, name, 
	birth_date, age, gender, login, password, created_at, updated_at 
	from staffs where id = $1`

	row := s.DB.QueryRow(ctx,query, pKey.ID)

	err := row.Scan(
		&staff.ID,
		&staff.BranchID,
		&staff.TarifID,
		&staff.TypeStaff,
		&staff.Name,
		&staff.BirthDate,
		&staff.Age,
		&staff.Gender,
		&staff.Login,
		&staff.Password,
		&staff.CreatedAt,
		&staff.UpdatedAt,
	)

	if err != nil {
		fmt.Println("Error while selecting staff by id!", err.Error())
		return models.Staff{}, err
	}

	return staff, nil
}

func (s *staffRepo) GetList(ctx context.Context,request models.GetListRequest) (models.StaffsResponse, error) {
	var (
		staffs            = []models.Staff{}
		count             = 0
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `select count(1) from staffs`

	if search != "" {
		countQuery += fmt.Sprintf(` where name ilike '%%%s%%'`, search)
	}
	if err := s.DB.QueryRow(ctx,countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting staff count", err.Error())
		return models.StaffsResponse{}, err
	}

	query = `select id, branch_id, tarif_id, type_staff, name, birth_date, 
	age, gender, login, password, created_at, updated_at 
	from staffs`

	if search != "" {
		query += fmt.Sprintf(` where name ilike '%%%s%%'`, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := s.DB.Query(ctx,query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting staff", err.Error())
		return models.StaffsResponse{}, err
	}

	for rows.Next() {
		staff := models.Staff{}
		err := rows.Scan(
			&staff.ID,
			&staff.BranchID,
			&staff.TarifID,
			&staff.TypeStaff,
			&staff.Name,
			&staff.BirthDate,
			&staff.Age,
			&staff.Gender,
			&staff.Login,
			&staff.Password,
			&staff.CreatedAt,
			&staff.UpdatedAt,
		)
		if err != nil {
			fmt.Println("error is while scanning staff data", err.Error())
			return models.StaffsResponse{}, err
		}

		staffs = append(staffs, staff)

	}

	return models.StaffsResponse{
		Staffs: staffs,
		Count:  count,
	}, nil
}

func (s *staffRepo) Update(ctx context.Context,updateStaff models.UpdateStaff) (string, error) {

	query := `UPDATE staffs
   set branch_id = $1, tarif_id = $2, type_staff = $3,
   name = $4, birth_date = $5, age = $6, gender = $7, login = $8, password = $9, updated_at = now()
   where id = $11
   `
	_, err := s.DB.Exec(ctx,query,
		updateStaff.BranchID,
		updateStaff.TarifID,
		updateStaff.TypeStaff,
		updateStaff.Name,
		updateStaff.BirthDate,
		check.CalculateAge(updateStaff.BirthDate),
		updateStaff.Gender,
		updateStaff.Login,
		updateStaff.Password,
		updateStaff.ID,
	)
	if err != nil {
		fmt.Println("error while updating staff data...", err.Error())
		return "", err
	}

	return updateStaff.ID, nil
}

func (s *staffRepo) Delete(ctx context.Context, pKey models.PrimaryKey) error {

	query := `
	DELETE from staffs where id = $1
	`

	_, err := s.DB.Exec(ctx,query, pKey.ID)
	if err != nil {
		fmt.Println("error while deleting staff by id", err.Error())
		return err
	}

	return nil
}

func (s *staffRepo) UpdateSalary(ctx context.Context, pKey models.PrimaryKey, balance int)error{
	query := `UPDATE staffs set balance = balance + $1 where id = $2 `

	_, err := s.DB.Exec(ctx, query, balance, pKey.ID)
	if err != nil{
		fmt.Println("error while updating balance of staff!", err.Error())
		return err
	}
	return nil
}