# Repository

The `repository` package within the application is responsible for managing data persistence and database operations. It includes a `db` folder containing multiple database connectors, and for the Postgres database, it contains two files: `db.go` and `wallet.go`.

### Database Connector (`db.go`)

The `db.go` file provides functionality to establish a connection to the Postgres database using GORM. It defines a `DBConnector` struct and methods to connect to the database.\


**Attributes**

* `DB`: A pointer to the `gorm.DB` instance representing the database connection.

This method initializes a connection to the PostgreSQL database using the provided connection details. It utilizes a **singleton pattern** to ensure a single database connection instance is created.

* `dsn`: Data Source Name used to connect to the PostgreSQL database.
* `once.Do`: Ensures that the connection is established only once.
* `gorm.Open`: Establishes a connection to the PostgreSQL database using GORM.

### Wallet Repository

The `wallet.go` file contains the `Repo` type, responsible for performing CRUD operations on wallet entities within the PostgreSQL database.

#### Attributes

* `DBConnector`: A pointer to the `DBConnector` instance for interacting with the PostgreSQL database.

#### GetWallet

```go
// GetWallet retrieves a wallet by its address.
func (r *Repo) GetWallet(address string) *entity.Wallet {
    var wallet = entity.Wallet{}
    r.DBConnector.DB.First(&wallet, &entity.Wallet{Address: address})
    return &wallet
}
```

**Description:** This method retrieves a wallet entity from the database based on its address. It queries the database using the provided address and populates the `wallet` variable with the retrieved data. It then returns a pointer to the wallet entity.

#### GetWalletBySignature

```go
// GetWalletBySignature retrieves a wallet by its signature.
func (r *Repo) GetWalletBySignature(signature string) *entity.Wallet {
    var wallet = entity.Wallet{}
    r.DBConnector.DB.First(&wallet, &entity.Wallet{Signature: signature})
    return &wallet
}
```

**Description:** Similar to `GetWallet`, this method retrieves a wallet entity from the database but searches by signature instead of address. It queries the database using the provided signature and populates the `wallet` variable with the retrieved data. It then returns a pointer to the wallet entity.

#### CreateWallet

```go
// CreateWallet creates a new wallet with the given address and signature.
func (r *Repo) CreateWallet(address string, signature string) {
    var wallet = entity.Wallet{Address: address, Signature: signature}
    r.DBConnector.DB.Create(&wallet)
}
```

**Description:** This method creates a new wallet entity in the database with the provided address and signature. It initializes a new `Wallet` struct with the given parameters, and then uses GORM's `Create` method to insert the wallet entity into the database.

#### LoginWallet

```go
// LoginWallet updates the login timestamp of a wallet.
func (r *Repo) LoginWallet(address string) {
    var wallet = entity.Wallet{}
    r.DBConnector.DB.First(&wallet, &entity.Wallet{Address: address})
    wallet.LoginAt = time.Now()
    r.DBConnector.DB.Save(&wallet)
}
```

**Description:** This method updates the login timestamp of a wallet in the database. It retrieves the wallet entity from the database using the provided address, updates the `LoginAt` field with the current timestamp, and then saves the changes back to the database.

#### GetCreateWallet

```go
// GetCreateWallet retrieves a wallet by address, or creates it if it doesn't exist.
func (r *Repo) GetCreateWallet(address, signature string) *entity.Wallet {
    wallet := entity.Wallet{}
    r.DBConnector.DB.First(&wallet, &entity.Wallet{Address: address})
    if wallet.Address == "" {
        r.DBConnector.DB.Save(&entity.Wallet{Address: address, Signature: signature})
        r.DBConnector.DB.First(&wallet, &entity.Wallet{Address: address})
    }
    return &wallet
}
```

**Description:** This method retrieves a wallet entity from the database based on the provided address. If the wallet doesn't exist, it creates a new wallet entity with the provided address and signature. It first queries the database to check if a wallet with the given address exists. If not, it creates a new wallet entity with the provided address and signature, saves it to the database, and then retrieves the newly created wallet entity. Finally, it returns a pointer to the wallet entity.

### Transaction Repository

To be completed.

