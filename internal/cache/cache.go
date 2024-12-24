package cache

import (
	"log"
	"sync"
	"time"

	"wallet/internal/storage"
)

type Cache struct {
	existWallets   map[string]int
	updatedWallets map[string]int
	mu             sync.RWMutex
	str            Storager
}

type Storager interface {
	GetWallets() ([]storage.Wallet, error)
	UpdateBalances(map[string]int) error
}

func InitCache(str Storager) (*Cache, error) {
	wallets, err := str.GetWallets()
	if err != nil {
		return nil, err
	}

	mapWallets := make(map[string]int, len(wallets))
	for _, wallet := range wallets {
		mapWallets[wallet.Uuid] = wallet.Balance
	}

	updatedWallets := make(map[string]int, len(wallets))
	c := &Cache{existWallets: mapWallets, updatedWallets: updatedWallets, str: str}

	go func() {
		for range time.Tick(time.Minute) {
			c.mu.Lock()

			if err = c.str.UpdateBalances(c.updatedWallets); err != nil {
				log.Println(err)
			}

			clear(c.updatedWallets)
			c.mu.Unlock()
		}
	}()

	return c, nil
}

func (c *Cache) GetBalance(walletID string) (int, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	w, found := c.existWallets[walletID]
	return w, found
}

func (c *Cache) UpdateBalance(walletID string, amount int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.existWallets[walletID] = amount
	c.updatedWallets[walletID] = amount
}
