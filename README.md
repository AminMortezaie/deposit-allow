# Introduction



Deposit Allow is an application developed in Golang, structured around several key modules aimed at managing a transaction to be confirmed by two users. These modules include:

* **Contract**: This module contains the Solidity smart contracts used within the application.
* **Delivery**: The Delivery module focuses on facilitating the delivery process, encompassing functionalities related to APIs, CLIs, and other methods for exposing the application to users and external systems.
* **Entity**: Here, the Entity module consists of files related to transactions and wallets, crucial components for managing entities within the application.
* **Repository**: The Repository module encompasses the database operations, particularly focused on PostgreSQL for storage. It includes functionalities for interacting with the database and managing wallet data.
* **Service**: The Service module provides essential business logic and functionalities related to wallets and transactions.

Throughout this documentation, we will delve into the details of each module, explaining their purpose, structure, and usage. Additionally, we'll provide guidance on how to set up and utilize the application effectively.
