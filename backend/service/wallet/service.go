package wallet

import (
	"backend/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spruceid/siwe-go"
)

type Service struct {
	Repo Repo
}

type Repo interface {
	GetWallet(address string) *entity.Wallet
	CreateWallet(address string, signature string)
	LoginWallet(address string)
	GetCreateWallet(address, signature string) *entity.Wallet
	GetWalletBySignature(signature string) *entity.Wallet
}

func (s *Service) SetSession(c *gin.Context) *entity.Wallet {
	return s.Repo.GetWallet(c.Param("address"))
}

func GenerateNonce(c *gin.Context) {
	nonce := siwe.GenerateNonce()
	c.JSON(200, gin.H{
		"nonce": nonce,
	})

}

func (s *Service) GetCreateWallet(address, signature string) *entity.Wallet {
	return s.Repo.GetCreateWallet(address, signature)
}

func (s *Service) LoginWallet(address string) {
	s.Repo.LoginWallet(address)
}

func (s *Service) ValidateSignature(signature string) (string, bool) {
	wallet := s.Repo.GetWalletBySignature(signature)
	fmt.Println(wallet)
	fmt.Println(wallet.Address)

	if wallet.Address == "" {
		return "", false
	}
	return wallet.Address, true
}
