package postgres

import (
	"context"
	"developer/config"
	"developer/storage"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database"          //database is needed for migration
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //postgres is used for database
	_ "github.com/golang-migrate/migrate/v4/source/file"       //file is needed for migration url
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type Store struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Config) (storage.IfStorage, error) {
	url := fmt.Sprintf(
		`postgres://%s:%s@%s:%s/%s?sslmode=disable`,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		fmt.Println("error while parsing config", err.Error())
		return nil, err
	}

	poolConfig.MaxConns = 100

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil{
		fmt.Println("error while connecting to db", err.Error())
		return nil, err
	}

	return Store{
		Pool: pool,
	},nil
}


func (s Store) Close(){
	s.Pool.Close()
}

func (s Store) Branch() storage.IBranch{
	return NewBranchRepo(s.Pool)
}

func (s Store) Sale() storage.ISale{
	return NewSaleRepo(s.Pool)
}

func (s Store) Product() storage.IProduct{
	return NewProductRepo(s.Pool)
}

func (s Store) Basket() storage.IBasket {
	return NewBasketRepo(s.Pool)
}

func (s Store) Storage() storage.IStorage {
	return NewStorageRepo(s.Pool)
}
//New
func (s Store) Staff() storage.IStaff{
	return NewStaffRepo(s.Pool)
}

func (s Store) Tarif() storage.ITarif{
	return NewTarifRepo(s.Pool)
}

func (s Store) Category() storage.ICategory {
	return NewCategoryRepo(s.Pool)
}

func (s Store) Transaction() storage.ITransaction {
	return NewTransactionRepo(s.Pool)
}

func (s Store) StorageTransaction()storage.IStorageTransaction  {
	return NewStorageTransactionRepo(s.Pool)
}