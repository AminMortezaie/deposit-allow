package wallet

import (
	"backend/entity"
	"github.com/gin-gonic/gin"
	"github.com/spruceid/siwe-go"
)

type Service struct {
	Repo Repo
}

type Repo interface {
	GetWallet(address string) (*entity.Wallet, error)
	CreateWallet(address string, signature string) (*entity.Wallet, error)
	LoginWallet(address string) (*entity.Wallet, error)
	GetCreateWallet(address string, signature string) (*entity.Wallet, error)
}

func (s *Service) SetSession(c *gin.Context) (*entity.Wallet, error) {
	return s.Repo.GetWallet(c.Param("address"))
}

func GenerateNonce(c *gin.Context) {
	nonce := siwe.GenerateNonce()
	c.JSON(200, gin.H{
		"nonce": nonce,
	})

}

func (s *Service) GetCreateWallet(address, signature string) (*entity.Wallet, error) {
	return s.Repo.GetCreateWallet(address, signature)
}

func (s *Service) LoginWallet(address string) (*entity.Wallet, error) {
	return s.Repo.LoginWallet(address)
}
