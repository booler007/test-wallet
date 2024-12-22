package service

import (
	"time"

	"wallet/cmd/api/storage"

	"github.com/google/uuid"
)

type Storager interface {
	CreateTransaction(transaction *storage.Transaction) error
	GetWallet(walletID string) (*storage.Wallet, error)
	UpdateBalance(wl *storage.Wallet) error
}
type Service struct {
	Storage Storager
}

func NewService(str Storager) *Service {
	return &Service{str}
}

func (s *Service) ProcessTheTransaction(walletID, operation string, amount float64) error {
	wallet, err := s.Storage.GetWallet(walletID)
	if err != nil {
		return err
	}

	if operation == "WITHDRAW" {
		amount = -amount
	}

	wallet.Balance += amount
	err = s.Storage.UpdateBalance(wallet)
	if err != nil {
		return err
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	return s.Storage.CreateTransaction(&storage.Transaction{
		Uuid:            id.String(),
		TransactionTime: time.Now(),
		Operation:       operation,
		Amount:          amount,
		Wallet:          walletID,
	})
}
