package entity

type Transaction struct {
	ID     uint `gorm:"primaryKey"`
	Status TxStatus
}

type SubTransaction struct {
	ID           uint        `gorm:"primaryKey"`
	Transaction  Transaction `gorm:"foreignKey:ID"`
	WalletID     Wallet      `gorm:"foreignKey:ID"`
	Amount       uint
	TokenAddress string
	Status       SubTxStatus
	TxHash       string
	Timestamp    uint
}

type TxStatus uint8

const (
	TxStatusInit TxStatus = iota + 1
	TxStatusPending
	TxStatusFailed
	TxStatusDone
)

func (s TxStatus) IsValid() bool {
	if s >= TxStatusInit && s <= TxStatusDone {
		return true
	}
	return false
}

type SubTxStatus uint8

const (
	SubTxStatusPending SubTxStatus = iota + 1
	SubTxStatusDeposited
	SubTxStatusSubmitted
	SubTxStatusExecuted
	SubTxStatusRevoked
	SubTxStatusWithdrawn
)

func (s SubTxStatus) IsValid() bool {
	if s >= SubTxStatusPending && s <= SubTxStatusWithdrawn {
		return true
	}
	return false
}
