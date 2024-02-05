package main

import (
	delivery "backend/delivery/wallet"
	"backend/entity"
	"backend/repository/db/postgres"
	"backend/service/wallet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	var db postgres.DBConnector
	err := db.ConnectToPostgres()
	db.DB.Debug()

	// Perform database migration
	err = db.DB.AutoMigrate(&entity.Wallet{})
	err = db.DB.AutoMigrate(&entity.Transaction{})
	err = db.DB.AutoMigrate(&entity.SubTransaction{})

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"} // Adjust as needed
	config.AllowCredentials = true
	r.Use(cors.New(config))

	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24})

	r.Use(sessions.Sessions("session", store))
	r.POST("/verify", delivery.Verify)
	r.GET("/nonce", wallet.GenerateNonce)
	r.GET("/login", delivery.ValidateSession)
	err = r.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}

}
