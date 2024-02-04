package entity

import (
	"gorm.io/gorm"
	"time"
)

type Wallet struct {
	gorm.Model
	Address   string
	Signature string
	LoginAt   time.Time
}
