package delivery

import (
	"backend/repository/db/postgres"
	"backend/service/wallet"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spruceid/siwe-go"
)

func Verify(c *gin.Context) {
	var dbConnection postgres.DBConnector
	var message *siwe.Message

	var body struct {
		Message   string
		Signature string
	}
	err := dbConnection.ConnectToPostgres()
	repo := postgres.Repo{DBConnector: &dbConnection}
	walletSr := wallet.Service{Repo: repo}

	err = c.Bind(&body)
	if err != nil {
		return
	}
	message, err = siwe.ParseMessage(body.Message)
	address := message.GetAddress().String()
	_, err = walletSr.GetCreateWallet(address, body.Signature)
	if err != nil {
		panic("error in creating wallet:" + err.Error())
	}
	session := sessions.Default(c)
	session.Set("signature", body.Signature)
	err = session.Save()
	_, err = walletSr.LoginWallet(address)

	c.JSON(200, gin.H{
		"address": address,
		"status":  "login success",
	})

}

func ValidateSession(c *gin.Context) {
	session := sessions.Default(c)
	s := session.Get("signature")
	fmt.Printf("%s %T \n", "this is type:::", s)
}
