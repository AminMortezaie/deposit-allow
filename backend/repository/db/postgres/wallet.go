package postgres

import (
	"backend/entity"
	"fmt"
	"time"
)

type Repo struct {
	DBConnector *DBConnector
}

func (r Repo) GetWallet(address string) (*entity.Wallet, error) {
	var wallet = entity.Wallet{Address: address}
	err := r.DBConnector.DB.First(&wallet)
	if err != nil {
		return nil, fmt.Errorf("cannot retrieve wallet with address %s: %w", address, err)
	}
	return &wallet, nil
}

func (r Repo) CreateWallet(address string, signature string) (*entity.Wallet, error) {
	var wallet = entity.Wallet{Address: address, Signature: signature}
	err := r.DBConnector.DB.Create(&wallet)
	if err != nil {
		return nil, fmt.Errorf("cannot create wallet with address %s: %w", address, err)
	}
	return &wallet, nil
}

func (r Repo) LoginWallet(address string) (*entity.Wallet, error) {
	var wallet = entity.Wallet{Address: address, LoginAt: time.Now()}
	err := r.DBConnector.DB.Save(&wallet)
	if err != nil {
		return nil, fmt.Errorf("cannot retrieve wallet with address %s: %w", address, err)
	}
	return &wallet, nil
}

func (r Repo) GetCreateWallet(address string, signature string) (*entity.Wallet, error) {
	var wallet = entity.Wallet{Address: address, Signature: signature}
	err := r.DBConnector.DB.FirstOrCreate(&wallet)
	if err != nil {
		return nil, fmt.Errorf("cannot retrieve wallet with address %s: %w", address, err)
	}
	return &wallet, nil
}
