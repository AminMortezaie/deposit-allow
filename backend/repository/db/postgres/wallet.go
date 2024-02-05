package postgres

import (
	"backend/entity"
	"fmt"
	"time"
)

type Repo struct {
	DBConnector *DBConnector
}

func (r *Repo) GetWallet(address string) *entity.Wallet {
	var wallet = entity.Wallet{}
	r.DBConnector.DB.First(&wallet, &entity.Wallet{Address: address})
	return &wallet
}
func (r *Repo) GetWalletBySignature(signature string) *entity.Wallet {
	wallet := entity.Wallet{}
	r.DBConnector.DB.First(&wallet, &entity.Wallet{Signature: signature})
	return &wallet
}

func (r *Repo) CreateWallet(address string, signature string) {
	var wallet = entity.Wallet{Address: address, Signature: signature}
	r.DBConnector.DB.Create(&wallet)
}

func (r *Repo) LoginWallet(address string) {
	var wallet = entity.Wallet{}
	r.DBConnector.DB.First(&wallet, &entity.Wallet{Address: address})
	wallet.LoginAt = time.Now()
	r.DBConnector.DB.Save(&wallet)
}

func (r *Repo) GetCreateWallet(address, signature string) *entity.Wallet {
	wallet := entity.Wallet{}

	r.DBConnector.DB.First(&wallet, &entity.Wallet{Address: address})
	fmt.Println(wallet)
	fmt.Println(wallet.Address)
	if wallet.Address == "" {
		r.DBConnector.DB.Save(&entity.Wallet{Address: address, Signature: signature})
		r.DBConnector.DB.First(&wallet, &entity.Wallet{Address: address})
	}
	return &wallet
}
