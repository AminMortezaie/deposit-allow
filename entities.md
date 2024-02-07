# Entities

### Introduction

The `entity` package defines transaction models used within the application to manage transactions and sub-transactions. This documentation provides an overview of the transaction models and wallet model, their attributes, and functionalities.

### Transaction Model

The `Transaction` model represents a high-level transaction entity within the application. It captures the status of the transaction using enumerated values.

Sure, here's the documentation for the transaction models:

```go
package entity

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Status TxStatus
}
```

#### Attributes

* `ID`: Primary key of the transaction.
* `CreatedAt`, `UpdatedAt`, `DeletedAt`: Standard GORM fields for tracking the creation, update, and deletion timestamps of the transaction.
* `Status`: Status of the transaction, represented by the `TxStatus` enumerated type.

#### Enumerated Types

The `TxStatus` enumerated type defines the possible statuses of a transaction:

```go
type TxStatus uint8

const (
	TxStatusInit TxStatus = iota + 1
	TxStatusPending
	TxStatusFailed
	TxStatusDone
)
```

### SubTransaction Model

The `SubTransaction` model represents a sub-level transaction entity associated with a parent transaction. It captures details such as the wallet ID, amount, token address, status, transaction hash, and timestamp.

```go
type SubTransaction struct {
	gorm.Model
	Transaction  Transaction `gorm:"foreignKey:ID"`
	WalletID     Wallet      `gorm:"foreignKey:ID"`
	Amount       uint
	TokenAddress string
	Status       SubTxStatus
	TxHash       string
	Timestamp    uint
}
```

#### Attributes

* `ID`: Primary key of the sub-transaction.
* `CreatedAt`, `UpdatedAt`, `DeletedAt`: Standard GORM fields for tracking the creation, update, and deletion timestamps of the sub-transaction.
* `Transaction`: Reference to the parent transaction associated with the sub-transaction.
* `WalletID`: Reference to the wallet involved in the sub-transaction.
* `Amount`: Amount of tokens transacted.
* `TokenAddress`: Address of the token involved in the transaction.
* `Status`: Status of the sub-transaction, represented by the `SubTxStatus` enumerated type.
* `TxHash`: Hash of the transaction.
* `Timestamp`: Timestamp of the transaction.

#### Enumerated Types

The `SubTxStatus` enumerated type defines the possible statuses of a sub-transaction:

```go
type SubTxStatus uint8

const (
	SubTxStatusPending SubTxStatus = iota + 1
	SubTxStatusDeposited
	SubTxStatusSubmitted
	SubTxStatusExecuted
	SubTxStatusRevoked
	SubTxStatusWithdrawn
)
```

### Wallet Model

The `Wallet` model represents a user's wallet entity within the application. It captures details such as the wallet address, signature, and login timestamp.

```go
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
```

#### Attributes

* `ID`: Primary key of the wallet.
* `CreatedAt`, `UpdatedAt`, `DeletedAt`: Standard GORM fields for tracking the creation, update, and deletion timestamps of the wallet.
* `Address`: Wallet address, uniquely identifying the wallet within the application.
* `Signature`: Signature associated with the wallet.
* `LoginAt`: Timestamp indicating the last login time of the wallet.

### Conclusion

The models provided by the `entity` package serve as fundamental components for managing entity-related data within the application. By leveraging these models, developers can effectively capture, track, and process entity transactions, ensuring integrity and consistency in data management and transaction processing.
