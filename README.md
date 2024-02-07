# Introduction



Deposit Allow is an application developed in Golang, structured around several key modules aimed at managing a transaction to be confirmed by two users. These modules include:

* [**Contract**](contract.md): This module contains the Solidity smart contracts used within the application.
* [**Delivery**](delivery.md): The Delivery module focuses on facilitating the delivery process, encompassing functionalities related to APIs, CLIs, and other methods for exposing the application to users and external systems.
* [**Entity**](entities.md): Here, the Entity module consists of files related to transactions and wallets, crucial components for managing entities within the application.
* [**Repository**](repository.md): The Repository module encompasses the database operations, particularly focused on PostgreSQL for storage. It includes functionalities for interacting with the database and managing wallet data.
* [**Service**](service.md): The Service module provides essential business logic and functionalities related to wallets and transactions.
* [Front-end](front-end.md): The front-end part implemented using Metamask with web browser and it has some unimplemented parts for transaction handling and others.
* [Worker-Compose](worker-compose.md): This docker file contains all the parts regarding the database managements and it has decoupled from docker-compose which is responsible for the all other services.

\
