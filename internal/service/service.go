package service

import (
	"errors"
)

type Cacher interface {
	UpdateBalance(string, int)
	GetBalance(string) (int, bool)
}

type Service struct {
	Cache Cacher
}

var (
	ErrInsufficientBalance = errors.New("insufficient balance for this transaction")
	ErrWalletNotFound      = errors.New("wallet not exist")
)

func NewService(c Cacher) *Service {
	return &Service{c}
}

func (s *Service) ProcessTheTransaction(walletID, operation string, amount int) error {
	if balance, exist := s.Cache.GetBalance(walletID); exist {
		if operation == "WITHDRAW" {
			amount = -amount
		}

		newBalance := balance + amount
		if newBalance < 0 {
			return ErrInsufficientBalance
		}

		s.Cache.UpdateBalance(walletID, newBalance)

		return nil
	}

	return ErrWalletNotFound
}

func (s *Service) GetBalance(walletID string) (int, error) {
	if balance, exist := s.Cache.GetBalance(walletID); exist {
		return balance, nil
	}

	return 0, ErrWalletNotFound
}
