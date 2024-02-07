# Contract



### Introduction

The `DepositAllowance` smart contract facilitates secure and efficient transaction processing between two parties by allowing them to deposit tokens and execute transactions based on predefined conditions. This documentation provides an overview of the contract's functionalities, usage, and key components.

### Contract Overview

The `DepositAllowance` contract defines the following:

* A transaction structure to store transaction details such as user addresses, token addresses, amounts, statuses, and time locks.
* Functions for initiating transactions, depositing tokens, submitting transactions, revoking transactions, and withdrawing tokens.
* Events for deposit, submission, execution, withdrawal, and revocation of transactions.

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

#### Depositing Tokens

Participants can deposit tokens using the `deposit` function by providing the transaction key and their index in the transaction array.

#### Submitting a Transaction

Once both parties have deposited tokens, they can submit the transaction using the `submit` function. This triggers the execution of the transaction if both parties have submitted.

#### Revoking a Transaction

Before the transaction is executed, participants can revoke it using the `revoke` function if needed.

#### Withdrawing Tokens

After the time lock has expired, participants can withdraw tokens using the `withdraw` function.

### Conclusion

The `DepositAllowance` smart contract provides a flexible and secure mechanism for managing transactions between parties, enabling efficient token transfers while ensuring proper validation and time lock mechanisms. By following the outlined usage guidelines, participants can leverage the contract's functionalities to facilitate smooth transaction processing within their applications.
