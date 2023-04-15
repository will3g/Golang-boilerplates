# Golang with Kafka Boilerplate
This is a boilerplate project for building applications with **Golang**, **Kafka**, and **Confluentinc**.

![Dashboard](https://imgur.com/qJISiEH.png)

# Getting Started
To get started with this project, you will need to have Docker and Docker compose installed on your machine.

## Prerequisites
- Docker
- Docker compose

## Installation
- Clone this repository
- Navigate to the project directory
- Run **docker-compose up --build** to start the containers

## For new projects
- Remove the **go.mod** and **go.sum**
- Start this project
- Run **docker exec -it <container-app> bash** in your terminal
- Into the container, you need to run **go mod init github.com\/\<user\|company\>/\<project\>**
- After this execute **go mod tidy**

## Usage
After the installation, the following containers will be running:

- **go-with-kafka**
    > The Go application in **go-with-kafka** container can be modified to suit your needs. 
- **zookeeper**
    > Zookeeper is used as a **distributed coordination service and synchronization** tool to **manage the state of the Kafka cluster**
- **kafka**
    > The **Kafka cluster** is set up with three listeners: *PLAINTEXT*, *PLAINTEXT_HOST*, and *OUTSIDE*, with different advertised listeners. 
- **control-center**
    > The **Control Center container** can be used for **monitoring and managing the Kafka cluster**.

## Built with
- Golang
- Kafka
- Docker
- Confluentinc

## To execute this project in manual mode
```bash
root@go-with-kafka:/go/go-with-kafka# go run cmd/consumer/main.go

CONSUMER - INFO: Kafka consumer has started!
WORKER - INFO: Kafka worker has started!
```

> **PS:** *If you want to execute this project automatic, you need modified this **Dockerfile**.*

## To execute all tests

This will run all the tests in your project, including the **TestSavingOrder** test in **OrderRepositoryTestSuite**.

> **go test -v ./...**

**PS:** *The **-v** flag will enable verbose output, which will show you the details of each test that is run.*

## Authors
- **Will3g**

## Query used in this project for orders.db with sqlite3
> CREATE TABLE orders (id varchar(255) not null, price float not null, tax float not null, final_price float not null, primary key (id))
