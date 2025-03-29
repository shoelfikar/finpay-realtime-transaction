package repository

import (
	"database/sql"

	"github.com/shoelfikar/finpay-realtime-transaction/model"
)

type walletRepository struct {
	DB *sql.DB
}

type WalletRepository interface {
	CreateWallet(wallet model.Wallet) model.Wallet
}

func NewWalletRepository(db *sql.DB) WalletRepository {
	return &walletRepository{
		DB: db,
	}
}

func (w *walletRepository) CreateWallet(wallet model.Wallet) model.Wallet {
	return wallet
}