package postgres

import (
	"context"
	"developer/api/models"
	"developer/storage"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type tarifRepo struct {
	DB *pgxpool.Pool
}

func NewTarifRepo(db *pgxpool.Pool) storage.ITarif {
	return &tarifRepo{
		DB: db,
	}
}

func (t *tarifRepo) Create(ctx context.Context, createTarif models.CreateTarif) (string, error) {

	uid := uuid.New()

	query := `insert into tarifs (id, name, tarif_type, amount_for_cash,
		amount_for_card) 
	values 
	($1, $2, $3, $4, $5)`

	_, err := t.DB.Exec(ctx, query,
		uid,
		createTarif.Name,
		createTarif.TarifType,
		createTarif.AmountForCash,
		createTarif.AmountForCard,
	)
	if err != nil {
		fmt.Println("error while inserting tarif data", err.Error())
		return "", err
	}

	return uid.String(), nil

}

func (t *tarifRepo) GetByID(ctx context.Context, pKey models.PrimaryKey) (models.Tarif, error) {

	tarif := models.Tarif{}

	query := `select id, name, tarif_type, amount_for_cash,
	amount_for_card, 
	 created_at, updated_at from tarifs
	 where id = $1`

	row := t.DB.QueryRow(ctx, query, pKey.ID)

	err := row.Scan(
		&tarif.ID,
		&tarif.Name,
		&tarif.TarifType,
		&tarif.AmountForCash,
		&tarif.AmountForCard,
		&tarif.CreatedAt,
		&tarif.UpdatedAt,
	)

	if err != nil {
		fmt.Println("error while selecting tarif data", err.Error())
		return models.Tarif{}, err
	}

	return tarif, nil

}

func (t *tarifRepo) GetList(ctx context.Context, request models.GetListRequest) (models.TarifsResponse, error) {

	var (
		tarifs            = []models.Tarif{}
		count             = 0
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `select count(1) from tarif `

	if search != "" {
		countQuery += fmt.Sprintf(` where name ilike '%%%s%%'`, search)
	}
	if err := t.DB.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting tarif count", err.Error())
		return models.TarifsResponse{}, err
	}

	query = `select id, name, tarif_type, 
	amount_for_cash, amount_for_card,
	created_at, updated_at
	from tarifs `

	if search != "" {
		query += fmt.Sprintf(` where name ilike '%%%s%%'`, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := t.DB.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting tarif", err.Error())
		return models.TarifsResponse{}, err
	}

	for rows.Next() {
		tarif := models.Tarif{}
		if err = rows.Scan(
			&tarif.ID,
			&tarif.Name,
			&tarif.TarifType,
			&tarif.AmountForCash,
			&tarif.AmountForCard,
			&tarif.CreatedAt,
			&tarif.UpdatedAt,
		); err != nil {
			fmt.Println("error is while scanning tarif data", err.Error())
			return models.TarifsResponse{}, err
		}

		tarifs = append(tarifs, tarif)

	}

	return models.TarifsResponse{
		Tarifs: tarifs,
		Count:  count,
	}, nil
}

func (t *tarifRepo) Update(ctx context.Context, request models.UpdateTarif) (string, error) {

	query := `update tarif
   set name = $1, tarif_type = $2, amount_for_cash = $3,
   amount_for_card = $4, updated_at = now()
   where id = $5
   `
	_, err := t.DB.Exec(ctx, query,
		request.Name,
		request.TarifType,
		request.AmountForCash,
		request.AmountForCard,
		request.ID,
	)
	if err != nil {
		fmt.Println("error while updating tarif data...", err.Error())
		return "", err
	}
	return "", nil
}

func (t *tarifRepo) Delete(ctx context.Context, pKey models.PrimaryKey) error {

	query := `
	DELETE FROM tarifs where id = $1
	`

	_, err := t.DB.Exec(ctx, query, pKey.ID)
	if err != nil {
		fmt.Println("error while deleting tarif by id", err.Error())
		return err
	}

	return nil
}