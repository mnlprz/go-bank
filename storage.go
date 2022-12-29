package main

import (
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage interface {
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
	DeleteAccountByID(int) error
	GetAccounts() ([]*Account, error)
}

type PostgresStore struct {
	db *gorm.DB
}

func NewPostgresStore() (*PostgresStore, error) {

	dsn := "host=localhost user=root password=secret dbname=go-bank port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {

	err := s.db.AutoMigrate(&Account{})
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) CreateAccount(acc *Account) error {

	id := s.db.Create(&Account{
		FirstName: acc.FirstName,
		LastName:  acc.LastName,
		Number:    acc.Number,
		Balance:   acc.Balance,
		CreatedAt: acc.CreatedAt,
	})

	log.Println("Create accoun: ", id)

	return nil
}

func (s *PostgresStore) DeleteAccountByID(id int) error {

	s.db.Delete(&Account{}, id)
	return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error {

	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {

	account := new(Account)

	s.db.First(&account, id)

	return account, nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {

	/* var accounts []*Account
	result := s.db.Find(&accounts)
	result.Scan()

	for rows.Next() {
		account := new(Account)
		err := rows.Scan(
			&account.ID,
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

	return accounts, nil */
	return nil, nil
}
