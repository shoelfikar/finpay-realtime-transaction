package repository

import (
	"database/sql"

	"github.com/shoelfikar/finpay-realtime-transaction/model"
)

type transactionRepository struct {
	DB *sql.DB
}

type TransactionRepository interface {
	Transaction(trx model.Transaction) model.Transaction
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{
		DB: db,
	}
}

func (t *transactionRepository) Transaction(trx model.Transaction) model.Transaction {
	return trx
}