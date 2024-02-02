package entity

type Wallet struct {
	ID      uint `gorm:"primaryKey"`
	Address string
}
