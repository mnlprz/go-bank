package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
	DeleteAccountByID(int) error
	GetAccounts() ([]*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {

	dsn := "host=localhost user=root password=secret dbname=go-bank port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {

	err := s.createAccountTable()

	return err
}

func (s *PostgresStore) createAccountTable() error {

	q := `CREATE TABLE IF NOT EXISTS ACCOUNT(
			ID SERIAL PRIMARY KEY,
			FIRST_NAME VARCHAR(50),
			LAST_NAME VARCHAR(50),
			NUMBER SERIAL,
			BALANCE SERIAL,
			CREATED_AT TIMESTAMP
		)`

	_, err := s.db.Exec(q)

	return err
}

func (s *PostgresStore) CreateAccount(acc *Account) error {

	q := `INSERT INTO ACCOUNT(
			FIRST_NAME,
			LAST_NAME,
			NUMBER,
			BALANCE,
			CREATED_AT
			) VALUES
			($1, $2, $3, $4, $5)`

	_, err := s.db.Exec(q,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.Balance,
		acc.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) DeleteAccountByID(id int) error {

	q := `DELETE FROM ACCOUNT WHERE ID = $1`
	r, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}

	if rf, _ := r.RowsAffected(); rf == 0 {
		return fmt.Errorf("id: %d not found to delete", id)
	}

	return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error {

	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {

	q := `SELECT * FROM ACCOUNT WHERE ID = $1`

	account := new(Account)

	err := s.db.QueryRow(q, id).Scan(
		&account.Id,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("id: %d not found", id)
		}
		return nil, err
	}

	return account, nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {

	rows, err := s.db.Query(`SELECT * FROM ACCOUNT`)
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}

	for rows.Next() {
		account := new(Account)
		err := rows.Scan(
			&account.Id,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}
