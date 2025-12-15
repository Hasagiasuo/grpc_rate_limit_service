package postgresql

import (
	"fmt"
	"rlservice/internal/config"
	"rlservice/internal/domain/models"
	"rlservice/pkg/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const DEFAULT_BALANCE = 5

type PsqlStorage struct {
	db  *sqlx.DB
	log *logger.Logger
	cfg *config.PsqlConfig
}

func NewPsqlStorage(log *logger.Logger, cfg *config.PsqlConfig) (*PsqlStorage, error) {
	const op = "postgresql.NewPsqlStorage"
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.DbName, cfg.SslMode))
	if err != nil {
		panic(fmt.Sprintf("cannot open postgres by config: %v", err))
	}
	if err := updateTable(db, log); err != nil {
		return nil, err
	}
	log.Info(op, "postgres success inited")
	return &PsqlStorage{
		db:  db,
		log: log,
		cfg: cfg,
	}, nil
}

func (ps *PsqlStorage) GetServiceByToken(token string) *models.Service {
	var service models.Service
	query := `select token, balance from services where token = $1`
	if err := ps.db.Get(&service, query, token); err != nil {
		return nil
	}
	return &service
}

func (ps *PsqlStorage) AddNewService(token string) error {
	const op = "postgresql.AddNewService"
	query := `insert into services (token, balance) values ($1, $2);`
	if _, err := ps.db.Exec(query, token, DEFAULT_BALANCE); err != nil {
		ps.log.Error(op, fmt.Sprintf("cannot add new service: %v", err))
		return ErrCannotAddNewService
	}
	return nil
}

func (ps *PsqlStorage) RemoveService(token string) error {
	const op = "postgresql.RemoveService"
	query := `delete from services where token = $1`
	if _, err := ps.db.Exec(query, token); err != nil {
		ps.log.Error(op, fmt.Sprintf("cannot delete service by token [%s]: %v", token, err))
		return ErrCannotDeleteService
	}
	return nil
}

func (ps *PsqlStorage) UpdateServiceBalance(token string, cost int) error {
	const op = "postgresql.UpdateServiceBalance"
	service := ps.GetServiceByToken(token)
	if service == nil {
		return ErrServiceNotFound
	}
	query := `update services set balance = $1 where token = $2`
	if _, err := ps.db.Exec(query, service.Balance - cost, token); err != nil {
		ps.log.Error(op, fmt.Sprintf("cannot update balance by service: %v", err))
		return ErrCannotUpdateServiceBalance
	}
	return nil
}

func updateTable(db *sqlx.DB, log *logger.Logger) error {
	const op = "postgresql.updateTable"
	query := `create table if not exists services (
		token TEXT, 
		balance INTEGER
	);`
	if _, err := db.Exec(query); err != nil {
		log.Error(op, fmt.Sprintf("cannot create services table: %v", err))
		return ErrCannotCreateServiceTable
	}
	return nil
}
