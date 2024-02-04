package postgres

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

//var (
//	host     = os.Getenv("POSTGRES_HOST")
//	port, _  = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
//	user     = os.Getenv("POSTGRES_USER")
//	password = os.Getenv("POSTGRES_PASSWORD")
//	dbname   = os.Getenv("POSTGRES_DB")
//)

var (
	host       = "localhost"
	port       = 5432
	user       = "deposit_user"
	password   = "123456789"
	dbname     = "postgres"
	dbInstance *gorm.DB
	once       sync.Once
)

type DBConnector struct {
	DB *gorm.DB
}

func ConnectToPostgreSQL() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// ConnectToPostgres singleton
func (dbConnect *DBConnector) ConnectToPostgres() error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	once.Do(func() {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		dbInstance = db
	})
	dbConnect.DB = dbInstance
	return nil
}
