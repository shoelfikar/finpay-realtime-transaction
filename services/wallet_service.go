package services

import (
	"github.com/shoelfikar/finpay-realtime-transaction/model"
	"github.com/shoelfikar/finpay-realtime-transaction/repository"
)

type walletService struct {
	WalletRepository repository.WalletRepository
}

type WalletService interface {
	CreateWallet(wallet model.Wallet) model.Wallet
}

func NewWalletService(wallet repository.WalletRepository) WalletService {
	return &walletService{
		WalletRepository: wallet,
	}
}

func (w *walletService) CreateWallet(wallet model.Wallet) model.Wallet {
	return wallet
}