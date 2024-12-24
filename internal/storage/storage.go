package storage

import (
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

type Wallet struct {
	Uuid    string
	Balance int
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{db}
}

func (s *Storage) GetWallets() (wallets []Wallet, err error) {
	res := s.DB.Find(&wallets)

	return wallets, res.Error
}

func (s *Storage) UpdateBalances(wallets map[string]int) (err error) {
	for wallet, balance := range wallets {
		if err = s.DB.Model(&Wallet{}).Where("uuid = ?", wallet).Update("balance", balance).Error; err != nil {
			return err
		}
	}

	return nil
}
