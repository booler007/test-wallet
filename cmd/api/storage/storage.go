package storage

import (
	"time"

	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

type Wallet struct {
	uuid    string
	Balance float64
}

type Transaction struct {
	Uuid            string
	TransactionTime time.Time
	Wallet          string
	Operation       string
	Amount          float64
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{db}
}

func (s *Storage) CreateTransaction(new *Transaction) error {
	return s.DB.Create(new).Error
}

func (s *Storage) GetWallet(walletID string) (*Wallet, error) {
	var wl Wallet
	err := s.DB.Where("uuid = ?", walletID).First(&wl).Error
	if err != nil {
		return nil, err
	}

	return &wl, nil
}

func (s *Storage) UpdateBalance(wl *Wallet) error {
	return s.DB.Updates(&wl).Error
}
