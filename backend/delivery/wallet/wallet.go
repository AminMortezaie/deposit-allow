package delivery

import (
	"backend/repository/db/postgres"
	"backend/service/wallet"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spruceid/siwe-go"
	"log"
)

func Verify(c *gin.Context) {

	var dbConnection postgres.DBConnector
	var message *siwe.Message

	var body struct {
		Message   string
		Signature string
	}

	errBind := c.Bind(&body)
	if errBind != nil {
		panic("error in binding body:" + errBind.Error())
	}
	errDB := dbConnection.ConnectToPostgres()
	if errDB != nil {
		panic("error in connect db:" + errDB.Error())
	}

	message, _ = siwe.ParseMessage(body.Message)
	address := message.GetAddress().String()

	repo := &postgres.Repo{DBConnector: &dbConnection}
	walletSr := wallet.Service{Repo: repo}
	walletSr.GetCreateWallet(address, body.Signature)
	session := sessions.Default(c)
	session.Set("signature", body.Signature)
	errSession := session.Save()
	if errSession != nil {
		log.Fatal("error in saving session:" + errDB.Error())
	}
	walletSr.LoginWallet(address)
	c.JSON(200, gin.H{
		"address": address,
		"status":  "login success",
	})

}

func ValidateSession(c *gin.Context) {
	var dbConnection postgres.DBConnector
	errDB := dbConnection.ConnectToPostgres()
	if errDB != nil {
		panic("error in connect db:" + errDB.Error())
	}

	session := sessions.Default(c)
	s := session.Get("signature")
	if s == nil {
		c.JSON(401, gin.H{
			"address": "",
			"status":  "not authenticated",
		})
	}
	repo := &postgres.Repo{DBConnector: &dbConnection}
	walletSr := wallet.Service{Repo: repo}
	address, isExist := walletSr.ValidateSignature(s.(string))
	if isExist {
		c.JSON(200, gin.H{
			"address": address,
			"status":  "authenticated",
		})
	} else {
		c.JSON(401, gin.H{
			"address": address,
			"status":  "not authenticated",
		})

	}

}
