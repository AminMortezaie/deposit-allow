# Contract

### Key Concepts

#### Transaction Status

Transactions can have the following statuses:

* `Pending`: Initial state indicating the transaction has been initiated but not yet deposited or submitted.
* `Deposited`: Indicates that one party has deposited tokens, waiting for the other party to deposit.
* `Submitted`: Indicates that both parties have deposited tokens, waiting for execution.
* `Revoked`: Indicates that one party has revoked the transaction after submission but before execution.
* `Executed`: Indicates that the transaction has been successfully executed.
* `Withdrawn`: Indicates that one party has withdrawn tokens after the transaction time lock has expired.

#### Time Lock

A time lock mechanism ensures that transactions can only be executed after a specified period, providing a window for both parties to review and submit the transaction.

### Usage

#### Initiating a Transaction

To initiate a transaction, call the `initiateTransaction` function with the required parameters: user addresses, destination address, token addresses, and amounts.

```solidity
function initiateTransaction(
    address _userA,
    address _userB,
    address _destination,
    address _tokenA,
    address _tokenB,
    uint256 _amountA,
    uint256 _amountB
) external;
```

**Input:**

* `_userA`: Address of user A initiating the transaction.
* `_userB`: Address of user B participating in the transaction.
* `_destination`: Address where tokens will be sent after execution.
* `_tokenA`: Address of token A to be deposited.
* `_tokenB`: Address of token B to be deposited.
* `_amountA`: Amount of token A to be deposited.
* `_amountB`: Amount of token B to be deposited.

**Description:** This function initiates a new transaction between two parties. It creates a new transaction object with the provided details and stores it in the transaction set.

#### Depositing Tokens

Participants can deposit tokens using the `deposit` function by providing the transaction key and their index in the transaction array.

```solidity
function deposit(bytes32 _txKey, uint _index) external;
```

**Input:**

* `_txKey`: Key of the transaction to deposit tokens for.
* `_index`: Index of the transaction in the transaction array.

**Description:** This function allows a participant to deposit tokens for a specific transaction. It checks the participant's identity and status, transfers the specified token amount to the contract, and updates the transaction status accordingly.

#### Submitting a Transaction

Once both parties have deposited tokens, they can submit the transaction using the `submit` function. This triggers the execution of the transaction if both parties have submitted.

```solidity
function submit(bytes32 _txKey, uint _index) external;
```

**Input:**

* `_txKey`: Key of the transaction to submit.
* `_index`: Index of the transaction in the transaction array.

**Description:** This function allows a participant to submit a transaction once both parties have deposited tokens. It checks the participant's identity, status, and time lock, and if conditions are met, executes the transaction by transferring tokens to the destination address.

#### Revoking a Transaction

Before the transaction is executed, participants can revoke it using the `revoke` function if needed.

```solidity
function revoke(bytes32 _txKey, uint256 _index) external;
```

**Input:**

* `_txKey`: Key of the transaction to revoke.
* `_index`: Index of the transaction in the transaction array.

**Description:** This function allows a participant to revoke a transaction before it is executed. It checks the participant's identity and status and updates the transaction status to 'Revoked' if conditions are met.

#### Withdrawing Tokens

After the time lock has expired, participants can withdraw tokens using the `withdraw` function.

```solidity
function withdraw(bytes32 _txKey, uint256 _index) external;
```

**Input:**

* `_txKey`: Key of the transaction to withdraw tokens from.
* `_index`: Index of the transaction in the transaction array.

**Description:** This function allows a participant to withdraw tokens after the time lock has expired. It checks the participant's identity, time lock, and transaction status, and if conditions are met, transfers tokens back to the participant's address.

### Conclusion

The `DepositAllowance` smart contract provides a flexible and secure mechanism for managing transactions between parties, enabling efficient token transfers while ensuring proper validation and time lock mechanisms. By following the outlined usage guidelines, participants can leverage the contract's functionalities to facilitate smooth transaction processing within their applications.
